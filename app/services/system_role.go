package services

import (
	"Walker/app/http/model"
	"Walker/app/http/validator/admin"
	"Walker/app/utils/data"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SystemRole struct {
}
type Tree struct {
	Id             uint
	Pid            uint
	Name           string
	FuncList       []*TreeSysPageFunc
	SystemRolePage model.SystemRolePage
	Child          []*Tree
}
type TreeList map[uint]*Tree

func (service *SystemRole) Index(ctx *gin.Context) (map[string]interface{}, error) {
	db := model.UseCon()
	if str := ctx.Query("keywords"); str != "" {
		db = db.Where("name like ?", "%"+str+"%")
	}
	if str := ctx.Query("status"); str != "" {
		db = db.Where("status like ?", str)
	}
	var count int64
	db.Count(&count)
	var systemRoleList []model.SystemRole
	result := db.Scopes(model.Paginate(ctx)).Find(&systemRoleList)
	if result.Error != nil {
		return nil, result.Error
	}
	return map[string]interface{}{
		"data":       systemRoleList,
		"total_rows": count,
	}, nil
}
func (service *SystemRole) Edit(param admin.SystemRoleRequest, id uint) error {
	arr := model.SystemRole{
		Name:   param.Name,
		Pid:    param.Pid,
		Status: param.Status,
	}
	var result *gorm.DB
	if id > 0 {
		result = model.UseCon().Where("id=?", id).Updates(&arr)
	} else {
		result = model.UseCon().Create(&arr)
	}
	if result.Error != nil {
		return result.Error
	}
	if len(param.RoleData.PageIds) > 0 {
		rolePageData := make([]model.SystemRolePage, len(param.RoleData.PageIds))
		for k, _ := range rolePageData {
			rolePageData[k].RoleId = arr.Id
			rolePageData[k].PageId = uint(param.RoleData.PageIds[k])
		}
		model.UseCon().Create(&rolePageData)
	}
	var casbinRuleData []model.CasbinRule
	if len(param.RoleData.FuncIds) > 0 {
		var funcData []model.SystemFunc
		model.UseCon().Where("id In ?", param.RoleData.FuncIds).Find(&funcData)

		for k, _ := range funcData {
			casbinRuleData = append(casbinRuleData, model.CasbinRule{
				Ptype:  "p",
				V0:     data.UintToString(arr.Id),
				V1:     funcData[k].Url,
				V2:     funcData[k].Method,
				FuncId: funcData[k].Id,
			})
		}
	}
	if param.Pid > 0 {
		result = model.UseCon().Where("ptype='g' AND V0=? AND V1=?", id, param.Pid).First(&model.CasbinRule{})
		if result.RowsAffected > 0 {
			model.UseCon().Where("ptype='g' AND V0=? AND V1=?", id, param.Pid).Updates(&model.CasbinRule{
				V1: data.UintToString(param.Pid),
			})
		} else {
			casbinRuleData = append(casbinRuleData, model.CasbinRule{
				Ptype: "g",
				V0:    data.UintToString(arr.Id),
				V1:    data.UintToString(param.Pid),
			})
		}
	}
	model.UseCon().Create(&casbinRuleData)
	return nil
}

func (service *SystemRole) Destroy(ids []string) error {
	//角色账号删除
	model.UseCon().Delete(&model.SystemRole{}, ids)
	//删除该权限所能看到的页面
	model.UseCon().Debug().Where("role_id in ?", ids).Delete(&model.SystemRolePage{})
	//删除该角色所拥有的访问权限和继承跟被继承数据
	model.UseCon().Where("((ptype='g') and (V0 In ? OR V1 IN ?)) OR (ptype='p') and V0 In ?", ids, ids, ids).Delete(&model.CasbinRule{})

	return nil
}

func (service *SystemRole) Show(id uint) (map[string]interface{}, error) {
	var systemRole model.SystemRoleTo
	//查询所有的角色
	result := model.UseCon().Where("id=?", id).Preload("SystemRole").First(&systemRole)
	if result.RowsAffected == 0 || result.Error != nil {
		return nil, errors.New("数据不存在")
	}
	//查询所有的页面标记那些被选中
	var systemPage []model.SystemPageTo
	result = model.UseCon().Preload("SystemRolePage", "role_id=?", id).Find(&systemPage)
	if result.Error != nil {
		return nil, errors.New("角色页面数据未查询成功")
	}
	var systemPageFunc []model.SystemPageFuncTo
	model.UseCon().Preload("SystemFunc").Preload("SystemFunc.CasbinRule", "ptype='p' AND V0=?", id).Find(&systemPageFunc)
	//处理成树形结构
	data, _ := service.BuildTreePageFunc(systemPage, systemPageFunc)
	return map[string]interface{}{
		"role_info":   systemRole,
		"casbin_list": data,
	}, nil
}

func (service *SystemRole) BuildTreePageFunc(systemPage []model.SystemPageTo, systemPageFunc []model.SystemPageFuncTo) (interface{}, error) {

	list := make(map[uint][]*TreeSysPageFunc)
	for _, pageFunc := range systemPageFunc {
		var treeSysPageFunc TreeSysPageFunc
		treeSysPageFunc.Id = pageFunc.Id
		treeSysPageFunc.FuncId = pageFunc.FuncId
		treeSysPageFunc.PageId = pageFunc.PageId
		treeSysPageFunc.SystemFunc = pageFunc.SystemFunc
		list[pageFunc.PageId] = append(list[pageFunc.PageId], &treeSysPageFunc)

	}
	var newTree []*Tree
	data := make(TreeList)
	for _, v := range systemPage {
		var treeItem Tree
		treeItem.Id = v.Id
		treeItem.Pid = v.Pid
		treeItem.Name = v.Name
		treeItem.SystemRolePage = v.SystemRolePage
		if _, ok := list[v.Id]; ok == true {
			treeItem.FuncList = list[v.Id]
		}
		data[v.Id] = &treeItem
		if v.Pid == 0 {
			newTree = append(newTree, &treeItem)
		} else {
			data[v.Pid].Child = append(data[v.Pid].Child, &treeItem)
		}

	}
	return newTree, nil
}

type TreeSysPageFunc struct {
	Id         uint
	PageId     uint
	FuncId     uint
	SystemFunc model.SystemFuncTo
}
