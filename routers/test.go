package routers

import (
	"Walker/app/http/controller/api"
	"Walker/app/http/middleware"
	"Walker/pkg/contract"
)

func Test(router contract.Router) {
	Test := router.Group("/test").Use(middleware.GroupHandle)
	{
		Test.Get("/test", (&api.Demo{}).Test, middleware.Handle)
	}
	router.Get("/test_two", (&api.Demo{}).TestTwo)
	TestOne := router.Group("/test_one")
	{
		TestOne.Get("/test_one", (&api.Demo{}).TestOne)
	}

}
