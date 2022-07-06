package exception

import (
	"Walker/app/exceptions"
	"Walker/pkg/contract"
)

type ServiceProvider struct {
}

func (service *ServiceProvider) Register(app contract.Application) {
	//绑定关系
	app.Singleton("exception.handler", func() contract.ExceptionHandler {
		return &exceptions.ExceptionHandler{}
	})
}
func (service *ServiceProvider) Boot() {

}
