// automatically generated by stateify.

package vfs

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *Dentry) beforeSave() {}
func (x *Dentry) save(m state.Map) {
	x.beforeSave()
	m.Save("dead", &x.dead)
	m.Save("mounts", &x.mounts)
	m.Save("impl", &x.impl)
}

func (x *Dentry) afterLoad() {}
func (x *Dentry) load(m state.Map) {
	m.Load("dead", &x.dead)
	m.Load("mounts", &x.mounts)
	m.Load("impl", &x.impl)
}

func (x *registeredDevice) beforeSave() {}
func (x *registeredDevice) save(m state.Map) {
	x.beforeSave()
	m.Save("dev", &x.dev)
	m.Save("opts", &x.opts)
}

func (x *registeredDevice) afterLoad() {}
func (x *registeredDevice) load(m state.Map) {
	m.Load("dev", &x.dev)
	m.Load("opts", &x.opts)
}

func (x *RegisterDeviceOptions) beforeSave() {}
func (x *RegisterDeviceOptions) save(m state.Map) {
	x.beforeSave()
	m.Save("GroupName", &x.GroupName)
}

func (x *RegisterDeviceOptions) afterLoad() {}
func (x *RegisterDeviceOptions) load(m state.Map) {
	m.Load("GroupName", &x.GroupName)
}

