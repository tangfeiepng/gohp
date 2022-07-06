package contract

type Router interface {
	Group(path string) GroupRouter
	Get(url string, action interface{}, middleware ...interface{})
	Post(url string, action interface{}, middleware ...interface{})
	Put(url string, action interface{}, middleware ...interface{})
	Delete(url string, action interface{}, middleware ...interface{})
	Use(middle ...interface{}) Router
	Start() error
}

type Route interface {
	Method() string
	Url() string
	Action() MagicFunc
	Middleware() []MagicFunc
}
type GroupRouter interface {
	Group(path string) GroupRouter
	Get(url string, action interface{}, middleware ...interface{})
	Post(url string, action interface{}, middleware ...interface{})
	Put(url string, action interface{}, middleware ...interface{})
	Delete(url string, action interface{}, middleware ...interface{})
	Use(middle ...interface{}) GroupRouter
	Routes() []Route
	Middles() []MagicFunc
	Groups() []GroupRouter
}
