package api

import (
	"Walker/pkg/contract"
	"Walker/pkg/supports/response"
	"fmt"
)

type Demo struct {
}

func (s *Demo) Test(request contract.Request, logger contract.Logger) string {
	//测试日志文件
	//logger.Error("测试", contract.LoggerField{
	//	Key: "扩展字段",
	//	Val: "测试",
	//})
	panic("sssss")
	return "我返回的是字符串"
}
func (s *Demo) TestOne(request contract.Request) contract.Response {
	return response.Success("返回成功", request.Query("msg"))
}

func (s *Demo) TestTwo() {
	slice := []int{1, 2, 3, 4, 5}
	//slice = append(slice[:len(slice)-1]) //删除第一个元素
	//fmt.Println(slice)
	//进行反转
	for from, to := 0, len(slice)-1; from < to; from, to = from+1, to-1 {
		slice[from], slice[to] = slice[to], slice[from]
	}
	fmt.Println(slice)
}
