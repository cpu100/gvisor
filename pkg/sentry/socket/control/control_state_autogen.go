// automatically generated by stateify.

package control

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *RightsFiles) StateTypeName() string {
	return "pkg/sentry/socket/control.RightsFiles"
}

func (x *RightsFiles) StateFields() []string {
	return nil
}

func (x *scmCredentials) StateTypeName() string {
	return "pkg/sentry/socket/control.scmCredentials"
}

func (x *scmCredentials) StateFields() []string {
	return []string{
		"t",
		"kuid",
		"kgid",
	}
}

func (x *scmCredentials) beforeSave() {}

func (x *scmCredentials) StateSave(m state.Sink) {
	x.beforeSave()
	m.Save(0, &x.t)
	m.Save(1, &x.kuid)
	m.Save(2, &x.kgid)
}

func (x *scmCredentials) afterLoad() {}

func (x *scmCredentials) StateLoad(m state.Source) {
	m.Load(0, &x.t)
	m.Load(1, &x.kuid)
	m.Load(2, &x.kgid)
}

func init() {
	state.Register((*RightsFiles)(nil))
	state.Register((*scmCredentials)(nil))
}
