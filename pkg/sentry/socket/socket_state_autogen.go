// automatically generated by stateify.

package socket

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *SendReceiveTimeout) StateTypeName() string {
	return "pkg/sentry/socket.SendReceiveTimeout"
}

func (x *SendReceiveTimeout) StateFields() []string {
	return []string{
		"send",
		"recv",
	}
}

func (x *SendReceiveTimeout) beforeSave() {}

func (x *SendReceiveTimeout) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.send)
	m.Save(1, &x.recv)
}

func (x *SendReceiveTimeout) afterLoad() {}

func (x *SendReceiveTimeout) StateLoad(m state.Source) {
	m.Load(0, &x.send)
	m.Load(1, &x.recv)
}

func init() {
	state.Register((*SendReceiveTimeout)(nil))
}
