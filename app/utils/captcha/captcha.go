package captcha

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   base64Captcha.DriverAudio
	DriverString  base64Captcha.DriverString
	DriverChinese base64Captcha.DriverChinese
	DriverMath    base64Captcha.DriverMath
	DriverDigit   base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// CreateCaptcha
// 生成验证码
func CreateCaptcha() (string, string) {
	var param configJsonBody
	param.CaptchaType = "string"
	param.DriverString.Source = "1234567890"
	param.DriverString.Length = 4
	param.DriverString.Width = 80
	param.DriverString.Height = 30
	param.DriverString.NoiseCount=50
	param.DriverString.ShowLineOptions = 0|4|8
	// 干扰线
	//param.DriverString.ShowLineOptions = 10
	// 3Dumb.ttf wqy-microhei.ttc
	param.DriverString.Fonts = []string{"RitaSmith.ttf", "DENNEthree-dee.ttf", "wqy-microhei.ttc"}
	param.DriverString.NoiseCount = 0
	param.DriverString.BgColor = &color.RGBA{
		R: 3,
		G: 102,
		B: 214,
		A: 125,
	}
	var driver base64Captcha.Driver
	driver = param.DriverString.ConvertFonts()

	var store = base64Captcha.DefaultMemStore
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _ := c.Generate()
	//b64s = b64s[22:]
	return b64s, id
}

// VerifyCaptcha
// 验证验证码
func VerifyCaptcha(id string, verifyValue string) bool {
	res := store.Verify(id, verifyValue, false)
	if res {
		clear := store.Verify(id, verifyValue, true)
		return clear
	}
	return false
}