package services

import (
	"Walker/app/http/model"
	"Walker/app/http/validator/admin"
	"Walker/app/utils/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SystemPage struct {
	SystemPageListModel []model.SystemPage
}

func (services *SystemPage) List(ctx *gin.Context) (model.SystemPagePage, error) {
	//判断是否有条件输入
	list := model.UseCon()
	if str := ctx.DefaultQuery("keywords", ""); str != "" {
		list = list.Where("name like ?", "%"+str+"%")
	}
	if str := data.StringToInt(ctx.DefaultQuery("status", "-1")); str >= 0 {
		list = list.Where("status = ?", str)
	}
	var count int64
	list.Model(model.SystemPage{}).Count(&count)
	result := list.Scopes(model.Paginate(ctx)).Find(&services.SystemPageListModel)
	if result.Error != nil {
		return model.SystemPagePage{}, result.Error
	}
	return model.SystemPagePage{
		Data:      services.SystemPageListModel,
		TotalRows: count,
	}, nil
}

func (services *SystemPage) Edit(param admin.SystemPage, id uint) error {
	arr := model.SystemPage{
		Name:   param.Name,
		Icon:   param.Icon,
		Pid:    param.Pid,
		Url:    param.Url,
		Status: param.Status,
	}
	var result *gorm.DB
	if id == 0 {
		result = model.UseCon().Create(&arr)
	} else {
		result = model.UseCon().Where("id=?", id).Updates(&arr)

	}
	if result.Error != nil {
		return result.Error
	}
	if len(param.PageFunc) > 0 {
		//绑定页面得关系
		var pageFuncs = make([]model.SystemPageFunc, len(param.PageFunc))
		for k, _ := range pageFuncs {
			pageFuncs[k].PageId = arr.Id
			pageFuncs[k].FuncId = uint(param.PageFunc[k])
		}
		result = model.UseCon().Create(&pageFuncs)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (services *SystemPage) Show(id int) (interface{}, error) {
	//查询数据
	var arr model.SystemPageFuncList
	err := model.UseCon().Where("id =?", id).Preload("SystemPageFuncTo.SystemFunc").Preload("SystemPageFuncTo").First(&arr)
	if err.Error != nil {
		return nil, err.Error
	}
	return arr, nil
}
