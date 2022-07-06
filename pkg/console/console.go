package console

import (
	"Walker/pkg/contract"
)

type Console struct {
	app contract.Application
}

func (console *Console) Run() {
	//启动所有的boot
	console.app.Boot()
}
