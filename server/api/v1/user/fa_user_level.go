package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/user"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type FaUserLevelApi struct {
}

var faUserLevelService = service.ServiceGroupApp.UserServiceGroup.FaUserLevelService


// CreateFaUserLevel 创建FaUserLevel
// @Tags FaUserLevel
// @Summary 创建FaUserLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.FaUserLevel true "创建FaUserLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faUserLevel/createFaUserLevel [post]
func (faUserLevelApi *FaUserLevelApi) CreateFaUserLevel(c *gin.Context) {
	var faUserLevel user.FaUserLevel
	err := c.ShouldBindJSON(&faUserLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faUserLevelService.CreateFaUserLevel(faUserLevel); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFaUserLevel 删除FaUserLevel
// @Tags FaUserLevel
// @Summary 删除FaUserLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.FaUserLevel true "删除FaUserLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /faUserLevel/deleteFaUserLevel [delete]
func (faUserLevelApi *FaUserLevelApi) DeleteFaUserLevel(c *gin.Context) {
	var faUserLevel user.FaUserLevel
	err := c.ShouldBindJSON(&faUserLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faUserLevelService.DeleteFaUserLevel(faUserLevel); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFaUserLevelByIds 批量删除FaUserLevel
// @Tags FaUserLevel
// @Summary 批量删除FaUserLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FaUserLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /faUserLevel/deleteFaUserLevelByIds [delete]
func (faUserLevelApi *FaUserLevelApi) DeleteFaUserLevelByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faUserLevelService.DeleteFaUserLevelByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFaUserLevel 更新FaUserLevel
// @Tags FaUserLevel
// @Summary 更新FaUserLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.FaUserLevel true "更新FaUserLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /faUserLevel/updateFaUserLevel [put]
func (faUserLevelApi *FaUserLevelApi) UpdateFaUserLevel(c *gin.Context) {
	var faUserLevel user.FaUserLevel
	err := c.ShouldBindJSON(&faUserLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := faUserLevelService.UpdateFaUserLevel(faUserLevel); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFaUserLevel 用id查询FaUserLevel
// @Tags FaUserLevel
// @Summary 用id查询FaUserLevel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.FaUserLevel true "用id查询FaUserLevel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /faUserLevel/findFaUserLevel [get]
func (faUserLevelApi *FaUserLevelApi) FindFaUserLevel(c *gin.Context) {
	var faUserLevel user.FaUserLevel
	err := c.ShouldBindQuery(&faUserLevel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refaUserLevel, err := faUserLevelService.GetFaUserLevel(faUserLevel.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refaUserLevel": refaUserLevel}, c)
	}
}

// GetFaUserLevelList 分页获取FaUserLevel列表
// @Tags FaUserLevel
// @Summary 分页获取FaUserLevel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.FaUserLevelSearch true "分页获取FaUserLevel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /faUserLevel/getFaUserLevelList [get]
func (faUserLevelApi *FaUserLevelApi) GetFaUserLevelList(c *gin.Context) {
	var pageInfo userReq.FaUserLevelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := faUserLevelService.GetFaUserLevelInfoList(pageInfo); err != nil {
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
