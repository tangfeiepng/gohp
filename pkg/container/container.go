package container

import (
	"Walker/pkg/contract"
	"reflect"
)

type Container struct {
	singletons map[string]contract.MagicFunc
	alias      map[string]string
	instance   map[string]interface{}
}

func NewContainer() contract.Container {
	//初始化配置文件
	return &Container{
		//初始化
		singletons: make(map[string]contract.MagicFunc, 0),
		alias:      make(map[string]string, 0),
		instance:   make(map[string]interface{}, 0),
	}
}

func (container *Container) Bind(abstract string, concrete interface{}) {

}
func (container *Container) Singleton(abstract string, concrete interface{}) {
	//绑定关系
	//判断concrete是否时接口类型
	value := reflect.ValueOf(concrete)
	if value.Kind() != reflect.Func {
		panic("实现服务需要函数类型")
	}
	tem := container.CreateMagicFunc(concrete)
	//绑定关系
	container.singletons[abstract] = tem
	//设置绑定关系
	container.alias[tem.Outs()[0].String()] = abstract
}

//获取绑定关系的内容
func (container *Container) Get(key string) interface{} {
	if value, ok := container.instance[key]; ok {
		return value
	}
	if value, ok := container.singletons[key]; ok {
		result := container.MagicFuncCall(value)
		//不能每次都调用否则每次都初始化一次
		container.instance[key] = result
		return result
	}
	//判断alias集合当中是否存在如果存在直接返回
	if value, ok := container.alias[key]; ok {
		return container.Get(value)
	}
	return nil
}

func (container *Container) MagicFuncCall(magicFunc interface{}, bindArg ...interface{}) interface{} {
	if magicFunc, ok := magicFunc.(contract.MagicFunc); ok {
		InValues := make([]reflect.Value, 0)
		for _, v := range magicFunc.Ins() {
			//假如注入参数在额外绑定参数当中则无需去找该参数的实现
			InValues = append(InValues, container.judgeArg(v, bindArg))
		}
		if len(InValues) != magicFunc.InNums() {
			return nil
		}
		tem := magicFunc.Value().Call(InValues)
		if magicFunc.OutNums() <= 0 {
			return nil
		}
		return tem[0].Interface()

	}
	return container.MagicFuncCall(container.CreateMagicFunc(magicFunc))
}
func (container *Container) judgeArg(reType reflect.Type, bindArg []interface{}) reflect.Value {
	//判断该方法是否需要反射执行如果不需要则返回
	var res reflect.Value

	if result := container.Get(reType.String()); result != nil {
		return reflect.ValueOf(result)
	}

	//判断外部的参数类型转换后是否可以匹配上内部参数，或者外部参数执行后的返回值是否跟内部参数匹配上，或者不需要转换直接能匹配上内部的参数
	for _, arg := range bindArg {
		if reflect.TypeOf(arg).ConvertibleTo(reType) {
			res = reflect.ValueOf(arg)
			break
		}
		if reflect.TypeOf(arg).Kind() == reflect.Ptr {
			continue
		}
		tem := container.MagicFuncCall(arg)
		if reflect.TypeOf(tem).ConvertibleTo(reType) {
			res = reflect.ValueOf(tem)
			break
		}
	}
	if res.IsValid() == false {
		panic("没有找到该依赖的实现请检查是否提供服务注册")
	}
	return res
}
func (container *Container) Instance(abstract string, concrete interface{}) {
	//注册单例
	container.instance[abstract] = concrete
}

func (container *Container) SetPath(rootPath string) {
	container.instance["path"] = rootPath
	container.instance["path.storage"] = rootPath + "/storage"
	container.instance["path.public"] = rootPath + "/public"
}

//把函数反射成固定类型
func (container *Container) CreateMagicFunc(magicFunc interface{}) contract.MagicFunc {
	value := reflect.ValueOf(magicFunc)
	Type := reflect.TypeOf(magicFunc)
	//获取函数的传入惨与返回惨
	tem := MagicFunc{
		inNums:  Type.NumIn(),
		outNums: Type.NumOut(),
		value:   value,
		ins:     make([]reflect.Type, 0),
		outs:    make([]reflect.Type, 0),
	}
	//获取传入惨
	for i := 0; i < tem.inNums; i++ {
		tem.ins = append(tem.ins, Type.In(i))
	}
	//获取返回惨
	for i := 0; i < tem.outNums; i++ {
		tem.outs = append(tem.outs, Type.Out(i))
	}
	return &tem
}
