package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
)

type FaUserLevelService struct {
}

// CreateFaUserLevel 创建FaUserLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserLevelService *FaUserLevelService) CreateFaUserLevel(faUserLevel user.FaUserLevel) (err error) {
	err = global.GVA_DB.Create(&faUserLevel).Error
	return err
}

// DeleteFaUserLevel 删除FaUserLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserLevelService *FaUserLevelService)DeleteFaUserLevel(faUserLevel user.FaUserLevel) (err error) {
	err = global.GVA_DB.Delete(&faUserLevel).Error
	return err
}

// DeleteFaUserLevelByIds 批量删除FaUserLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserLevelService *FaUserLevelService)DeleteFaUserLevelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]user.FaUserLevel{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFaUserLevel 更新FaUserLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserLevelService *FaUserLevelService)UpdateFaUserLevel(faUserLevel user.FaUserLevel) (err error) {
	err = global.GVA_DB.Save(&faUserLevel).Error
	return err
}

// GetFaUserLevel 根据id获取FaUserLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserLevelService *FaUserLevelService)GetFaUserLevel(id uint) (faUserLevel user.FaUserLevel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&faUserLevel).Error
	return
}

// GetFaUserLevelInfoList 分页获取FaUserLevel记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserLevelService *FaUserLevelService)GetFaUserLevelInfoList(info userReq.FaUserLevelSearch) (list []user.FaUserLevel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&user.FaUserLevel{})
    var faUserLevels []user.FaUserLevel
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&faUserLevels).Error
	return  faUserLevels, total, err
}
