package console

import (
	"Walker/pkg/contract"
)

type ServiceProvider struct {
	app contract.Application
}

func (service *ServiceProvider) Register(app contract.Application) {
	//绑定关系
	service.app = app
	app.Singleton("console", func() contract.Console {
		return &Console{app}
	})
}
func (service *ServiceProvider) Boot() {
	//打印一些服务
	service.app.MagicFuncCall(func(config contract.Config) {
		println(config.GetString("app.name"))
	})
}
