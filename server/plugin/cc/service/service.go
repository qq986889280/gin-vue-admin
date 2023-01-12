package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cc/model"
)

type CcService struct{}

func (e *CcService) Do(req model.Request) (res model.Response, err error) {
	// 写你的CC业务逻辑
	return res, nil
}
