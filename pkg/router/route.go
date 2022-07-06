package router

import "Walker/pkg/contract"

type Route struct {
	method     string
	url        string
	action     contract.MagicFunc
	middleware []contract.MagicFunc
}

func (route *Route) Method() string {
	return route.method
}
func (route *Route) Url() string {
	return route.url
}
func (route *Route) Action() contract.MagicFunc {
	return route.action
}

func (route *Route) Middleware() []contract.MagicFunc {
	return route.middleware
}
