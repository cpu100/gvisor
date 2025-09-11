package tun2socks

import (
    "io"
    "log"
    "runtime"
    "sync"

    "gvisor.dev/gvisor/pkg/buffer"
    "gvisor.dev/gvisor/pkg/tcpip"
    "gvisor.dev/gvisor/pkg/tcpip/adapters/gonet"
    "gvisor.dev/gvisor/pkg/tcpip/header"
    "gvisor.dev/gvisor/pkg/tcpip/network/ipv4"
    "gvisor.dev/gvisor/pkg/tcpip/network/ipv6"
    "gvisor.dev/gvisor/pkg/tcpip/stack"
    "gvisor.dev/gvisor/pkg/tcpip/transport/tcp"
    "gvisor.dev/gvisor/pkg/tcpip/transport/udp"
    "gvisor.dev/gvisor/pkg/waiter"
)

type Tun2socks struct {
    s  *stack.Stack
    th TransportHandler
    // TransportEndpointID => chan buffer.View
    udpViews sync.Map
}

func New(tun io.ReadWriteCloser, th TransportHandler) *Tun2socks {

    log.SetFlags(log.LstdFlags | log.Lshortfile)

    s := stack.New(stack.Options{
        // icmp.NewProtocol4,
        NetworkProtocols:   []stack.NetworkProtocolFactory{ipv4.NewProtocol, ipv6.NewProtocol},
        TransportProtocols: []stack.TransportProtocolFactory{tcp.NewProtocol, udp.NewProtocol},
    })

    if err := s.CreateNIC(1, &endpoint{tun: tun}); err != nil {
        panic(err)
    }

    if err := s.AddProtocolAddress(1, tcpip.ProtocolAddress{ipv4.ProtocolNumber, tcpip.AddressWithPrefix{tcpip.AddrFrom4([4]byte{111,230,9,178}), 32}}, stack.AddressProperties{}); err != nil {
        panic(err)
    }

    // if err := s.AddProtocolAddress(1, tcpip.ProtocolAddress{ipv4.ProtocolNumber, tcpip.AddressWithPrefix{header.IPv4Any, 0}}, stack.AddressProperties{}); err != nil {
    //     panic(err)
    // }

    // if err := s.AddProtocolAddress(1, tcpip.ProtocolAddress{ipv6.ProtocolNumber, tcpip.AddressWithPrefix{header.IPv6Any, 0}}, stack.AddressProperties{}); err != nil {
    //     panic(err)
    // }

    // s.AddAddressRange(1, ipv4.ProtocolNumber, header.IPv4EmptySubnet)
    // s.AddAddressRange(1, ipv6.ProtocolNumber, header.IPv6EmptySubnet)
    s.SetPromiscuousMode(1, true)

    // a default route is required by udp sending
    s.SetRouteTable([]tcpip.Route{
        {NIC: 1, Destination: header.IPv4EmptySubnet},
        {NIC: 1, Destination: header.IPv6EmptySubnet},
    })

    t2s := &Tun2socks{th: th, s: s}

    t2s.tcpAccept(ipv4.ProtocolNumber)
    t2s.tcpAccept(ipv6.ProtocolNumber)
    // t2s.udpAccept(ipv4.ProtocolNumber)
    // t2s.udpAccept(ipv6.ProtocolNumber)

    return t2s
}

func (t2s *Tun2socks) tcpAccept(netProto tcpip.NetworkProtocolNumber) {
    var wq waiter.Queue
    ep, err2 := t2s.s.NewEndpoint(tcp.ProtocolNumber, netProto, &wq)
    if err2 != nil {
        panic(err2)
    }

    if err := ep.Bind(tcpip.FullAddress{NIC: 1, Port: uint16(netProto)}); err != nil {
        panic(err)
    }

    // make(chan *endpoint, backlog)
    if err := ep.Listen(16); err != nil {
        panic(err)
    }

    go func() {
        waitEntry, notifyCh := waiter.NewChannelEntry(waiter.EventErr|waiter.ReadableEvents)
        wq.EventRegister(&waitEntry)
        defer wq.EventUnregister(&waitEntry)
        for {
            ep2, wq2, err := ep.Accept(nil)
            if err != nil {
                if _, ok := err.(*tcpip.ErrWouldBlock); ok {
                    <-notifyCh
                    continue
                } else {
                    log.Println(err)
                    break
                }
            }

            go t2s.tcpConnect(wq2, ep2)
        }
    }()
}

