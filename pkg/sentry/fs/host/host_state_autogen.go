// automatically generated by stateify.

package host

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *descriptor) StateTypeName() string {
	return "pkg/sentry/fs/host.descriptor"
}

func (x *descriptor) StateFields() []string {
	return []string{
		"origFD",
		"wouldBlock",
	}
}

func (x *descriptor) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.origFD)
	m.Save(1, &x.wouldBlock)
}

func (x *descriptor) StateLoad(m state.Source) {
	m.Load(0, &x.origFD)
	m.Load(1, &x.wouldBlock)
	m.AfterLoad(x.afterLoad)
}

func (x *fileOperations) StateTypeName() string {
	return "pkg/sentry/fs/host.fileOperations"
}

func (x *fileOperations) StateFields() []string {
	return []string{
		"iops",
		"dirCursor",
	}
}

func (x *fileOperations) beforeSave() {}

func (x *fileOperations) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.iops)
	m.Save(1, &x.dirCursor)
}

func (x *fileOperations) afterLoad() {}

func (x *fileOperations) StateLoad(m state.Source) {
	m.LoadWait(0, &x.iops)
	m.Load(1, &x.dirCursor)
}

func (x *filesystem) StateTypeName() string {
	return "pkg/sentry/fs/host.filesystem"
}

func (x *filesystem) StateFields() []string {
	return []string{}
}

func (x *filesystem) beforeSave() {}

func (x *filesystem) StateSave(m state.Sink) {
	x.beforeSave()
}

func (x *filesystem) afterLoad() {}

func (x *filesystem) StateLoad(m state.Source) {
}

func (x *inodeOperations) StateTypeName() string {
	return "pkg/sentry/fs/host.inodeOperations"
}

func (x *inodeOperations) StateFields() []string {
	return []string{
		"fileState",
		"cachingInodeOps",
	}
}

func (x *inodeOperations) beforeSave() {}

func (x *inodeOperations) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.fileState)
	m.Save(1, &x.cachingInodeOps)
}

func (x *inodeOperations) afterLoad() {}

func (x *inodeOperations) StateLoad(m state.Source) {
	m.LoadWait(0, &x.fileState)
	m.Load(1, &x.cachingInodeOps)
}

func (x *inodeFileState) StateTypeName() string {
	return "pkg/sentry/fs/host.inodeFileState"
}

func (x *inodeFileState) StateFields() []string {
	return []string{
		"descriptor",
		"sattr",
		"savedUAttr",
	}
}

func (x *inodeFileState) beforeSave() {}

func (x *inodeFileState) StateSave(m state.Sink) {
	x.beforeSave()
	if !state.IsZeroValue(&x.queue) {
		state.Failf("queue is %#v, expected zero", &x.queue)
	}
	m.Save(0, &x.descriptor)
	m.Save(1, &x.sattr)
	m.Save(2, &x.savedUAttr)
}

func (x *inodeFileState) StateLoad(m state.Source) {
	m.LoadWait(0, &x.descriptor)
	m.LoadWait(1, &x.sattr)
	m.Load(2, &x.savedUAttr)
	m.AfterLoad(x.afterLoad)
}

func (x *ConnectedEndpoint) StateTypeName() string {
	return "pkg/sentry/fs/host.ConnectedEndpoint"
}

func (x *ConnectedEndpoint) StateFields() []string {
	return []string{
		"ref",
		"queue",
		"path",
		"srfd",
		"stype",
	}
}

func (x *ConnectedEndpoint) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.ref)
	m.Save(1, &x.queue)
	m.Save(2, &x.path)
	m.Save(3, &x.srfd)
	m.Save(4, &x.stype)
}

func (x *ConnectedEndpoint) StateLoad(m state.Source) {
	m.Load(0, &x.ref)
	m.Load(1, &x.queue)
	m.Load(2, &x.path)
	m.LoadWait(3, &x.srfd)
	m.Load(4, &x.stype)
	m.AfterLoad(x.afterLoad)
}

func (x *TTYFileOperations) StateTypeName() string {
	return "pkg/sentry/fs/host.TTYFileOperations"
}

func (x *TTYFileOperations) StateFields() []string {
	return []string{
		"fileOperations",
		"session",
		"fgProcessGroup",
		"termios",
	}
}

func (x *TTYFileOperations) beforeSave() {}

func (x *TTYFileOperations) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.fileOperations)
	m.Save(1, &x.session)
	m.Save(2, &x.fgProcessGroup)
	m.Save(3, &x.termios)
}

func (x *TTYFileOperations) afterLoad() {}

func (x *TTYFileOperations) StateLoad(m state.Source) {
	m.Load(0, &x.fileOperations)
	m.Load(1, &x.session)
	m.Load(2, &x.fgProcessGroup)
	m.Load(3, &x.termios)
}

func init() {
	state.Register((*descriptor)(nil))
	state.Register((*fileOperations)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*inodeOperations)(nil))
	state.Register((*inodeFileState)(nil))
	state.Register((*ConnectedEndpoint)(nil))
	state.Register((*TTYFileOperations)(nil))
}
