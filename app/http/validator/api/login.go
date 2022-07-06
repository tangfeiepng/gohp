package api

type RegisterRequest struct {
	OpenId string `binding:"required" form:"openid" json:"openid"`
	Iv string `binding:"required" form:"iv" json:"iv"`
	EncryptedData string `binding:"required" form:"encryptedData" json:"encryptedData"`
}
type LoginRequest struct {
	Code string `json:"code" binding:"required" form:"code"`
}