package admin

type SystemFuncRequest struct {
	Name string `json:"name" binding:"required"`
	Desc string `json:"desc"`
	Status uint `json:"status" binding:"required"`
	Url string `json:"url" binding:"required"`
	Method string `json:"method" binding:"required"`
}