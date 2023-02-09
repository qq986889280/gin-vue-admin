package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
	videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/video/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VideoDetailsApi struct {
}

var videoDetailsService = service.ServiceGroupApp.VideoServiceGroup.VideoDetailsService

// CreateVideoDetails 创建VideoDetails
// @Tags VideoDetails
// @Summary 创建VideoDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body video.VideoDetails true "创建VideoDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoDetails/createVideoDetails [post]
func (videoDetailsApi *VideoDetailsApi) CreateVideoDetails(c *gin.Context) {
	var videoDetails video.VideoDetails
	err := c.ShouldBindJSON(&videoDetails)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoDetailsService.CreateVideoDetails(videoDetails); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteVideoDetails 删除VideoDetails
// @Tags VideoDetails
// @Summary 删除VideoDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body video.VideoDetails true "删除VideoDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /videoDetails/deleteVideoDetails [delete]
func (videoDetailsApi *VideoDetailsApi) DeleteVideoDetails(c *gin.Context) {
	var videoDetails video.VideoDetails
	err := c.ShouldBindJSON(&videoDetails)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoDetailsService.DeleteVideoDetails(videoDetails); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteVideoDetailsByIds 批量删除VideoDetails
// @Tags VideoDetails
// @Summary 批量删除VideoDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VideoDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /videoDetails/deleteVideoDetailsByIds [delete]
func (videoDetailsApi *VideoDetailsApi) DeleteVideoDetailsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoDetailsService.DeleteVideoDetailsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateVideoDetails 更新VideoDetails
// @Tags VideoDetails
// @Summary 更新VideoDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body video.VideoDetails true "更新VideoDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /videoDetails/updateVideoDetails [put]
func (videoDetailsApi *VideoDetailsApi) UpdateVideoDetails(c *gin.Context) {
	var videoDetails video.VideoDetails
	err := c.ShouldBindJSON(&videoDetails)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := videoDetailsService.UpdateVideoDetails(videoDetails); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindVideoDetails 用id查询VideoDetails
// @Tags VideoDetails
// @Summary 用id查询VideoDetails
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query video.VideoDetails true "用id查询VideoDetails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /videoDetails/findVideoDetails [get]
func (videoDetailsApi *VideoDetailsApi) FindVideoDetails(c *gin.Context) {
	var videoDetails request.GetById
	err := c.ShouldBindQuery(&videoDetails)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if revideoDetails, err := videoDetailsService.GetVideoDetails(videoDetails.Uint()); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"revideoDetails": revideoDetails}, c)
	}
}

// GetVideoDetailsList 分页获取VideoDetails列表
// @Tags VideoDetails
// @Summary 分页获取VideoDetails列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query videoReq.VideoDetailsSearch true "分页获取VideoDetails列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /videoDetails/getVideoDetailsList [get]
func (videoDetailsApi *VideoDetailsApi) GetVideoDetailsList(c *gin.Context) {
	var pageInfo videoReq.VideoDetailsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := videoDetailsService.GetVideoDetailsInfoList(pageInfo); err != nil {
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
