// automatically generated by stateify.

package stack

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *linkAddrEntryList) beforeSave() {}
func (x *linkAddrEntryList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *linkAddrEntryList) afterLoad() {}
func (x *linkAddrEntryList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *linkAddrEntryEntry) beforeSave() {}
func (x *linkAddrEntryEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *linkAddrEntryEntry) afterLoad() {}
func (x *linkAddrEntryEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func (x *PacketBufferList) beforeSave() {}
func (x *PacketBufferList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *PacketBufferList) afterLoad() {}
func (x *PacketBufferList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *PacketBufferEntry) beforeSave() {}
func (x *PacketBufferEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *PacketBufferEntry) afterLoad() {}
func (x *PacketBufferEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func (x *TransportEndpointID) beforeSave() {}
func (x *TransportEndpointID) save(m state.Map) {
	x.beforeSave()
	m.Save("LocalPort", &x.LocalPort)
	m.Save("LocalAddress", &x.LocalAddress)
	m.Save("RemotePort", &x.RemotePort)
	m.Save("RemoteAddress", &x.RemoteAddress)
}

func (x *TransportEndpointID) afterLoad() {}
func (x *TransportEndpointID) load(m state.Map) {
	m.Load("LocalPort", &x.LocalPort)
	m.Load("LocalAddress", &x.LocalAddress)
	m.Load("RemotePort", &x.RemotePort)
	m.Load("RemoteAddress", &x.RemoteAddress)
}

func (x *GSOType) save(m state.Map) {
	m.SaveValue("", (int)(*x))
}

func (x *GSOType) load(m state.Map) {
	m.LoadValue("", new(int), func(y interface{}) { *x = (GSOType)(y.(int)) })
}

func (x *GSO) beforeSave() {}
func (x *GSO) save(m state.Map) {
	x.beforeSave()
	m.Save("Type", &x.Type)
	m.Save("NeedsCsum", &x.NeedsCsum)
	m.Save("CsumOffset", &x.CsumOffset)
	m.Save("MSS", &x.MSS)
	m.Save("L3HdrLen", &x.L3HdrLen)
	m.Save("MaxSize", &x.MaxSize)
}

func (x *GSO) afterLoad() {}
func (x *GSO) load(m state.Map) {
	m.Load("Type", &x.Type)
	m.Load("NeedsCsum", &x.NeedsCsum)
	m.Load("CsumOffset", &x.CsumOffset)
	m.Load("MSS", &x.MSS)
	m.Load("L3HdrLen", &x.L3HdrLen)
	m.Load("MaxSize", &x.MaxSize)
}

func (x *TransportEndpointInfo) beforeSave() {}
func (x *TransportEndpointInfo) save(m state.Map) {
	x.beforeSave()
	m.Save("NetProto", &x.NetProto)
	m.Save("TransProto", &x.TransProto)
	m.Save("ID", &x.ID)
	m.Save("BindNICID", &x.BindNICID)
	m.Save("BindAddr", &x.BindAddr)
	m.Save("RegisterNICID", &x.RegisterNICID)
}

func (x *TransportEndpointInfo) afterLoad() {}
func (x *TransportEndpointInfo) load(m state.Map) {
	m.Load("NetProto", &x.NetProto)
	m.Load("TransProto", &x.TransProto)
	m.Load("ID", &x.ID)
	m.Load("BindNICID", &x.BindNICID)
	m.Load("BindAddr", &x.BindAddr)
	m.Load("RegisterNICID", &x.RegisterNICID)
}

func (x *multiPortEndpoint) beforeSave() {}
func (x *multiPortEndpoint) save(m state.Map) {
	x.beforeSave()
	m.Save("demux", &x.demux)
	m.Save("netProto", &x.netProto)
	m.Save("transProto", &x.transProto)
	m.Save("endpoints", &x.endpoints)
	m.Save("flags", &x.flags)
}

func (x *multiPortEndpoint) afterLoad() {}
func (x *multiPortEndpoint) load(m state.Map) {
	m.Load("demux", &x.demux)
	m.Load("netProto", &x.netProto)
	m.Load("transProto", &x.transProto)
	m.Load("endpoints", &x.endpoints)
	m.Load("flags", &x.flags)
}

func init() {
	state.Register("pkg/tcpip/stack.linkAddrEntryList", (*linkAddrEntryList)(nil), state.Fns{Save: (*linkAddrEntryList).save, Load: (*linkAddrEntryList).load})
	state.Register("pkg/tcpip/stack.linkAddrEntryEntry", (*linkAddrEntryEntry)(nil), state.Fns{Save: (*linkAddrEntryEntry).save, Load: (*linkAddrEntryEntry).load})
	state.Register("pkg/tcpip/stack.PacketBufferList", (*PacketBufferList)(nil), state.Fns{Save: (*PacketBufferList).save, Load: (*PacketBufferList).load})
	state.Register("pkg/tcpip/stack.PacketBufferEntry", (*PacketBufferEntry)(nil), state.Fns{Save: (*PacketBufferEntry).save, Load: (*PacketBufferEntry).load})
	state.Register("pkg/tcpip/stack.TransportEndpointID", (*TransportEndpointID)(nil), state.Fns{Save: (*TransportEndpointID).save, Load: (*TransportEndpointID).load})
	state.Register("pkg/tcpip/stack.GSOType", (*GSOType)(nil), state.Fns{Save: (*GSOType).save, Load: (*GSOType).load})
	state.Register("pkg/tcpip/stack.GSO", (*GSO)(nil), state.Fns{Save: (*GSO).save, Load: (*GSO).load})
	state.Register("pkg/tcpip/stack.TransportEndpointInfo", (*TransportEndpointInfo)(nil), state.Fns{Save: (*TransportEndpointInfo).save, Load: (*TransportEndpointInfo).load})
	state.Register("pkg/tcpip/stack.multiPortEndpoint", (*multiPortEndpoint)(nil), state.Fns{Save: (*multiPortEndpoint).save, Load: (*multiPortEndpoint).load})
}
