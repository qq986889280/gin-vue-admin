package charge

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/charge"
	chargeReq "github.com/flipped-aurora/gin-vue-admin/server/model/charge/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FaChargeService struct {
}

// CreateFaCharge 创建FaCharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (faChargeService *FaChargeService) CreateFaCharge(faCharge charge.FaCharge) (err error) {
	err = global.GVA_DB.Create(&faCharge).Error
	return err
}

// DeleteFaCharge 删除FaCharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (faChargeService *FaChargeService) DeleteFaCharge(faCharge charge.FaCharge) (err error) {
	err = global.GVA_DB.Delete(&faCharge).Error
	return err
}

// DeleteFaChargeByIds 批量删除FaCharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (faChargeService *FaChargeService) DeleteFaChargeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]charge.FaCharge{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFaCharge 更新FaCharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (faChargeService *FaChargeService) UpdateFaCharge(faCharge charge.FaCharge) (err error) {
	err = global.GVA_DB.Save(&faCharge).Error
	return err
}

// GetFaCharge 根据id获取FaCharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (faChargeService *FaChargeService) GetFaCharge(id uint) (faCharge charge.FaCharge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&faCharge).Error
	return
}

// GetFaChargeInfoList 分页获取FaCharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (faChargeService *FaChargeService) GetFaChargeInfoList(info chargeReq.FaChargeSearch) (list []charge.FaCharge, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&charge.FaCharge{})
	var faCharges []charge.FaCharge
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&faCharges).Error
	return faCharges, total, err
}
