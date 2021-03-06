// automatically generated by stateify.

package pipe

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *inodeOperations) StateTypeName() string {
	return "pkg/sentry/kernel/pipe.inodeOperations"
}

func (x *inodeOperations) StateFields() []string {
	return []string{
		"InodeSimpleAttributes",
		"p",
	}
}

func (x *inodeOperations) beforeSave() {}

func (x *inodeOperations) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.InodeSimpleAttributes)
	m.Save(1, &x.p)
}

func (x *inodeOperations) afterLoad() {}

func (x *inodeOperations) StateLoad(m state.Source) {
	m.Load(0, &x.InodeSimpleAttributes)
	m.Load(1, &x.p)
}

func (x *Pipe) StateTypeName() string {
	return "pkg/sentry/kernel/pipe.Pipe"
}

func (x *Pipe) StateFields() []string {
	return []string{
		"isNamed",
		"atomicIOBytes",
		"readers",
		"writers",
		"view",
		"max",
		"hadWriter",
	}
}

func (x *Pipe) beforeSave() {}

func (x *Pipe) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.isNamed)
	m.Save(1, &x.atomicIOBytes)
	m.Save(2, &x.readers)
	m.Save(3, &x.writers)
	m.Save(4, &x.view)
	m.Save(5, &x.max)
	m.Save(6, &x.hadWriter)
}

func (x *Pipe) afterLoad() {}

func (x *Pipe) StateLoad(m state.Source) {
	m.Load(0, &x.isNamed)
	m.Load(1, &x.atomicIOBytes)
	m.Load(2, &x.readers)
	m.Load(3, &x.writers)
	m.Load(4, &x.view)
	m.Load(5, &x.max)
	m.Load(6, &x.hadWriter)
}

func (x *Reader) StateTypeName() string {
	return "pkg/sentry/kernel/pipe.Reader"
}

func (x *Reader) StateFields() []string {
	return []string{
		"ReaderWriter",
	}
}

func (x *Reader) beforeSave() {}

func (x *Reader) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.ReaderWriter)
}

func (x *Reader) afterLoad() {}

func (x *Reader) StateLoad(m state.Source) {
	m.Load(0, &x.ReaderWriter)
}

func (x *ReaderWriter) StateTypeName() string {
	return "pkg/sentry/kernel/pipe.ReaderWriter"
}

func (x *ReaderWriter) StateFields() []string {
	return []string{
		"Pipe",
	}
}

func (x *ReaderWriter) beforeSave() {}

func (x *ReaderWriter) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.Pipe)
}

func (x *ReaderWriter) afterLoad() {}

func (x *ReaderWriter) StateLoad(m state.Source) {
	m.Load(0, &x.Pipe)
}

func (x *Writer) StateTypeName() string {
	return "pkg/sentry/kernel/pipe.Writer"
}

func (x *Writer) StateFields() []string {
	return []string{
		"ReaderWriter",
	}
}

func (x *Writer) beforeSave() {}

func (x *Writer) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.ReaderWriter)
}

func (x *Writer) afterLoad() {}

func (x *Writer) StateLoad(m state.Source) {
	m.Load(0, &x.ReaderWriter)
}

func init() {
	state.Register((*inodeOperations)(nil))
	state.Register((*Pipe)(nil))
	state.Register((*Reader)(nil))
	state.Register((*ReaderWriter)(nil))
	state.Register((*Writer)(nil))
}