func (x *epollInterestList) beforeSave() {}
func (x *epollInterestList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *epollInterestList) afterLoad() {}
func (x *epollInterestList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *epollInterestEntry) beforeSave() {}
func (x *epollInterestEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *epollInterestEntry) afterLoad() {}
func (x *epollInterestEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func (x *eventList) beforeSave() {}
func (x *eventList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *eventList) afterLoad() {}
func (x *eventList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *eventEntry) beforeSave() {}
func (x *eventEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *eventEntry) afterLoad() {}
func (x *eventEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func (x *Filesystem) beforeSave() {}
func (x *Filesystem) save(m state.Map) {
	x.beforeSave()
	m.Save("refs", &x.refs)
	m.Save("vfs", &x.vfs)
	m.Save("fsType", &x.fsType)
	m.Save("impl", &x.impl)
}

func (x *Filesystem) afterLoad() {}
func (x *Filesystem) load(m state.Map) {
	m.Load("refs", &x.refs)
	m.Load("vfs", &x.vfs)
	m.Load("fsType", &x.fsType)
	m.Load("impl", &x.impl)
}

func (x *registeredFilesystemType) beforeSave() {}
func (x *registeredFilesystemType) save(m state.Map) {
	x.beforeSave()
	m.Save("fsType", &x.fsType)
	m.Save("opts", &x.opts)
}

func (x *registeredFilesystemType) afterLoad() {}
func (x *registeredFilesystemType) load(m state.Map) {
	m.Load("fsType", &x.fsType)
	m.Load("opts", &x.opts)
}

func (x *Inotify) beforeSave() {}
func (x *Inotify) save(m state.Map) {
	x.beforeSave()
	m.Save("vfsfd", &x.vfsfd)
	m.Save("FileDescriptionDefaultImpl", &x.FileDescriptionDefaultImpl)
	m.Save("DentryMetadataFileDescriptionImpl", &x.DentryMetadataFileDescriptionImpl)
	m.Save("NoLockFD", &x.NoLockFD)
	m.Save("id", &x.id)
	m.Save("events", &x.events)
	m.Save("scratch", &x.scratch)
	m.Save("nextWatchMinusOne", &x.nextWatchMinusOne)
	m.Save("watches", &x.watches)
}

func (x *Inotify) afterLoad() {}
func (x *Inotify) load(m state.Map) {
	m.Load("vfsfd", &x.vfsfd)
	m.Load("FileDescriptionDefaultImpl", &x.FileDescriptionDefaultImpl)
	m.Load("DentryMetadataFileDescriptionImpl", &x.DentryMetadataFileDescriptionImpl)
	m.Load("NoLockFD", &x.NoLockFD)
	m.Load("id", &x.id)
	m.Load("events", &x.events)
	m.Load("scratch", &x.scratch)
	m.Load("nextWatchMinusOne", &x.nextWatchMinusOne)
	m.Load("watches", &x.watches)
}

func (x *Watches) beforeSave() {}
func (x *Watches) save(m state.Map) {
	x.beforeSave()
	m.Save("ws", &x.ws)
}

func (x *Watches) afterLoad() {}
func (x *Watches) load(m state.Map) {
	m.Load("ws", &x.ws)
}

func (x *Watch) beforeSave() {}
func (x *Watch) save(m state.Map) {
	x.beforeSave()
	m.Save("owner", &x.owner)
	m.Save("wd", &x.wd)
	m.Save("set", &x.set)
	m.Save("mask", &x.mask)
}

func (x *Watch) afterLoad() {}
func (x *Watch) load(m state.Map) {
	m.Load("owner", &x.owner)
	m.Load("wd", &x.wd)
	m.Load("set", &x.set)
	m.Load("mask", &x.mask)
}

func (x *Event) beforeSave() {}
func (x *Event) save(m state.Map) {
	x.beforeSave()
	m.Save("eventEntry", &x.eventEntry)
	m.Save("wd", &x.wd)
	m.Save("mask", &x.mask)
	m.Save("cookie", &x.cookie)
	m.Save("len", &x.len)
	m.Save("name", &x.name)
}

func (x *Event) afterLoad() {}
func (x *Event) load(m state.Map) {
	m.Load("eventEntry", &x.eventEntry)
	m.Load("wd", &x.wd)
	m.Load("mask", &x.mask)
	m.Load("cookie", &x.cookie)
	m.Load("len", &x.len)
	m.Load("name", &x.name)
}

func (x *Mount) beforeSave() {}
func (x *Mount) save(m state.Map) {
	x.beforeSave()
	m.Save("vfs", &x.vfs)
	m.Save("fs", &x.fs)
	m.Save("root", &x.root)
	m.Save("ID", &x.ID)
	m.Save("Flags", &x.Flags)
	m.Save("key", &x.key)
	m.Save("ns", &x.ns)
	m.Save("refs", &x.refs)
	m.Save("children", &x.children)
	m.Save("umounted", &x.umounted)
	m.Save("writers", &x.writers)
}

func (x *Mount) afterLoad() {}
func (x *Mount) load(m state.Map) {
	m.Load("vfs", &x.vfs)
	m.Load("fs", &x.fs)
	m.Load("root", &x.root)
	m.Load("ID", &x.ID)
	m.Load("Flags", &x.Flags)
	m.Load("key", &x.key)
	m.Load("ns", &x.ns)
	m.Load("refs", &x.refs)
	m.Load("children", &x.children)
	m.Load("umounted", &x.umounted)
	m.Load("writers", &x.writers)
}

func (x *MountNamespace) beforeSave() {}
func (x *MountNamespace) save(m state.Map) {
	x.beforeSave()
	m.Save("Owner", &x.Owner)
	m.Save("root", &x.root)
	m.Save("refs", &x.refs)
	m.Save("mountpoints", &x.mountpoints)
}

func (x *MountNamespace) afterLoad() {}
func (x *MountNamespace) load(m state.Map) {
	m.Load("Owner", &x.Owner)
	m.Load("root", &x.root)
	m.Load("refs", &x.refs)
	m.Load("mountpoints", &x.mountpoints)
}

func (x *VirtualFilesystem) beforeSave() {}
func (x *VirtualFilesystem) save(m state.Map) {
	x.beforeSave()
	m.Save("mounts", &x.mounts)
	m.Save("mountpoints", &x.mountpoints)
	m.Save("lastMountID", &x.lastMountID)
	m.Save("anonMount", &x.anonMount)
	m.Save("devices", &x.devices)
	m.Save("anonBlockDevMinorNext", &x.anonBlockDevMinorNext)
	m.Save("anonBlockDevMinor", &x.anonBlockDevMinor)
	m.Save("fsTypes", &x.fsTypes)
	m.Save("filesystems", &x.filesystems)
}

func (x *VirtualFilesystem) afterLoad() {}
func (x *VirtualFilesystem) load(m state.Map) {
	m.Load("mounts", &x.mounts)
	m.Load("mountpoints", &x.mountpoints)
	m.Load("lastMountID", &x.lastMountID)
	m.Load("anonMount", &x.anonMount)
	m.Load("devices", &x.devices)
	m.Load("anonBlockDevMinorNext", &x.anonBlockDevMinorNext)
	m.Load("anonBlockDevMinor", &x.anonBlockDevMinor)
	m.Load("fsTypes", &x.fsTypes)
	m.Load("filesystems", &x.filesystems)
}

func (x *VirtualDentry) beforeSave() {}
func (x *VirtualDentry) save(m state.Map) {
	x.beforeSave()
	m.Save("mount", &x.mount)
	m.Save("dentry", &x.dentry)
}

func (x *VirtualDentry) afterLoad() {}
func (x *VirtualDentry) load(m state.Map) {
	m.Load("mount", &x.mount)
	m.Load("dentry", &x.dentry)
}

func init() {
	state.Register("pkg/sentry/vfs.Dentry", (*Dentry)(nil), state.Fns{Save: (*Dentry).save, Load: (*Dentry).load})
	state.Register("pkg/sentry/vfs.registeredDevice", (*registeredDevice)(nil), state.Fns{Save: (*registeredDevice).save, Load: (*registeredDevice).load})
	state.Register("pkg/sentry/vfs.RegisterDeviceOptions", (*RegisterDeviceOptions)(nil), state.Fns{Save: (*RegisterDeviceOptions).save, Load: (*RegisterDeviceOptions).load})
	state.Register("pkg/sentry/vfs.epollInterestList", (*epollInterestList)(nil), state.Fns{Save: (*epollInterestList).save, Load: (*epollInterestList).load})
	state.Register("pkg/sentry/vfs.epollInterestEntry", (*epollInterestEntry)(nil), state.Fns{Save: (*epollInterestEntry).save, Load: (*epollInterestEntry).load})
	state.Register("pkg/sentry/vfs.eventList", (*eventList)(nil), state.Fns{Save: (*eventList).save, Load: (*eventList).load})
	state.Register("pkg/sentry/vfs.eventEntry", (*eventEntry)(nil), state.Fns{Save: (*eventEntry).save, Load: (*eventEntry).load})
	state.Register("pkg/sentry/vfs.Filesystem", (*Filesystem)(nil), state.Fns{Save: (*Filesystem).save, Load: (*Filesystem).load})
	state.Register("pkg/sentry/vfs.registeredFilesystemType", (*registeredFilesystemType)(nil), state.Fns{Save: (*registeredFilesystemType).save, Load: (*registeredFilesystemType).load})
	state.Register("pkg/sentry/vfs.Inotify", (*Inotify)(nil), state.Fns{Save: (*Inotify).save, Load: (*Inotify).load})
	state.Register("pkg/sentry/vfs.Watches", (*Watches)(nil), state.Fns{Save: (*Watches).save, Load: (*Watches).load})
	state.Register("pkg/sentry/vfs.Watch", (*Watch)(nil), state.Fns{Save: (*Watch).save, Load: (*Watch).load})
	state.Register("pkg/sentry/vfs.Event", (*Event)(nil), state.Fns{Save: (*Event).save, Load: (*Event).load})
	state.Register("pkg/sentry/vfs.Mount", (*Mount)(nil), state.Fns{Save: (*Mount).save, Load: (*Mount).load})
	state.Register("pkg/sentry/vfs.MountNamespace", (*MountNamespace)(nil), state.Fns{Save: (*MountNamespace).save, Load: (*MountNamespace).load})
	state.Register("pkg/sentry/vfs.VirtualFilesystem", (*VirtualFilesystem)(nil), state.Fns{Save: (*VirtualFilesystem).save, Load: (*VirtualFilesystem).load})
	state.Register("pkg/sentry/vfs.VirtualDentry", (*VirtualDentry)(nil), state.Fns{Save: (*VirtualDentry).save, Load: (*VirtualDentry).load})
}
