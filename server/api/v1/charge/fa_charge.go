package charge

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/charge"
	chargeReq "github.com/flipped-aurora/gin-vue-admin/server/model/charge/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FaChargeApi struct {
}

var faChargeService = service.ServiceGroupApp.ChargeServiceGroup.FaChargeService

// CreateFaCharge 创建FaCharge
// @Tags FaCharge
// @Summary 创建FaCharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body charge.FaCharge true "创建FaCharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faCharge/createFaCharge [post]
func (faChargeApi *FaChargeApi) CreateFaCharge(c *gin.Context) {
	var faCharge charge.FaCharge
	err := c.ShouldBindJSON(&faCharge)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	if uid > 1 {
		faCharge.ShangId = uid
	}
	if err := faChargeService.CreateFaCharge(faCharge); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFaCharge 删除FaCharge
// @Tags FaCharge
// @Summary 删除FaCharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body charge.FaCharge true "删除FaCharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faCharge/deleteFaCharge [delete]
func (faChargeApi *FaChargeApi) DeleteFaCharge(c *gin.Context) {
	var faCharge charge.FaCharge
	err := c.ShouldBindJSON(&faCharge)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faChargeService.DeleteFaCharge(faCharge); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFaChargeByIds 批量删除FaCharge
// @Tags FaCharge
// @Summary 批量删除FaCharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FaCharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /faCharge/deleteFaChargeByIds [delete]
func (faChargeApi *FaChargeApi) DeleteFaChargeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faChargeService.DeleteFaChargeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFaCharge 更新FaCharge
// @Tags FaCharge
// @Summary 更新FaCharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body charge.FaCharge true "更新FaCharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /faCharge/updateFaCharge [put]
func (faChargeApi *FaChargeApi) UpdateFaCharge(c *gin.Context) {
	var faCharge charge.FaCharge
	err := c.ShouldBindJSON(&faCharge)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	if uid > 1 {
		faCharge.ShangId = uid
	}
	if err := faChargeService.UpdateFaCharge(faCharge); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFaCharge 用id查询FaCharge
// @Tags FaCharge
// @Summary 用id查询FaCharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query charge.FaCharge true "用id查询FaCharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /faCharge/findFaCharge [get]
func (faChargeApi *FaChargeApi) FindFaCharge(c *gin.Context) {
	var faCharge charge.FaCharge
	err := c.ShouldBindQuery(&faCharge)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refaCharge, err := faChargeService.GetFaCharge(faCharge.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refaCharge": refaCharge}, c)
	}
}

// GetFaChargeList 分页获取FaCharge列表
// @Tags FaCharge
// @Summary 分页获取FaCharge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query chargeReq.FaChargeSearch true "分页获取FaCharge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faCharge/getFaChargeList [get]
func (faChargeApi *FaChargeApi) GetFaChargeList(c *gin.Context) {
	var pageInfo chargeReq.FaChargeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	if uid > 1 {
		pageInfo.ShangId = uid
	}
	if list, total, err := faChargeService.GetFaChargeInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
