package charge

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/charge"
	chargeReq "github.com/flipped-aurora/gin-vue-admin/server/model/charge/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FaShangService struct {
}

// CreateFaShang 创建FaShang记录
// Author [piexlmax](https://github.com/piexlmax)
func (faShangService *FaShangService) CreateFaShang(faShang charge.FaShang) (err error) {
	err = global.GVA_DB.Create(&faShang).Error
	return err
}

// DeleteFaShang 删除FaShang记录
// Author [piexlmax](https://github.com/piexlmax)
func (faShangService *FaShangService) DeleteFaShang(faShang charge.FaShang) (err error) {
	err = global.GVA_DB.Delete(&faShang).Error
	return err
}

// DeleteFaShangByIds 批量删除FaShang记录
// Author [piexlmax](https://github.com/piexlmax)
func (faShangService *FaShangService) DeleteFaShangByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]charge.FaShang{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFaShang 更新FaShang记录
// Author [piexlmax](https://github.com/piexlmax)
func (faShangService *FaShangService) UpdateFaShang(faShang charge.FaShang) (err error) {
	err = global.GVA_DB.Save(&faShang).Error
	return err
}

// GetFaShang 根据id获取FaShang记录
// Author [piexlmax](https://github.com/piexlmax)
func (faShangService *FaShangService) GetFaShang(id uint) (faShang charge.FaShang, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&faShang).Error
	return
}

// GetFaShangInfoList 分页获取FaShang记录
// Author [piexlmax](https://github.com/piexlmax)
func (faShangService *FaShangService) GetFaShangInfoList(info chargeReq.FaShangSearch) (list []charge.FaShang, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&charge.FaShang{})
	var faShangs []charge.FaShang
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Net != nil {
		db = db.Where("net = ?", info.Net)
	}
	if info.ShangId > 1 {
		db = db.Where("shang_id = ?", info.ShangId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&faShangs).Error
	return faShangs, total, err
}
