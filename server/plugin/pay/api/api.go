package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/charge"
	chargeReq "github.com/flipped-aurora/gin-vue-admin/server/model/charge/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pay/model"
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
func (p *PayApi) PayCheck(c *gin.Context) {

	var info model.RequestPayCheck
	_ = c.ShouldBindJSON(&info)
	var faCharge charge.FaCharge
	db := global.GVA_DB.Model(&charge.FaCharge{}).Where("shang_id = ?", info.ShangId).Where("order_sn = ?", info.OrderSn).Find(&faCharge)
	var total int64
	db.Count(&total)
	if total > 0 {
		response.FailWithMessage("订单已存在，请勿重复下单", c)
		return
	}
	response.Ok(c)
}

// @Tags Pay
// @Summary 请手动填写接口功能
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /pay/routerName [post]
func (p *PayApi) Order(c *gin.Context) {

	// var plug model.Request
	var faCharge charge.FaCharge
	_ = c.ShouldBindJSON(&faCharge)

	if err := service.ServiceGroupApp.Order(faCharge); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.Ok(c)
	}
}

// GetFaShangList 分页获取FaShang列表
// @Tags FaShang
// @Summary 分页获取FaShang列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query chargeReq.FaShangSearch true "分页获取FaShang列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faShang/getFaShangList [get]
func (p *PayApi) GetFaShangList(c *gin.Context) {
	var pageInfo chargeReq.FaShangSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := service.ServiceGroupApp.GetFaShangInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.Fail(c)
	} else {
		response.OkWithData(list, c)
	}
}

// @Tags Pay
// @Summary 充值订单检查
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /pay/ChargeCheck [post]
func (p *PayApi) ChargeCheck(c *gin.Context) {
	service.ServiceGroupApp.PayService.ChargeNotice()
	response.Ok(c)
}

// @Tags Pay
// @Summary 充值订单检查
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /pay/ChargeCheck [post]
func (p *PayApi) ChargeCheck2(c *gin.Context) {
	list, err := service.ServiceGroupApp.PayService.ChargeCheck()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(list, c)
}
