package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/charge"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pay/model"
)

type PayService struct{}

func (e *PayService) Order(req charge.FaCharge) (res model.Response, err error) {
	// 写你的业务逻辑
	err = global.GVA_DB.Create(&req).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
