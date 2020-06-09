package tun2socks

import (
    "io"
    "log"

    "gvisor.dev/gvisor/pkg/tcpip"
    "gvisor.dev/gvisor/pkg/tcpip/buffer"
    "gvisor.dev/gvisor/pkg/tcpip/network/ipv4"
    "gvisor.dev/gvisor/pkg/tcpip/network/ipv6"
    "gvisor.dev/gvisor/pkg/tcpip/stack"
)

// 1500 - ipHeader - tcpHeader - tcpTimestamp - packetCustomHeader
const PAYLOAD = 1500 - 20 - 20 - 12 - 10
const MTU = PAYLOAD + 20 + 20

type endpoint struct {
    tun io.ReadWriteCloser
}

func (e *endpoint) MTU() uint32 {
    return MTU
}

func (e *endpoint) Capabilities() stack.LinkEndpointCapabilities {
    return stack.CapabilityNone
}

// Tun has none link-layer header, just returns 0.
func (e *endpoint) MaxHeaderLength() uint16 {
    return 0
}

// Mac address is required by ICMPv6
func (e *endpoint) LinkAddress() tcpip.LinkAddress {
    // https://stackoverflow.com/questions/21018729/generate-mac-address-in-go
    // mac[0] = (mac[0] | 2) & 0xfe // Set local bit, ensure unicast address
    return "\xFE\xFF\xFF\xFF\xFF\xFF"
}

func (e *endpoint) WritePacket(r *stack.Route, gso *stack.GSO, protocol tcpip.NetworkProtocolNumber, pkt *stack.PacketBuffer) *tcpip.Error {
    vv := pkt.Header.View().ToVectorisedView()
    vv.Append(pkt.Data)
    nw, err := e.tun.Write(vv.ToView())
    if nil != err {
        log.Println(err)
        return tcpip.ErrInvalidEndpointState
    } else if nw != vv.Size() {
        panic(io.ErrShortWrite)
    } else {
        return nil
    }
}

func (e *endpoint) WritePackets(r *stack.Route, gso *stack.GSO, pkts stack.PacketBufferList, protocol tcpip.NetworkProtocolNumber) (int, *tcpip.Error) {
    panic(ErrNotSupported)
}

func (e *endpoint) WriteRawPacket(vv buffer.VectorisedView) *tcpip.Error {
    panic(ErrNotSupported)
}

func (e *endpoint) Attach(dispatcher stack.NetworkDispatcher) {
    go func(tun io.Reader) {
        for {
            v := buffer.NewView(MTU)
            nr, err := tun.Read(v)
            if nil != err {
                log.Println(err)
                break
            }

            v.CapLength(nr)

            if 0x40 == v[0]&0xf0 {
                dispatcher.DeliverNetworkPacket("", "", ipv4.ProtocolNumber, &stack.PacketBuffer{
                    Data: v.ToVectorisedView(),
                })
            } else {
                // https://en.wikipedia.org/wiki/List_of_IP_Protocol_numbers
                // header.IPv6(v).TransportProtocol() == header.ICMPv6ProtocolNumber
                dispatcher.DeliverNetworkPacket("", "", ipv6.ProtocolNumber, &stack.PacketBuffer{
                    Data: v.ToVectorisedView(),
                })
            }
        }
    }(e.tun)
}

func (e *endpoint) IsAttached() bool {
    panic(ErrNotSupported)
}

// Called by Stack.Wait()
// Once tun is closed, tun.Read(v) return error, the loop in Attach() break.
func (e *endpoint) Wait() {
    e.tun.Close()
}
