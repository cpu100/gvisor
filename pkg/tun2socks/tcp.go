package tun2socks

import (
    "fmt"
    "io"
    "log"
    "net"
    "time"

    "gvisor.dev/gvisor/pkg/tcpip"
    "gvisor.dev/gvisor/pkg/tcpip/transport/tcp"
    "gvisor.dev/gvisor/pkg/waiter"
)

type TCPConn struct {
    wq *waiter.Queue
    ep tcpip.Endpoint
}

func (conn *TCPConn) WriteTo(w io.Writer) (n int64, e error) {
    waitEntry, notifyCh := waiter.NewChannelEntry(waiter.EventIn|waiter.EventErr)
    conn.wq.EventRegister(&waitEntry)
    defer conn.wq.EventUnregister(&waitEntry)

    var rTimer = time.NewTimer(time.Second)

    for {
        res, err := conn.ep.Read(w, tcpip.ReadOptions{})
        n += int64(res.Count)
        if err == nil {
            continue
        } else if _, ok := err.(*tcpip.ErrWouldBlock); ok {
            timerReset(rTimer, time.Second)
            select {
            case <-notifyCh:
                continue
            case <-rTimer.C:
                e = ErrTimeout
            }
        } else if err2, ok := err.(net.Error); ok {
            // 可能代码永远都不会走这里 todo
            if err2.Temporary() {
                continue
            } else {
                e = err2
                conn.ep.Shutdown(tcpip.ShutdownRead)
            }
        } else if _, ok := err.(*tcpip.ErrClosedForReceive); ok {
        } else if _, ok := err.(*tcpip.ErrConnectionReset); ok {
            e = ErrClosedPipe
        } else if _, ok := err.(*tcpip.ErrConnectionAborted); ok {
            e = ErrClosedPipe
        } else {
            e = ErrClosedPipe
            log.Println(err)
        }
        return
    }
}

func (conn *TCPConn) ReadFrom(r io.Reader) (n int64, e error) {
    waitEntry, notifyCh := waiter.NewChannelEntry(waiter.EventErr|waiter.WritableEvents)
    conn.wq.EventRegister(&waitEntry)
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
        IP:   id.RemoteAddress.AsSlice(),
        Port: int(id.RemotePort),
    }
}

func (conn *TCPConn) RemoteAddr() net.Addr {
    id := tcp.TransportEndpointID(conn.ep)
    return &net.TCPAddr{
        IP:   id.LocalAddress.AsSlice(),
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
