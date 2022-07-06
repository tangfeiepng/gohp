package router

import (
	"Walker/pkg/contract"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ServiceProvider struct {
	app    contract.Application
	Routes []interface{}
}

func (service *ServiceProvider) Register(app contract.Application) {
	//绑定关系
	service.app = app
	app.Singleton("router", func(config contract.Config) contract.Router {
		return &Router{
			app,
			gin.New(),
			config,
			make([]contract.Route, 0),
			make([]contract.GroupRouter, 0),
			make([]contract.MagicFunc, 0),
		}
	})
}
func (service *ServiceProvider) Boot() {
	//启动gin的http服务

	for _, value := range service.Routes {
		service.app.MagicFuncCall(value)
	}
	err := service.app.MagicFuncCall(func(router contract.Router) error {
		return router.Start()
	}).(error)
	if err != nil {
		fmt.Println("启动失败")
	} else {
		fmt.Println("启动成功")
	}
}
