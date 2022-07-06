package contract

import "reflect"

type Container interface {
	Bind(abstract string, concrete interface{})
	Singleton(abstract string, concrete interface{})
	Get(key string) interface{}
	SetPath(rootPath string)
	MagicFuncCall(magicFunc interface{}, bindArg ...interface{}) interface{}
	CreateMagicFunc(magicFunc interface{}) MagicFunc
}
type MagicFunc interface {
	Ins() []reflect.Type
	Outs() []reflect.Type
	InNums() int
	OutNums() int
	Value() reflect.Value
}
