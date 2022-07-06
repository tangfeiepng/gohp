package exceptions

import (
	"Walker/pkg/contract"
	"Walker/pkg/router/http"
	"fmt"
)

type ExceptionHandler struct {
}

func (handler *ExceptionHandler) Handle(exception contract.Exception) interface{} {
	switch exception.(type) {
	case http.Exception:
		fmt.Println("我收到类型了")
	}
	fmt.Println("我收到信息了" + exception.GetMessage())
	return nil
}
