package charge

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FaShangRouter struct {
}

// InitFaShangRouter 初始化 FaShang 路由信息
func (s *FaShangRouter) InitFaShangRouter(Router *gin.RouterGroup) {
	faShangRouter := Router.Group("faShang").Use(middleware.OperationRecord())
	faShangRouterWithoutRecord := Router.Group("faShang")
	var faShangApi = v1.ApiGroupApp.ChargeApiGroup.FaShangApi
	{
		faShangRouter.POST("createFaShang", faShangApi.CreateFaShang)   // 新建FaShang
		faShangRouter.DELETE("deleteFaShang", faShangApi.DeleteFaShang) // 删除FaShang
		faShangRouter.DELETE("deleteFaShangByIds", faShangApi.DeleteFaShangByIds) // 批量删除FaShang
		faShangRouter.PUT("updateFaShang", faShangApi.UpdateFaShang)    // 更新FaShang
	}
	{
		faShangRouterWithoutRecord.GET("findFaShang", faShangApi.FindFaShang)        // 根据ID获取FaShang
		faShangRouterWithoutRecord.GET("getFaShangList", faShangApi.GetFaShangList)  // 获取FaShang列表
	}
}
