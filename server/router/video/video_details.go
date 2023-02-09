package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VideoDetailsRouter struct {
}

// InitVideoDetailsRouter 初始化 VideoDetails 路由信息
func (s *VideoDetailsRouter) InitVideoDetailsRouter(Router *gin.RouterGroup) {
	videoDetailsRouter := Router.Group("videoDetails").Use(middleware.OperationRecord())
	videoDetailsRouterWithoutRecord := Router.Group("videoDetails")
	var videoDetailsApi = v1.ApiGroupApp.VideoApiGroup.VideoDetailsApi
	{
		videoDetailsRouter.POST("createVideoDetails", videoDetailsApi.CreateVideoDetails)   // 新建VideoDetails
		videoDetailsRouter.DELETE("deleteVideoDetails", videoDetailsApi.DeleteVideoDetails) // 删除VideoDetails
		videoDetailsRouter.DELETE("deleteVideoDetailsByIds", videoDetailsApi.DeleteVideoDetailsByIds) // 批量删除VideoDetails
		videoDetailsRouter.PUT("updateVideoDetails", videoDetailsApi.UpdateVideoDetails)    // 更新VideoDetails
	}
	{
		videoDetailsRouterWithoutRecord.GET("findVideoDetails", videoDetailsApi.FindVideoDetails)        // 根据ID获取VideoDetails
		videoDetailsRouterWithoutRecord.GET("getVideoDetailsList", videoDetailsApi.GetVideoDetailsList)  // 获取VideoDetails列表
	}
}
