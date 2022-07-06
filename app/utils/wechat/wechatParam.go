package wechat

type Code2SessionData struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
	UnionId string `json:"unionid"`
	SessionKey string `json:"session_key"`
	Openid string `json:"openid"`
}

type UserProfileData struct {
	NickName string
	Gender int
	Language string
	City string
	Province string
	Country string
	AvatarUrl string
	Watermark struct{
		Timestamp int64
		Appid string
	}
}

type UserPhoneData struct {
	PhoneNumber string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode string `json:"countryCode"`
	Watermark struct{
		Timestamp uint64 `json:"timestamp"`
		Appid string `json:"appid"`
	}
}