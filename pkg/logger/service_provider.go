package logger

import "Walker/pkg/contract"

type ServiceProvider struct {
}

func (service *ServiceProvider) Register(app contract.Application) {
	//绑定关系
	app.Singleton("logger", func(config contract.Config) contract.Logger {
		return (&ZapClass{config: config, app: app}).InitLogger()
	})
}
func (service *ServiceProvider) Boot() {

}
