package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FaUserLevelRouter struct {
}

// InitFaUserLevelRouter 初始化 FaUserLevel 路由信息
func (s *FaUserLevelRouter) InitFaUserLevelRouter(Router *gin.RouterGroup) {
	faUserLevelRouter := Router.Group("faUserLevel").Use(middleware.OperationRecord())
	faUserLevelRouterWithoutRecord := Router.Group("faUserLevel")
	var faUserLevelApi = v1.ApiGroupApp.UserApiGroup.FaUserLevelApi
	{
		faUserLevelRouter.POST("createFaUserLevel", faUserLevelApi.CreateFaUserLevel)   // 新建FaUserLevel
		faUserLevelRouter.DELETE("deleteFaUserLevel", faUserLevelApi.DeleteFaUserLevel) // 删除FaUserLevel
		faUserLevelRouter.DELETE("deleteFaUserLevelByIds", faUserLevelApi.DeleteFaUserLevelByIds) // 批量删除FaUserLevel
		faUserLevelRouter.PUT("updateFaUserLevel", faUserLevelApi.UpdateFaUserLevel)    // 更新FaUserLevel
	}
	{
		faUserLevelRouterWithoutRecord.GET("findFaUserLevel", faUserLevelApi.FindFaUserLevel)        // 根据ID获取FaUserLevel
		faUserLevelRouterWithoutRecord.GET("getFaUserLevelList", faUserLevelApi.GetFaUserLevelList)  // 获取FaUserLevel列表
	}
}
