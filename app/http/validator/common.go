package validator

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
)
func DataAddContext(data interface{},ctx *gin.Context) *gin.Context {
	strRet, err := json.Marshal(data)
	if err != nil{
		return nil
	}
	// json转map
	var mRet map[string]interface{}
	err1 := json.Unmarshal(strRet, &mRet)
	if err1 != nil{
		return nil
	}
	for k, v := range mRet {
		ctx.Set(k,v)
	}
	return ctx
}


func Verify(ctx *gin.Context,Data interface{}) error {
	if err := ctx.ShouldBind(Data); err != nil {
		return err
	}
	//解析json数据绑定到上下文
	data :=DataAddContext(Data,ctx)
	if data == nil{
		return errors.New("表单数据处理失败")
	}
	return nil
}

