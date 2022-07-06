package services

import (
	"Walker/app/http/model"
	"Walker/app/http/validator/admin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SystemFunc struct {
	Model model.SystemFunc
}

func (service *SystemFunc) Edit(param admin.SystemFuncRequest, id int) error {
	data := model.SystemFunc{
		Name:   param.Name,
		Desc:   param.Desc,
		Status: param.Status,
		Url:    param.Url,
		Method: param.Method,
	}
	var result *gorm.DB
	if id == 0 {
		result = model.UseCon().Create(&data)
	} else {
		result = model.UseCon().Where("id=?", id).Updates(&data)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (service *SystemFunc) Index(ctx *gin.Context) (model.SystemFuncPage, error) {
	var list []model.SystemFunc
	db := model.UseCon()
	if str := ctx.Query("keywords"); str != "" {
		db = db.Where("`name` like ? or `desc` like ?", "%"+str+"%", "%"+str+"%")
	}
	if str := ctx.Query("status"); str != "" {
		db = db.Where("status = ?", str)
	}
	var count int64
	db.Model(model.SystemFunc{}).Count(&count)
	result := db.Scopes(model.Paginate(ctx)).Find(&list)
	if result.Error != nil {
		return model.SystemFuncPage{}, result.Error
	}
	return model.SystemFuncPage{
		Data:      list,
		TotalRows: count,
	}, nil
}
