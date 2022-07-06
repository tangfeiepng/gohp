package application

import (
	"Walker/pkg/contract"
)

type Application struct {
	contract.Container
	Services []contract.ServiceProvider
}

func (app *Application) ServiceProviders(service ...contract.ServiceProvider) {
	app.Services = append(app.Services, service...)
	for _, value := range service {
		value.Register(app)
	}
}
func (app *Application) Boot() {
	//启动所有的服务
	for _, value := range app.Services {
		value.Boot()
	}
}
