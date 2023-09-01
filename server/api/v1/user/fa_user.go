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

type FaUserApi struct {
}

var faUserService = service.ServiceGroupApp.UserServiceGroup.FaUserService

// CreateFaUser 创建FaUser
// @Tags FaUser
// @Summary 创建FaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.FaUser true "创建FaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faUser/createFaUser [post]
func (faUserApi *FaUserApi) CreateFaUser(c *gin.Context) {
	var faUser user.FaUser
	err := c.ShouldBindJSON(&faUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faUserService.CreateFaUser(faUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFaUser 删除FaUser
// @Tags FaUser
// @Summary 删除FaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.FaUser true "删除FaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faUser/deleteFaUser [delete]
func (faUserApi *FaUserApi) DeleteFaUser(c *gin.Context) {
	var faUser user.FaUser
	err := c.ShouldBindJSON(&faUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faUserService.DeleteFaUser(faUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFaUserByIds 批量删除FaUser
// @Tags FaUser
// @Summary 批量删除FaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /faUser/deleteFaUserByIds [delete]
func (faUserApi *FaUserApi) DeleteFaUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faUserService.DeleteFaUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFaUser 更新FaUser
// @Tags FaUser
// @Summary 更新FaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.FaUser true "更新FaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /faUser/updateFaUser [put]
func (faUserApi *FaUserApi) UpdateFaUser(c *gin.Context) {
	var faUser user.FaUser
	err := c.ShouldBindJSON(&faUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faUserService.UpdateFaUser(faUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFaUser 用id查询FaUser
// @Tags FaUser
// @Summary 用id查询FaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.FaUser true "用id查询FaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /faUser/findFaUser [get]
func (faUserApi *FaUserApi) FindFaUser(c *gin.Context) {
	var faUser user.FaUser
	err := c.ShouldBindQuery(&faUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refaUser, err := faUserService.GetFaUser(faUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refaUser": refaUser}, c)
	}
}

// GetFaUserList 分页获取FaUser列表
// @Tags FaUser
// @Summary 分页获取FaUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.FaUserSearch true "分页获取FaUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faUser/getFaUserList [get]
func (faUserApi *FaUserApi) GetFaUserList(c *gin.Context) {
	var pageInfo userReq.FaUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := faUserService.GetFaUserInfoList(pageInfo); err != nil {
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

// UserCharge 后台充值
// @Tags FaUser
// @Summary 后台充值
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.UserCharge true "后台充值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /faUser/UserCharge [post]
func (faUserApi *FaUserApi) UserCharge(c *gin.Context) {
	var info user.UserCharge
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faUserService.Charge(info); err != nil {
		global.GVA_LOG.Error("后台充值失败!", zap.Error(err))
		response.Fail(c)
	} else {
		response.Ok(c)
	}
}

// UserTeam 后台查看团队
// @Tags FaUser
// @Summary 后台查看团队
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.UserCharge true "后台查看团队"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /faUser/UserTeam [get]
func (faUserApi *FaUserApi) UserTeam(c *gin.Context) {
	// claims, _ := utils.GetClaims(c)
	// claims, _ := c.Get("claims")
	// dd, _ := json.MarshalIndent(claims, "", "  ")
	// fmt.Printf("%s\n", dd)
	var get request.GetById
	err := c.ShouldBindQuery(&get)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if RE, err := faUserService.Team(get.ID); err != nil {
		global.GVA_LOG.Error("", zap.Error(err))
		response.Fail(c)
	} else {
		response.OkWithData(RE, c)
	}
}
