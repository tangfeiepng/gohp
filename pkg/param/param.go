package param

type Param struct {
	oneParam int
	twoParam int
}

func (param *Param) OneParam() int {
	return param.oneParam
}
func (param *Param) TwoParam() int {
	return param.twoParam
}
