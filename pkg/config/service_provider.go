package config

import (
	"Walker/pkg/contract"
	"fmt"
	"github.com/fsnotify/fsnotify"
	viper2 "github.com/spf13/viper"
)

type ServiceProvider struct {
	app contract.Application
}

func (service *ServiceProvider) Register(app contract.Application) {
	//绑定关系
	service.app = app
	app.Singleton("config", func() contract.Config {
		//获取内容
		rootPath := app.Get("path").(string)
		viper := viper2.New()
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
		viper.AddConfigPath(rootPath)
		if err := viper.ReadInConfig(); err != nil {
			panic("无法读取到配置文件")
			//exception.Throw(&Exception{Msg: "无法读取到配置文件"})
		}
		return &Config{viper}
	})
}
func (service *ServiceProvider) Boot() {
	//启动配置文件监听服务
	service.app.MagicFuncCall(func(config contract.Config) {
		config.GetViper().OnConfigChange(func(in fsnotify.Event) {
			fmt.Println("参数正在改变", in.Name)
		})
		config.GetViper().WatchConfig()
	})
}
