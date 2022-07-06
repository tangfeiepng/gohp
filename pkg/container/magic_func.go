package container

import (
	"reflect"
)

type MagicFunc struct {
	ins     []reflect.Type
	outs    []reflect.Type
	inNums  int
	outNums int
	value   reflect.Value
}

func (m *MagicFunc) Ins() []reflect.Type {
	return m.ins
}
func (m *MagicFunc) Outs() []reflect.Type {
	return m.outs
}
func (m *MagicFunc) InNums() int {
	return m.inNums
}
func (m *MagicFunc) OutNums() int {
	return m.outNums
}
func (m *MagicFunc) Value() reflect.Value {
	return m.value
}
