package main

import (
	application2 "Walker/pkg/application"
	"Walker/pkg/config"
	"Walker/pkg/console"
	"Walker/pkg/contract"
	"Walker/pkg/example"
	"Walker/pkg/exception"
	"Walker/pkg/logger"
	"Walker/pkg/param"
	"Walker/pkg/router"
	"Walker/routers"
	"os"
)

func main() {
	//初始化一个ioc容器实例
	application := application2.App()
	//设置项目根目录
	rootPath, _ := os.Getwd()
	application.SetPath(rootPath)
	application.ServiceProviders(
		&exception.ServiceProvider{},
		&config.ServiceProvider{},
		&logger.ServiceProvider{},
		&router.ServiceProvider{Routes: []interface{}{routers.Test}},
		&console.ServiceProvider{},
		&param.ServiceProvider{},
		&example.ServiceProvider{},
	)

	application.MagicFuncCall(func(console contract.Console, config2 contract.Config) {
		//获取参数启动服务提供者的boot方法
		console.Run()
	})
}
