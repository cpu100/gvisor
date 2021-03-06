// automatically generated by stateify.

package bpf

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *Program) StateTypeName() string {
	return "pkg/bpf.Program"
}

func (x *Program) StateFields() []string {
	return []string{
		"instructions",
	}
}

func (x *Program) beforeSave() {}

func (x *Program) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.instructions)
}

func (x *Program) afterLoad() {}

func (x *Program) StateLoad(m state.Source) {
	m.Load(0, &x.instructions)
}

func init() {
	state.Register((*Program)(nil))
}
