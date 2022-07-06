package response

import (
	"Walker/pkg/contract"
	Response2 "Walker/pkg/response"
)

var Response contract.Response

func init() {
	Response = &Response2.Response{
		Status: 200,
	}
}
func Status(status int) contract.Response {
	return Response.SetStatus(status)
}
func Success(msg string, data interface{}) contract.Response {
	return Response.SetSuccess(msg, data)
}
func Error(msg string) contract.Response {
	return Response.SetError(msg)
}
