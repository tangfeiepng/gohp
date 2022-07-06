package application

import (
	"Walker/pkg/container"
	"Walker/pkg/contract"
)

var Instance contract.Application

func App() contract.Application {
	if Instance != nil {
		return Instance
	}
	return &Application{
		container.NewContainer(),
		make([]contract.ServiceProvider, 0),
	}
}
