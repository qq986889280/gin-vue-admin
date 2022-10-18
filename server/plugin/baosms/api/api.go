package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/baosms/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/baosms/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaoSmsApi struct{}

// @Tags BaoSms
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /baosms/routerName [post]
func (p *BaoSmsApi) Send(c *gin.Context) {

	var plug model.Request
	_ = c.ShouldBindJSON(&plug)

	if err := service.ServiceGroupApp.SendService(plug); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
	} else {

		response.OkWithData("成功", c)
	}
}
