package middleware

import (
	"Walker/pkg/contract"
)

type Demo struct {
}

func Handle(request contract.Request, next contract.Pipe) interface{} {
	//模拟中间件响应
	//fmt.Println("我是路由中间件业务处理前")
	response := next(request)
	//fmt.Println("我是路由中间件业务处理后")
	return response
}
func GlobalHandle(request contract.Request, next contract.Pipe) interface{} {
	//fmt.Println("我是全局中间件业务处理前")
	response := next(request)
	//fmt.Println("我是全局中间件业务处理后")
	return response
}
func GroupHandle(request contract.Request, next contract.Pipe) interface{} {
	//fmt.Println("我是组中间件业务处理前")
	response := next(request)
	//fmt.Println("我是组中间件业务处理后")
	return response
}
