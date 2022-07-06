package admin

type LoginRequest struct {
	UserName string `json:"username" binding:"required" form:"username"`
	PassWord string `json:"password" binding:"required" form:"password"`
}
