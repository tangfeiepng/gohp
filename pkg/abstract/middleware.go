package abstract

import "Walker/pkg/contract"

type Middleware interface {
	Handle(request contract.Request, next contract.Pipe) interface{}
}
