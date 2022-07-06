package router

import (
	"Walker/pkg/contract"
	"net/http"
)

type GroupRouter struct {
	path   string
	groups []contract.GroupRouter
	routes []contract.Route
	app    contract.Application
	middle []contract.MagicFunc
}

//创建新的组
func (group *GroupRouter) NewGroup(path string) contract.GroupRouter {
	return &GroupRouter{
		path:   group.path + path,
		app:    group.app,
		routes: make([]contract.Route, 0),
		groups: make([]contract.GroupRouter, 0),
		middle: make([]contract.MagicFunc, 0),
	}
}

//给路由分组
func (group *GroupRouter) Group(path string) contract.GroupRouter {
	groupItem := group.NewGroup(path)
	group.groups = append(group.groups, groupItem)
	return groupItem
}
func (group *GroupRouter) Get(url string, action interface{}, middleware ...interface{}) {
	group.Add(http.MethodGet, url, action, middleware...)
}
func (group *GroupRouter) Post(url string, action interface{}, middleware ...interface{}) {
	group.Add(http.MethodPost, url, action, middleware...)
}
func (group *GroupRouter) Put(url string, action interface{}, middleware ...interface{}) {
	group.Add(http.MethodPut, url, action, middleware...)
}
func (group *GroupRouter) Delete(url string, action interface{}, middleware ...interface{}) {
	group.Add(http.MethodDelete, url, action, middleware...)
}
func (group *GroupRouter) Use(middle ...interface{}) contract.GroupRouter {
	//此处为组中间件和路由中间件
	for _, value := range middle {
		group.middle = append(group.middle, group.app.CreateMagicFunc(value))
	}
	return group
}
func (group *GroupRouter) Add(method string, url string, action interface{}, middleware ...interface{}) {
	group.routes = append(group.routes, &Route{
		method:     method,
		url:        group.path + url,
		action:     group.app.CreateMagicFunc(action),
		middleware: group.middlewareMagicFunc(middleware...),
	})
}
func (group *GroupRouter) middlewareMagicFunc(middleware ...interface{}) []contract.MagicFunc {
	middlewareList := make([]contract.MagicFunc, 0)
	for _, value := range middleware {
		middlewareList = append(middlewareList, group.app.CreateMagicFunc(value))
	}
	return middlewareList
}
func (group *GroupRouter) Routes() []contract.Route {
	return group.routes
}
func (group *GroupRouter) Groups() []contract.GroupRouter {
	return group.groups
}
func (group *GroupRouter) Middles() []contract.MagicFunc {
	return group.middle
}