func (t2s *Tun2socks) udpAccept(netProto tcpip.NetworkProtocolNumber) {
    var wq waiter.Queue
    ep, err := t2s.s.NewEndpoint(udp.ProtocolNumber, netProto, &wq)
    if err != nil {
        panic(err)
    }

    // 好像并未发送 RemoteAddress 是缺失的
    if err := ep.Bind(tcpip.FullAddress{NIC: 1, Port: uint16(netProto)}); err != nil {
        panic(err)
    }

    // ep.Accept() 是可以用的吗

    go func() {
        waitEntry, notifyCh := waiter.NewChannelEntry(waiter.EventIn|waiter.EventErr)
        wq.EventRegister(&waitEntry)
        defer wq.EventUnregister(&waitEntry)

        // var ipLen = header.IPv4AddressSize
        // if ipv6.ProtocolNumber == netProto {
        //     ipLen = header.IPv6AddressSize
        // }

         v := buffer.NewView(MTU*2)
        var chView chan *buffer.View
        // var addr = tcpip.FullAddress{}
        for {
            res, err := ep.Read(v, tcpip.ReadOptions{
                NeedRemoteAddr: true,
            })
            if err != nil {
                if _, ok := err.(*tcpip.ErrWouldBlock); ok {
                    <-notifyCh
                    continue
                } else {
                    log.Println(err)
                    break
                }
            }

            id := stack.TransportEndpointID{
                // RemoteAddress: addr.Addr[:ipLen],
                // RemotePort:    addr.Port,
                RemoteAddress: res.RemoteAddr.Addr,
                RemotePort: res.RemoteAddr.Port,
                // LocalAddress:  addr.Addr[ipLen:],
                // LocalPort:     uint16(addr.NIC),
            }

            if ch, ok := t2s.udpViews.Load(id); ok {
                chView = ch.(chan *buffer.View)
            } else {
                actual, loaded := t2s.udpViews.LoadOrStore(id, make(chan *buffer.View, 8))
                chView = actual.(chan *buffer.View)
                if !loaded {
                    // 好像不对，多个read?
                    go t2s.udpConnect(udp.EndpointWithWriteOptions(ep, &id), chView)
                }
            }
            if len(chView) < cap(chView) {
                chView <- v
            }
        }
    }()
}

func (t2s *Tun2socks) tcpConnect(wq *waiter.Queue, ep tcpip.Endpoint) {
    // conn := &TCPConn{wq: wq, ep: ep}
    conn := gonet.NewTCPConn(wq, ep)
    if nil != t2s.th.TcpHandle(conn) {
        conn.Close()
    } else {
        runtime.SetFinalizer(conn, (*gonet.TCPConn).Close)
    }
}

func (t2s *Tun2socks) udpConnect(ep tcpip.Endpoint, ch chan *buffer.View) {
    conn := &UDPConn{ep: ep, ch: ch, t2s: t2s}
    // conn := gonet.NewUDPConn(t2s.s, )
    if nil != t2s.th.UdpHandle(conn) {
        conn.Close()
    } else {
        runtime.SetFinalizer(conn, (*gonet.UDPConn).Close)
    }
}

func (t2s *Tun2socks) Close() error {
    t2s.s.Close()
    t2s.s.Wait()
    t2s.udpViews.Range(func(_, ch interface{}) bool {
        // ch.(chan any) <- buffer.View{} // close
        ch.(chan any) <- nil // close
        return true
    })
    return nil
}
