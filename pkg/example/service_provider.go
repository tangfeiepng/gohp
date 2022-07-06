package example

import "Walker/pkg/contract"

type ServiceProvider struct {
}

func (service *ServiceProvider) Register(app contract.Application) {
	//绑定关系
	app.Singleton("demo", func(param contract.Param) contract.Demo {
		return &Demo{Param: param}
	})
}
func (service *ServiceProvider) Boot() {

}
