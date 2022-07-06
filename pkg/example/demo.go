package example

import (
	"Walker/pkg/contract"
	"fmt"
	"strconv"
)

type Demo struct {
	Param contract.Param
}

func (demo *Demo) Plus() {
	tem := demo.Param.OneParam() + demo.Param.TwoParam()
	fmt.Println("加法等于13" + strconv.Itoa(tem))
}
