package exception

//把异常处理成标准异常

type NormException struct {
}

//func (norm *NormException) Throw(err interface{}) {
//	if err != nil {
//		panic(norm.judgeException(err))
//	}
//}
//
//func (norm *NormException) judgeException(err interface{}) contract.Exception {
//	switch result := err.(type) {
//	case error:
//
//	}
//}
