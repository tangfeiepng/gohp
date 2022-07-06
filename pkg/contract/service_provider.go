package contract

type ServiceProvider interface {
	Register(application Application)
	Boot()
}
