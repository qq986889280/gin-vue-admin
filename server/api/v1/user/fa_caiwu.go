package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FaCaiwuApi struct {
}

var faCaiwuService = service.ServiceGroupApp.UserServiceGroup.FaCaiwuService

// CreateFaCaiwu 创建FaCaiwu
// @Tags FaCaiwu
// @Summary 创建FaCaiwu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.FaCaiwu true "创建FaCaiwu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faCaiwu/createFaCaiwu [post]
func (faCaiwuApi *FaCaiwuApi) CreateFaCaiwu(c *gin.Context) {
	var faCaiwu user.FaCaiwu
	err := c.ShouldBindJSON(&faCaiwu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faCaiwuService.CreateFaCaiwu(faCaiwu); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFaCaiwu 删除FaCaiwu
// @Tags FaCaiwu
// @Summary 删除FaCaiwu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.FaCaiwu true "删除FaCaiwu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faCaiwu/deleteFaCaiwu [delete]
func (faCaiwuApi *FaCaiwuApi) DeleteFaCaiwu(c *gin.Context) {
	var faCaiwu user.FaCaiwu
	err := c.ShouldBindJSON(&faCaiwu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faCaiwuService.DeleteFaCaiwu(faCaiwu); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFaCaiwuByIds 批量删除FaCaiwu
// @Tags FaCaiwu
// @Summary 批量删除FaCaiwu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FaCaiwu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /faCaiwu/deleteFaCaiwuByIds [delete]
func (faCaiwuApi *FaCaiwuApi) DeleteFaCaiwuByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faCaiwuService.DeleteFaCaiwuByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFaCaiwu 更新FaCaiwu
// @Tags FaCaiwu
// @Summary 更新FaCaiwu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.FaCaiwu true "更新FaCaiwu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /faCaiwu/updateFaCaiwu [put]
func (faCaiwuApi *FaCaiwuApi) UpdateFaCaiwu(c *gin.Context) {
	var faCaiwu user.FaCaiwu
	err := c.ShouldBindJSON(&faCaiwu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faCaiwuService.UpdateFaCaiwu(faCaiwu); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFaCaiwu 用id查询FaCaiwu
// @Tags FaCaiwu
// @Summary 用id查询FaCaiwu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.FaCaiwu true "用id查询FaCaiwu"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /faCaiwu/findFaCaiwu [get]
func (faCaiwuApi *FaCaiwuApi) FindFaCaiwu(c *gin.Context) {
	var faCaiwu user.FaCaiwu
	err := c.ShouldBindQuery(&faCaiwu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refaCaiwu, err := faCaiwuService.GetFaCaiwu(faCaiwu.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refaCaiwu": refaCaiwu}, c)
	}
}

// GetFaCaiwuList 分页获取FaCaiwu列表
// @Tags FaCaiwu
// @Summary 分页获取FaCaiwu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.FaCaiwuSearch true "分页获取FaCaiwu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faCaiwu/getFaCaiwuList [get]
func (faCaiwuApi *FaCaiwuApi) GetFaCaiwuList(c *gin.Context) {
	var pageInfo userReq.FaCaiwuSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, price, err := faCaiwuService.GetFaCaiwuInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResultCaiwu{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
			Price:    price,
		}, "获取成功", c)
	}
}
