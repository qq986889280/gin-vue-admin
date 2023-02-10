package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FaCaiwuRouter struct {
}

// InitFaCaiwuRouter 初始化 FaCaiwu 路由信息
func (s *FaCaiwuRouter) InitFaCaiwuRouter(Router *gin.RouterGroup) {
	faCaiwuRouter := Router.Group("faCaiwu").Use(middleware.OperationRecord())
	faCaiwuRouterWithoutRecord := Router.Group("faCaiwu")
	var faCaiwuApi = v1.ApiGroupApp.UserApiGroup.FaCaiwuApi
	{
		faCaiwuRouter.POST("createFaCaiwu", faCaiwuApi.CreateFaCaiwu)   // 新建FaCaiwu
		faCaiwuRouter.DELETE("deleteFaCaiwu", faCaiwuApi.DeleteFaCaiwu) // 删除FaCaiwu
		faCaiwuRouter.DELETE("deleteFaCaiwuByIds", faCaiwuApi.DeleteFaCaiwuByIds) // 批量删除FaCaiwu
		faCaiwuRouter.PUT("updateFaCaiwu", faCaiwuApi.UpdateFaCaiwu)    // 更新FaCaiwu
	}
	{
		faCaiwuRouterWithoutRecord.GET("findFaCaiwu", faCaiwuApi.FindFaCaiwu)        // 根据ID获取FaCaiwu
		faCaiwuRouterWithoutRecord.GET("getFaCaiwuList", faCaiwuApi.GetFaCaiwuList)  // 获取FaCaiwu列表
	}
}
