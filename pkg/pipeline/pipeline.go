package pipeline

import (
	"Walker/pkg/contract"
)

type Pipeline struct {
	app      contract.Application
	pipes    []contract.MagicFunc
	passable interface{}
}
type Callback func(stack contract.Pipe, next contract.MagicFunc) contract.Pipe

func (p *Pipeline) CreatePipe(app contract.Application) contract.Pipeline {
	p.app = app
	return p
}
func (p *Pipeline) Send(passable interface{}) contract.Pipeline {
	p.passable = passable
	return p
}
func (p *Pipeline) Through(pipes ...contract.MagicFunc) contract.Pipeline {
	p.pipes = append(p.pipes, pipes...)
	return p
}

func (p *Pipeline) Then(destination interface{}) interface{} {
	return p.arrayReduce(p.arrayReverse(p.pipes), p.carry(), p.prepareDestination(destination))(p.passable)
}

//中间件进行反转
func (p *Pipeline) arrayReverse(pipes []contract.MagicFunc) []contract.MagicFunc {
	for from, to := 0, len(pipes)-1; from < to; from, to = from+1, to-1 {
		pipes[from], pipes[to] = pipes[to], pipes[from]
	}
	return pipes
}
func (p *Pipeline) arrayReduce(pipes []contract.MagicFunc, Callback Callback, destination contract.Pipe) contract.Pipe {
	for _, magicFunc := range pipes {
		destination = Callback(destination, magicFunc)
	}
	return destination
}
func (p *Pipeline) carry() Callback {
	return func(stack contract.Pipe, next contract.MagicFunc) contract.Pipe {
		return func(passable interface{}) interface{} {
			return p.app.MagicFuncCall(next, func() interface{} {
				return passable
			}, stack)
		}
	}
}
func (p *Pipeline) prepareDestination(destination interface{}) contract.Pipe {
	return func(passable interface{}) interface{} {
		return p.app.MagicFuncCall(destination, passable)
	}
}
