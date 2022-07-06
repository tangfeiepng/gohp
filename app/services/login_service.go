package services

import (
	"Walker/app/http/model"
	"Walker/app/http/validator/admin"
	"Walker/app/http/validator/api"
	"Walker/app/utils"
	"Walker/app/utils/jwt"
	"Walker/app/utils/wechat"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Login struct {
	UserWechatModel model.UserWechat
	Code2SessionData wechat.Code2SessionData
	UserModel model.User
}

func (login *Login) Register(param api.RegisterRequest) (string,error) {
	//查询sessionkey
	result :=model.UseCon().Where("openid = ?",param.OpenId).First(&login.UserWechatModel)
	if result.RowsAffected==0{
		return "",errors.New("用户信息不存在")
	}
	var UserPhoneData wechat.UserPhoneData
	err:=wechat.DecryptData(param.EncryptedData,param.Iv,login.UserWechatModel.SessionKey,&UserPhoneData)
	if err!=nil{
		return "",err
	}
	UserId := login.UserWechatModel.UserId
	if login.UserWechatModel.UserId==0{
		login.UserModel = model.User{
			Ip: utils.GetLocalIP(),
			Phone: UserPhoneData.PhoneNumber,
			LoginTime: time.Now(),
		}
		result=model.UseCon().Create(&login.UserModel)
		if result.Error != nil{
			return "",err
		}
		login.UserWechatModel.UserId =  login.UserModel.Id
		model.UseCon().Where("id=?",login.UserWechatModel.Id).Updates(&login.UserWechatModel)
		UserId = login.UserModel.Id
	}
	//注册用户账号
	token,err:=login.UserLogin(UserId)
	if err!=nil{
		return "",err
	}
	return token,nil
}
func (login *Login) Login(param api.LoginRequest) (interface{},error) {
	login.Code2SessionData = wechat.Code2SessionData{}
	err:=wechat.Code2Session(param.Code,&login.Code2SessionData)
	if err!=nil{
		return nil,err
	}
	//处理数据
	login.UserWechatModel = model.UserWechat{}
	result  :=model.UseCon().Where("openid = ?",login.Code2SessionData.Openid).First(&login.UserWechatModel)
	err=login.editUserWechat(result.RowsAffected)
	if err!=nil{
		return nil,err
	}
	res :=map[string]string{
		"openid":"",
		"type":"success",
		"token":"",
	}
	if login.UserWechatModel.UserId==0{
		res =map[string]string{
			"openid":login.Code2SessionData.Openid,
			"type":"register",
		}
	}else{
		token,err:=login.UserLogin(login.UserWechatModel.UserId)
		if err!=nil{
			return nil,err
		}
		res["token"] = token
	}
	return res,nil
}

func (login *Login) editUserWechat(num int64) error {
	var result *gorm.DB
	if num==0{
		login.UserWechatModel = model.UserWechat{
			Openid: login.Code2SessionData.Openid,
			UnionId: login.Code2SessionData.UnionId,
			SessionKey: login.Code2SessionData.SessionKey,
		}
		result = model.UseCon().Create(&login.UserWechatModel)
	}else{
		login.UserWechatModel.SessionKey = login.Code2SessionData.SessionKey
		result = model.UseCon().Where("id=?",login.UserWechatModel.Id).Updates(&login.UserWechatModel)
	}
	if result !=nil{
		return result.Error
	}
	return nil

}

func (login *Login) UserLogin(UserId uint) (string,error)  {
	token,err:=jwt.GenerateToken(UserId)
	if err!=nil{
		return "",err
	}
	//更新用户登录时间更换code
	login.UserModel = model.User{
		Ip: utils.GetLocalIP(),
		LoginTime: time.Now(),
	}
	model.UseCon().Where("id=?",UserId).Updates(&login.UserModel)
	return token,nil
}

func (login *Login) AdminLogin(loginRequest admin.LoginRequest) (string,error) {
	var adminUserModel model.AdminUser
	result:=model.UseCon().Where("username=?",loginRequest.UserName).First(&adminUserModel)
	if result.RowsAffected<=0{
		return "",result.Error
	}
	err:=bcrypt.CompareHashAndPassword([]byte(adminUserModel.Password),[]byte(loginRequest.PassWord))
	if err!=nil{
		return "",err
	}
	token,err:=jwt.GenerateToken(adminUserModel.Id)
	if err!=nil{
		return "",err
	}
	//更新用户登录时间更换code
	adminUserModel.LoginTime = time.Now()
	adminUserModel.Ip = utils.GetLocalIP()
	model.UseCon().Where("id=?",adminUserModel.Id).Updates(adminUserModel)
	return token,nil
}
