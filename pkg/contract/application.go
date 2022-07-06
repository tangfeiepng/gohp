package contract

type Application interface {
	Container
	ServiceProviders(service ...ServiceProvider)
	Boot()
}
