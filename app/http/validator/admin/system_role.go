package admin

type SystemRoleRequest struct {
	Name string `json:"name" binding:"required"`
	Pid uint `json:"pid"`
	Status uint `json:"status"`
	RoleData struct{
		PageIds []int `json:"page_ids"`
		FuncIds []int `json:"func_ids"`
	} `json:"role_data"`
}
