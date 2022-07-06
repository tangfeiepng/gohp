package contract

type Request interface {
	Param(key string) string
	Query(key string) string
	Next()
	Context() interface{}
}
