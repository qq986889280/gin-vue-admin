package charge

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FaChargeRouter struct {
}

// InitFaChargeRouter 初始化 FaCharge 路由信息
func (s *FaChargeRouter) InitFaChargeRouter(Router *gin.RouterGroup) {
	faChargeRouter := Router.Group("faCharge").Use(middleware.OperationRecord())
	faChargeRouterWithoutRecord := Router.Group("faCharge")
	var faChargeApi = v1.ApiGroupApp.ChargeApiGroup.FaChargeApi
	{
		faChargeRouter.POST("createFaCharge", faChargeApi.CreateFaCharge)   // 新建FaCharge
		faChargeRouter.DELETE("deleteFaCharge", faChargeApi.DeleteFaCharge) // 删除FaCharge
		faChargeRouter.DELETE("deleteFaChargeByIds", faChargeApi.DeleteFaChargeByIds) // 批量删除FaCharge
		faChargeRouter.PUT("updateFaCharge", faChargeApi.UpdateFaCharge)    // 更新FaCharge
	}
	{
		faChargeRouterWithoutRecord.GET("findFaCharge", faChargeApi.FindFaCharge)        // 根据ID获取FaCharge
		faChargeRouterWithoutRecord.GET("getFaChargeList", faChargeApi.GetFaChargeList)  // 获取FaCharge列表
	}
}
