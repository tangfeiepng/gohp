package api

import (
	"Walker/pkg/contract"
	"Walker/pkg/supports/response"
	"fmt"
)

type Demo struct {
}
type Money struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Moneys float32 `json:"moneys"`
}

func (s *Demo) Test(model contract.Model) interface{} {
	//测试日志文件
	var result []map[string]interface{}
	model.Model(Money{}).Where("id=?", 1).Find(&result)
	return result
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
