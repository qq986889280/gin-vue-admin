package user

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
)

type FaCaiwuService struct {
}

// CreateFaCaiwu 创建FaCaiwu记录
// Author [piexlmax](https://github.com/piexlmax)
func (faCaiwuService *FaCaiwuService) CreateFaCaiwu(faCaiwu user.FaCaiwu) (err error) {
	err = global.GVA_DB.Create(&faCaiwu).Error
	return err
}

// DeleteFaCaiwu 删除FaCaiwu记录
// Author [piexlmax](https://github.com/piexlmax)
func (faCaiwuService *FaCaiwuService) DeleteFaCaiwu(faCaiwu user.FaCaiwu) (err error) {
	err = global.GVA_DB.Delete(&faCaiwu).Error
	return err
}

// DeleteFaCaiwuByIds 批量删除FaCaiwu记录
// Author [piexlmax](https://github.com/piexlmax)
func (faCaiwuService *FaCaiwuService) DeleteFaCaiwuByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]user.FaCaiwu{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFaCaiwu 更新FaCaiwu记录
// Author [piexlmax](https://github.com/piexlmax)
func (faCaiwuService *FaCaiwuService) UpdateFaCaiwu(faCaiwu user.FaCaiwu) (err error) {
	err = global.GVA_DB.Save(&faCaiwu).Error
	return err
}

// GetFaCaiwu 根据id获取FaCaiwu记录
// Author [piexlmax](https://github.com/piexlmax)
func (faCaiwuService *FaCaiwuService) GetFaCaiwu(id uint) (faCaiwu user.FaCaiwu, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&faCaiwu).Error
	return
}

// GetFaCaiwuInfoList 分页获取FaCaiwu记录
// Author [piexlmax](https://github.com/piexlmax)
func (faCaiwuService *FaCaiwuService) GetFaCaiwuInfoList(info userReq.FaCaiwuSearch) (list []user.FaCaiwu, total int64, price float64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user.FaCaiwu{})
	var faCaiwus []user.FaCaiwu
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}
	if info.Ptype != "" {
		db = db.Where("ptype = ?", info.Ptype)
	}
	if info.OrderId != "" {
		db = db.Where("order_id = ?", info.OrderId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["type"] = true
	orderMap["ptype"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&faCaiwus).Error
	if err != nil {
		return
	}
	err = db.Select("SUM(price) as price").Find(&price).Error

	fmt.Println(price)
	return faCaiwus, total, price, err
}
