package response

import (
	"Walker/pkg/contract"
	"time"
)

type Response struct {
	Status int
	Data   dateResponse
}
type dateResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time string      `json:"time"`
}

func (i Response) SetStatus(status int) contract.Response {
	i.Status = status
	return i
}

func (i Response) SetError(msg string) contract.Response {
	i.Data = dateResponse{
		Code: 200,
		Msg:  msg,
		Data: nil,
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
	return i
}

func (i Response) SetSuccess(msg string, data interface{}) contract.Response {
	i.Data = dateResponse{
		Code: 400,
		Msg:  msg,
		Data: data,
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
	return i
}

func (i Response) ToJson() interface{} {
	return i.Data
}

func (i Response) GetStatus() int {
	return i.Status
}
