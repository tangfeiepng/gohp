package wechat

import (
	data2 "Walker/app/utils/data"
	"Walker/global"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//微信相关服务
var
(
	Code2SessionPath string = "https://api.weixin.qq.com/sns/jscode2session"
)
func Code2Session(code string,data *Code2SessionData) error {

	url := Code2SessionPath+"?appid="+global.ConfigValue.GetString("Wechat.Appid")+"&secret=" + global.ConfigValue.GetString("Wechat.Secret") + "&js_code=" + code+ "&grant_type=authorization_code"
	resp,err := http.Get(url)
	if err!=nil{
		return err
	}
	res,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		return err
	}
	//字节转对象
	err=json.Unmarshal([]byte(string(res)),data)
	if err!=nil{
		return err
	}
	if data.ErrCode!=0{
		return errors.New(data.ErrMsg)
	}
	return nil
}
/**
 * error code 说明.
 * <ul>

 *    <li>-41001: encodingAesKey 非法</li>
 *    <li>-41003: aes 解密失败</li>
 *    <li>-41004: 解密后得到的buffer非法</li>
 *    <li>-41005: base64加密失败</li>
 *    <li>-41016: base64解密失败</li>
 * </ul>
 */

func DecryptData(encryptedData string, iv string ,sessionKey string,data interface{}) error {
	WxBizDataCrypt :=WxBizDataCrypt{
		AppId: global.ConfigValue.GetString("Wechat.Appid"),
		SessionKey: sessionKey,
	}
	result,err:=WxBizDataCrypt.Decrypt(encryptedData,iv,false)
	if err!=nil{
		return err
	}
	err=data2.MapToStruct(result,data)
	if err!=nil{
		return err
	}
	return nil
}