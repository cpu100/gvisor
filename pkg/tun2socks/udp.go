package tun2socks

import (
    "io"
    "log"
    "net"
    "time"

    "gvisor.dev/gvisor/pkg/tcpip"
    "gvisor.dev/gvisor/pkg/tcpip/buffer"
    "gvisor.dev/gvisor/pkg/tcpip/transport/udp"
)

type UDPConn struct {
    ch chan buffer.View
    ep tcpip.Endpoint

    t2s *Tun2socks

    timeouts uint8

    readWriteOnce bool
}

func (conn *UDPConn) WriteTo(w io.Writer) (n int64, e error) {
    var v buffer.View
    var rTimer = time.NewTimer(time.Second*5)
    defer rTimer.Stop()
    for {
        select {
        case v = <-conn.ch:
            if nil == v {
                return
            } else {
                conn.timeouts = 0
            }
        case <-rTimer.C:
            if conn.readWriteOnce {
                conn.Close()
            } else if conn.timeouts < 3 {
                e = ErrTimeout
                conn.timeouts++
            } else {
                conn.Close()
            }
            return
        }
        for len(v) > 0 {
            nw, err := w.Write(v)
            if nil != err {
                if err2, ok := err.(net.Error); !ok || !err2.Temporary() {
                    e = err
                    return
                }
            }
            n += int64(nw)
            v.TrimFront(nw)
        }
        if conn.readWriteOnce {
            return
        }
    }
}

func (conn *UDPConn) ReadFrom(r io.Reader) (n int64, e error) {
    for {
        // todo udp payload size
        v := buffer.NewView(PAYLOAD)
        r.(HasReadDeadline).SetReadDeadline(time.Now().Add(time.Second))
        nr, err := r.Read(v)
        if nr > 0 {
            conn.timeouts = 0
            v.CapLength(nr)
            for len(v) > 0 {
                nw, ch, err2 := udp.Write(conn.ep, v)
                if nil != err2 {
                    if nil != ch {
                        log.Println("udp <-ch 1", e)
                        <-ch
                        log.Println("udp <-ch 2", e)
                    } else {
                        e = ErrClosedPipe
                        log.Println("[udp write]", err2)
                        return
                    }
                }
                n += nw
                v.TrimFront(int(nw))
            }
        }
        if nil != err {
            if err == io.EOF {
                // conn.ep.Shutdown(tcpip.ShutdownWrite)
            } else {
                e = err
            }
            return
        }
    }
}

func (conn *UDPConn) Read(b []byte) (int, error) {
    panic(ErrNotSupported)
}

func (conn *UDPConn) Write(b []byte) (int, error) {
    panic(ErrNotSupported)
}

func (conn *UDPConn) Close() error {
    conn.t2s.udpViews.Delete(*udp.TransportEndpointID(conn.ep))
    return nil
}

func (conn *UDPConn) LocalAddr() net.Addr {
    id := udp.TransportEndpointID(conn.ep)
    return &net.UDPAddr{
        IP:   []byte(id.RemoteAddress),
        Port: int(id.RemotePort),
    }
}

func (conn *UDPConn) RemoteAddr() net.Addr {
    id := udp.TransportEndpointID(conn.ep)
    return &net.UDPAddr{
        IP:   []byte(id.LocalAddress),
        Port: int(id.LocalPort),
    }
}

func (conn *UDPConn) SetDeadline(t time.Time) error {
    panic(ErrNotSupported)
}

func (conn *UDPConn) SetReadDeadline(t time.Time) error {
    panic(ErrNotSupported)
}

func (conn *UDPConn) SetWriteDeadline(t time.Time) error {
    panic(ErrNotSupported)
}

// For Question-Answer mode, such as dns query
// Release resources as quickly as possible
func (conn *UDPConn) SetReadWriteOnce() {
    conn.readWriteOnce = true
}
