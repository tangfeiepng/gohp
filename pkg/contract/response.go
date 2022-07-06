package contract

type Response interface {
	SetStatus(status int) Response
	SetError(msg string) Response
	SetSuccess(msg string, data interface{}) Response
	ToJson() interface{}
	GetStatus() int
}
