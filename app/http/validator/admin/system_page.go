package admin

type SystemPage struct {
	Name string `json:"name" binding:"required"`
	Url string `json:"url"`
	Pid uint `json:"pid"`
	Icon string `json:"icon"`
	Status uint `json:"status" binding:"min=0,max=1"`
	PageFunc []int `json:"page_func"`
}