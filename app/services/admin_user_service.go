package services

import (
	"Walker/app/http/model"
	"Walker/app/http/validator/admin"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminUserService struct {
}

func (service *AdminUserService) Index(ctx *gin.Context) (model.AdminUserPage, error) {
	var data []model.AdminUser
	db := model.UseCon()
	if str := ctx.Query("keywords"); str != "" {
		db = db.Where("`username` like ? ", "%"+str+"%")
	}
	if str := ctx.Query("status"); str != "" {
		db = db.Where("status = ?", str)
	}
	var count int64
	db.Scopes(model.Paginate(ctx)).Count(&count)
	result := db.Find(&data)
	if result.Error != nil {
		return model.AdminUserPage{}, result.Error
	}
	return model.AdminUserPage{
		Data:      data,
		TotalRows: count,
	}, nil
}

func (service *AdminUserService) Edit(id uint, param admin.UserRequest) error {
	if param.Password == "" && id != 0 {
		return errors.New("添加管理员密码不能为空")
	}
	data := model.AdminUser{
		Username: param.UserName,
		RoleId:   param.RoleId,
		Status:   param.Status,
	}
	if param.Password != "" {
		if hash, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.MinCost); err != nil {
			return err
		} else {
			data.Password = string(hash)
		}
	}
	var result *gorm.DB
	if id != 0 {
		result = model.UseCon().Where("id=?", id).Updates(&data)
	} else {
		result = model.UseCon().Create(&data)
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}
