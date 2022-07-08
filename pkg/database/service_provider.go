package database

import "Walker/pkg/contract"

type ServiceProvider struct {
}

func (service *ServiceProvider) Register(app contract.Application) {
	//绑定关系
	app.Singleton("model", func(config contract.Config) contract.Model {
		return (&Model{}).DBConnector(config)
	})
}
func (service *ServiceProvider) Boot() {

}
