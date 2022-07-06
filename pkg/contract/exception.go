package contract

//定义异常

type Exception interface {
	GetMessage() string
	GetLine() int
	GetFile() string
	HandleStack(interface{})
	GetStack() string
}

//定义异常处理者

type ExceptionHandler interface {
	Handle(e Exception) interface{}
}
