package contract

type Pipeline interface {
	//把请求传进来
	Send(passable interface{}) Pipeline
	//设置所有的中间件
	Through(pipes ...MagicFunc) Pipeline
	//运行函数THEN
	Then(action interface{}) interface{}
}

//passable就是request业务处理

type Pipe func(passable interface{}) interface{}
