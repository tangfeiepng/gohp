package param

import "Walker/pkg/contract"

type ServiceProvider struct {
}

func (service *ServiceProvider) Register(app contract.Application) {
	//绑定关系
	app.Singleton("param", func() contract.Param {
		return &Param{
			oneParam: 10,
			twoParam: 20,
		}
	})
}
func (service *ServiceProvider) Boot() {

}
