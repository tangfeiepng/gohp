package router

import (
	"Walker/pkg/contract"
	"Walker/pkg/pipeline"
	"Walker/pkg/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	app    contract.Application
	gin    *gin.Engine
	config contract.Config
	routes []contract.Route
	groups []contract.GroupRouter
	middle []contract.MagicFunc
}

//给路由分组

func (router *Router) Group(path string) contract.GroupRouter {
	group := (&GroupRouter{app: router.app}).NewGroup(path)
	//存放
	router.groups = append(router.groups, group)
	return group
}
func (router *Router) Get(url string, action interface{}, middleware ...interface{}) {
	router.Add(http.MethodGet, url, action, middleware...)
}
func (router *Router) Post(url string, action interface{}, middleware ...interface{}) {
	router.Add(http.MethodPost, url, action, middleware...)
}
func (router *Router) Put(url string, action interface{}, middleware ...interface{}) {
	router.Add(http.MethodPut, url, action, middleware...)
}
func (router *Router) Delete(url string, action interface{}, middleware ...interface{}) {
	router.Add(http.MethodDelete, url, action, middleware...)
}
func (router *Router) Use(middle ...interface{}) contract.Router {
	//组装中间件此处的为全局中间件和路由中间件
	for _, middleware := range middle {
		router.middle = append(router.middle, router.app.CreateMagicFunc(middleware))
	}
	return router
}
func (router *Router) Add(method string, url string, action interface{}, middleware ...interface{}) {
	router.routes = append(router.routes, &Route{
		method:     method,
		url:        url,
		action:     router.app.CreateMagicFunc(action),
		middleware: router.middlewareMagicFunc(middleware...),
	})
}
func (router *Router) middlewareMagicFunc(middleware ...interface{}) []contract.MagicFunc {
	middlewareList := make([]contract.MagicFunc, 0)
	for _, value := range middleware {
		middlewareList = append(middlewareList, router.app.CreateMagicFunc(value))
	}
	return middlewareList
}
func (router *Router) Start() error {
	var exceptionHandler = router.app.CreateMagicFunc(func(handler contract.ExceptionHandler, exception contract.Exception) interface{} {
		return handler.Handle(exception)
	})
	//把路由装填到gin服务当中启动路由
	//异常捕捉
	var Logger = router.app.Get("logger").(contract.Logger)
	router.gin.Use(WithWriter(Logger, func(exception contract.Exception, context *gin.Context) {
		err := router.app.MagicFuncCall(exceptionHandler, exception)
		req := (&request.Request{}).CreateRequest(context)
		if err != nil {
			router.HandlerResponse(req, err)
		} else {
			router.HandlerResponse(req, exception)
		}
	}))
	//router.gin.Use(func(context *gin.Context) {
	//	WithWriter()
	//})
	//router.gin.Use(gin.Recovery())
	middleware := router.middle
	router.FillingRoute(router.routes, middleware...)
	//分组的路由处理
	router.AnalyseGroup(router.groups, middleware...)

	return router.gin.Run(router.config.GetString("app.http.port"))
}

//组装gin路由

func (router *Router) FillingRoute(routers []contract.Route, middleware ...contract.MagicFunc) {
	for _, value := range routers {
		router.gin.Handle(value.Method(), value.Url(), router.HandlerFunc(value, middleware...))
	}
}

//把group里的路由进行组装

func (router *Router) AnalyseGroup(groups []contract.GroupRouter, middleware ...contract.MagicFunc) {
	for _, value := range groups {
		//该组的组中间件
		middleware := append(middleware, value.Middles()...)
		router.FillingRoute(value.Routes(), middleware...)
		if len(value.Groups()) > 0 {
			router.AnalyseGroup(value.Groups(), middleware...)
		}
	}
}
func (router *Router) HandlerFunc(route contract.Route, middleware ...contract.MagicFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		req := (&request.Request{}).CreateRequest(context)
		middleware := append(middleware, route.Middleware()...)
		result := (&pipeline.Pipeline{}).
			CreatePipe(router.app).
			Send(req).
			Through(middleware...).
			Then(route.Action())
		//处理响应结构
		router.HandlerResponse(req, result)
	}
}

func (router *Router) HandlerResponse(request contract.Request, response interface{}) {

	switch tem := response.(type) {
	case contract.Response:
		//标准返回
		request.Context().(*gin.Context).JSON(tem.GetStatus(), tem.ToJson())
	case contract.Exception:
		request.Context().(*gin.Context).JSON(http.StatusInternalServerError, tem)
	default:
		request.Context().(*gin.Context).JSON(200, response)
	}
}
