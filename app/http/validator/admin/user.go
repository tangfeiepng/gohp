package admin

type UserRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password"`
	RoleId   int    `json:"role_id" binding:"required"`
	Status   uint   `json:"status"`
}
