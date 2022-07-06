# gohp
一个继承了 laravel 思想的 golang web 框架

## 框架特点
goal 通过容器和服务提供者作为框架的核心，以 contracts 为桥梁，为开发者提供丰富的功能和服务，这点与 laravel 是相似的。
* 强大的容器
* 服务提供者
* 契约精神

## 链接（项目初始借鉴对象实现的是同一个理念）
* [goal 仓库](https://github.com/goal-web/goal)

## 功能特性
* [x] docker-compose 环境一键运行
* [x] contracts 定义模块接口
* [x] container 容器实现！！！
* [x] pipeline 简单但是很强大的洋葱模型的管道
* [x] application 应用
    * [x] exceptions 异常处理模块
* [x] config 配置模块
* [x] logger 配置模块
* [x] http http相关模块，请求、响应、中间件等
    * [x] routing http 路由服务
* [x] console 命令行模块

