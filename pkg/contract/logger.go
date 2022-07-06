package contract

type LoggerField struct {
	Key string
	Val interface{}
}
type Logger interface {
	//debug
	Debug(msg string, fields ...LoggerField)
	Info(msg string, fields ...LoggerField)
	Warn(msg string, fields ...LoggerField)
	Error(msg string, fields ...LoggerField)
}
