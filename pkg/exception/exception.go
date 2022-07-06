package exception

import (
	"Walker/pkg/contract"
	"Walker/pkg/router/http"
	"bytes"
	"encoding/json"
	"runtime"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

//异常处理工厂
type StackMap struct {
	File   string `json:"file"`
	Line   int    `json:"line"`
	Method string `json:"method"`
}

func Throw(exception contract.Exception) interface{} {
	//查询日志配置是否加载如果加载则写入日志文件
	return nil
}
func WithError(err interface{}) contract.Exception {
	stack := stack(3)
	var exception contract.Exception
	if e, isOk := err.(contract.Exception); !isOk {
		exception = CreateException(err)
	} else {
		exception = e
	}
	exception.HandleStack(stack)
	return exception
}
func CreateException(err interface{}) contract.Exception {
	return http.Exception{&Exception{
		Msg: err.(string),
	}}
}

type Exception struct {
	Msg   string `json:"msg"`
	File  string `json:"file"`
	Line  int    `json:"line"`
	Stack string `json:"stack"`
}

func (e *Exception) GetMessage() string {
	return e.Msg
}
func (e *Exception) GetLine() int {
	return e.Line
}
func (e *Exception) GetFile() string {
	return e.File
}
func (e *Exception) HandleStack(maps interface{}) {
	e.File = maps.([]StackMap)[0].File
	e.Line = maps.([]StackMap)[0].Line
	stackStr, _ := json.Marshal(maps)
	e.Stack = string(stackStr)
}
func (e *Exception) GetStack() string {
	return e.Stack
}
func stack(skip int) []StackMap {
	//buf := new(bytes.Buffer)
	var maps []StackMap
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		maps = append(maps, StackMap{file, line, string(function(pc))})
	}
	return maps
}
func function(pc uintptr) []byte {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return dunno
	}
	name := []byte(fn.Name())
	if lastSlash := bytes.LastIndex(name, slash); lastSlash >= 0 {
		name = name[lastSlash+1:]
	}
	if period := bytes.Index(name, dot); period >= 0 {
		name = name[period+1:]
	}
	name = bytes.Replace(name, centerDot, dot, -1)
	return name
}
