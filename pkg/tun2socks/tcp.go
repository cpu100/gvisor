package tun2socks

import (
    "io"
    "log"
    "net"
    "time"

    "gvisor.dev/gvisor/pkg/tcpip"
    "gvisor.dev/gvisor/pkg/tcpip/buffer"
    "gvisor.dev/gvisor/pkg/tcpip/transport/tcp"
    "gvisor.dev/gvisor/pkg/waiter"
)

type TCPConn struct {
    wq *waiter.Queue
    ep tcpip.Endpoint
}

func (conn *TCPConn) WriteTo(w io.Writer) (n int64, e error) {
    waitEntry, notifyCh := waiter.NewChannelEntry(nil)
    conn.wq.EventRegister(&waitEntry, waiter.EventIn|waiter.EventErr)
    defer conn.wq.EventUnregister(&waitEntry)

    var rTimer = time.NewTimer(time.Second)

    for {
        v, _, err := conn.ep.Read(nil)
        for len(v) > 0 {
            nw, e2 := w.Write(v)
            if nil != e2 {
                if e3, ok := e2.(net.Error); !ok || !e3.Temporary() {
                    conn.ep.Shutdown(tcpip.ShutdownRead)
                    e = e2
                    return
                }
            }
            n += int64(nw)
            v.TrimFront(nw)
        }
        if err != nil {
            switch err {
            case tcpip.ErrWouldBlock:
                timerReset(rTimer, time.Second)
                select {
                case <-notifyCh:
                    continue
                case <-rTimer.C:
                    e = ErrTimeout
                }
            case tcpip.ErrClosedForReceive:
            case tcpip.ErrConnectionReset, tcpip.ErrConnectionAborted:
                e = ErrClosedPipe
            default:
                e = ErrClosedPipe
                log.Println(err)
            }

            return
        }
    }
}

func (conn *TCPConn) ReadFrom(r io.Reader) (n int64, e error) {
    waitEntry, notifyCh := waiter.NewChannelEntry(nil)
    conn.wq.EventRegister(&waitEntry, waiter.EventOut|waiter.EventErr)
    defer conn.wq.EventUnregister(&waitEntry)

    for {
        v := buffer.NewView(PAYLOAD)
        r.(HasReadDeadline).SetReadDeadline(time.Now().Add(time.Second))
        nr, err := r.Read(v)
        if nr > 0 {
            v.CapLength(nr)
            for len(v) > 0 {
                nw, _, err2 := conn.ep.Write(tcpip.SlicePayload(v), tcpip.WriteOptions{})
                if nil != err2 {
                    if err2 == tcpip.ErrWouldBlock {
                        <-notifyCh
                        continue
                    } else {
                        e = ErrClosedPipe
                        switch err2 {
                        case tcpip.ErrClosedForSend, tcpip.ErrConnectionReset, tcpip.ErrConnectionAborted:
                        default:
                            log.Println("[tcp write]", err2)
                        }
                        return
                    }
                }
                n += nw
                v.TrimFront(int(nw))
            }
        }
        if nil != err {
            if err == io.EOF {
                conn.ep.Shutdown(tcpip.ShutdownWrite)
            } else {
                e = err
            }
            return
        }
    }
}

func (conn *TCPConn) Read(b []byte) (int, error) {
    panic(ErrNotSupported)
}

func (conn *TCPConn) Write(b []byte) (int, error) {
    panic(ErrNotSupported)
}

func (conn *TCPConn) Close() error {
    conn.ep.Close()
    return nil
}

func (conn *TCPConn) LocalAddr() net.Addr {
    id := tcp.TransportEndpointID(conn.ep)
    return &net.TCPAddr{
        IP:   []byte(id.RemoteAddress),
        Port: int(id.RemotePort),
    }
}

func (conn *TCPConn) RemoteAddr() net.Addr {
    id := tcp.TransportEndpointID(conn.ep)
    return &net.TCPAddr{
        IP:   []byte(id.LocalAddress),
        Port: int(id.LocalPort),
    }
}

func (conn *TCPConn) SetDeadline(t time.Time) error {
    panic(ErrNotSupported)
}

func (conn *TCPConn) SetReadDeadline(t time.Time) error {
    panic(ErrNotSupported)
}

func (conn *TCPConn) SetWriteDeadline(t time.Time) error {
    panic(ErrNotSupported)
}
