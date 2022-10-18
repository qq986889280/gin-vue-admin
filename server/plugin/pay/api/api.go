package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/charge"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pay/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PayApi struct{}

// @Tags Pay
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /pay/routerName [post]
func (p *PayApi) Order(c *gin.Context) {

	// var plug model.Request
	var faCharge charge.FaCharge
	_ = c.ShouldBindJSON(&faCharge)

	if res, err := service.ServiceGroupApp.Order(faCharge); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {

		response.OkWithDetailed(res, "成功", c)
	}
}
