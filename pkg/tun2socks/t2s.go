package tun2socks

import (
    "io"
    "log"
    "runtime"
    "sync"

    "gvisor.dev/gvisor/pkg/tcpip"
    "gvisor.dev/gvisor/pkg/tcpip/buffer"
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

    s := stack.New(stack.Options{
        NetworkProtocols:   []stack.NetworkProtocol{ipv4.NewProtocol(), ipv6.NewProtocol()},
        TransportProtocols: []stack.TransportProtocol{tcp.NewProtocol(), udp.NewProtocol()},
    })

    if err := s.CreateNIC(1, &endpoint{tun: tun}); err != nil {
        panic(err)
    }

    if err := s.AddAddressRange(1, ipv4.ProtocolNumber, header.IPv4EmptySubnet); err != nil {
        panic(err)
    }

    if err := s.AddAddressRange(1, ipv6.ProtocolNumber, header.IPv6EmptySubnet); err != nil {
        panic(err)
    }

    // a default route is required by udp sending
    s.SetRouteTable([]tcpip.Route{
        {NIC: 1, Destination: header.IPv4EmptySubnet},
        {NIC: 1, Destination: header.IPv6EmptySubnet},
    })

    t2s := &Tun2socks{th: th, s: s}

    t2s.tcpAccept(ipv4.ProtocolNumber)
    t2s.tcpAccept(ipv6.ProtocolNumber)
    t2s.udpAccept(ipv4.ProtocolNumber)
    t2s.udpAccept(ipv6.ProtocolNumber)

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
        waitEntry, notifyCh := waiter.NewChannelEntry(nil)
        wq.EventRegister(&waitEntry, waiter.EventIn|waiter.EventErr)
        defer wq.EventUnregister(&waitEntry)
        for {
            ep2, wq2, err := ep.Accept()
            if err != nil {
                if err == tcpip.ErrWouldBlock {
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

    if err := ep.Bind(tcpip.FullAddress{NIC: 1, Port: uint16(netProto)}); err != nil {
        panic(err)
    }

    go func() {
        waitEntry, notifyCh := waiter.NewChannelEntry(nil)
        wq.EventRegister(&waitEntry, waiter.EventIn|waiter.EventErr)
        defer wq.EventUnregister(&waitEntry)

        var ipLen = header.IPv4AddressSize
        if ipv6.ProtocolNumber == netProto {
            ipLen = header.IPv6AddressSize
        }

        var chView chan buffer.View
        var addr = tcpip.FullAddress{}
        for {
            v, _, err := ep.Read(&addr)
            if err != nil {
                if err == tcpip.ErrWouldBlock {
                    <-notifyCh
                    continue
                } else {
                    log.Println(err)
                    break
                }
            }

            id := stack.TransportEndpointID{
                RemoteAddress: addr.Addr[:ipLen],
                RemotePort:    addr.Port,
                LocalAddress:  addr.Addr[ipLen:],
                LocalPort:     uint16(addr.NIC),
            }

            if ch, ok := t2s.udpViews.Load(id); ok {
                chView = ch.(chan buffer.View)
            } else {
                actual, loaded := t2s.udpViews.LoadOrStore(id, make(chan buffer.View, 8))
                chView = actual.(chan buffer.View)
                if !loaded {
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
    conn := &TCPConn{wq: wq, ep: ep}
    if nil != t2s.th.TcpHandle(conn) {
        conn.Close()
    } else {
        runtime.SetFinalizer(conn, (*TCPConn).Close)
    }
}

func (t2s *Tun2socks) udpConnect(ep tcpip.Endpoint, ch chan buffer.View) {
    conn := &UDPConn{ep: ep, ch: ch, t2s: t2s}
    if nil != t2s.th.UdpHandle(conn) {
        conn.Close()
    } else {
        runtime.SetFinalizer(conn, (*UDPConn).Close)
    }
}

func (t2s *Tun2socks) Close() error {
    t2s.s.Close()
    t2s.s.Wait()
    t2s.udpViews.Range(func(_, ch interface{}) bool {
        ch.(chan buffer.View) <- nil
        return true
    })
    return nil
}
