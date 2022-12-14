package charge

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/charge"
	chargeReq "github.com/flipped-aurora/gin-vue-admin/server/model/charge/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FaShangApi struct {
}

var faShangService = service.ServiceGroupApp.ChargeServiceGroup.FaShangService

// CreateFaShang 创建FaShang
// @Tags FaShang
// @Summary 创建FaShang
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body charge.FaShang true "创建FaShang"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faShang/createFaShang [post]
func (faShangApi *FaShangApi) CreateFaShang(c *gin.Context) {
	var faShang charge.FaShang
	err := c.ShouldBindJSON(&faShang)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	if uid > 1 {
		faShang.ShangId = uid
	}
	faShang.Fun = utils.Hash(strconv.Itoa(faShang.ShangId)+faShang.Ptype+faShang.Contractaddr+faShang.To+strconv.Itoa(faShang.Net)+"fvbexop", "sha1", false)
	verify := utils.Rules{
		"Net": {utils.NotEmpty()},
	}
	if err := utils.Verify(faShang, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faShangService.CreateFaShang(faShang); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFaShang 删除FaShang
// @Tags FaShang
// @Summary 删除FaShang
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body charge.FaShang true "删除FaShang"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faShang/deleteFaShang [delete]
func (faShangApi *FaShangApi) DeleteFaShang(c *gin.Context) {
	var faShang charge.FaShang
	err := c.ShouldBindJSON(&faShang)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faShangService.DeleteFaShang(faShang); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFaShangByIds 批量删除FaShang
// @Tags FaShang
// @Summary 批量删除FaShang
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FaShang"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /faShang/deleteFaShangByIds [delete]
func (faShangApi *FaShangApi) DeleteFaShangByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faShangService.DeleteFaShangByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFaShang 更新FaShang
// @Tags FaShang
// @Summary 更新FaShang
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body charge.FaShang true "更新FaShang"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /faShang/updateFaShang [put]
func (faShangApi *FaShangApi) UpdateFaShang(c *gin.Context) {
	var faShang charge.FaShang
	err := c.ShouldBindJSON(&faShang)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	if uid > 1 {
		faShang.ShangId = uid
	}
	faShang.Fun = utils.Hash(strconv.Itoa(faShang.ShangId)+faShang.Ptype+faShang.Contractaddr+faShang.To+strconv.Itoa(faShang.Net)+"fvbexop", "sha1", false)

	verify := utils.Rules{
		"Net": {utils.NotEmpty()},
	}
	if err := utils.Verify(faShang, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faShangService.UpdateFaShang(faShang); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFaShang 用id查询FaShang
// @Tags FaShang
// @Summary 用id查询FaShang
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query charge.FaShang true "用id查询FaShang"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /faShang/findFaShang [get]
func (faShangApi *FaShangApi) FindFaShang(c *gin.Context) {
	var faShang charge.FaShang
	err := c.ShouldBindQuery(&faShang)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refaShang, err := faShangService.GetFaShang(faShang.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refaShang": refaShang}, c)
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
func (faShangApi *FaShangApi) GetFaShangList(c *gin.Context) {
	var pageInfo chargeReq.FaShangSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid, _ := strconv.Atoi(c.Request.Header.Get("x-user-id"))
	if uid > 1 {
		pageInfo.ShangId = uid
	}
	if list, total, err := faShangService.GetFaShangInfoList(pageInfo); err != nil {
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
