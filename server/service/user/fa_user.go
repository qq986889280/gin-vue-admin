package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
)

type FaUserService struct {
}

// CreateFaUser 创建FaUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService) CreateFaUser(faUser user.FaUser) (err error) {
	err = global.GVA_DB.Create(&faUser).Error
	return err
}

// DeleteFaUser 删除FaUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService)DeleteFaUser(faUser user.FaUser) (err error) {
	err = global.GVA_DB.Delete(&faUser).Error
	return err
}

// DeleteFaUserByIds 批量删除FaUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService)DeleteFaUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]user.FaUser{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFaUser 更新FaUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService)UpdateFaUser(faUser user.FaUser) (err error) {
	err = global.GVA_DB.Save(&faUser).Error
	return err
}

// GetFaUser 根据id获取FaUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService)GetFaUser(id uint) (faUser user.FaUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&faUser).Error
	return
}

// GetFaUserInfoList 分页获取FaUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (faUserService *FaUserService)GetFaUserInfoList(info userReq.FaUserSearch) (list []user.FaUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&user.FaUser{})
    var faUsers []user.FaUser
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&faUsers).Error
	return  faUsers, total, err
}
