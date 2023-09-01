package user

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FaUserRouter struct {
}

// InitFaUserRouter 初始化 FaUser 路由信息
func (s *FaUserRouter) InitFaUserRouter(Router *gin.RouterGroup) {
	faUserRouter := Router.Group("faUser").Use(middleware.OperationRecord())
	faUserRouterWithoutRecord := Router.Group("faUser")
	var faUserApi = v1.ApiGroupApp.UserApiGroup.FaUserApi
	{
		faUserRouter.POST("createFaUser", faUserApi.CreateFaUser)             // 新建FaUser
		faUserRouter.DELETE("deleteFaUser", faUserApi.DeleteFaUser)           // 删除FaUser
		faUserRouter.DELETE("deleteFaUserByIds", faUserApi.DeleteFaUserByIds) // 批量删除FaUser
		faUserRouter.PUT("updateFaUser", faUserApi.UpdateFaUser)              // 更新FaUser
		faUserRouter.POST("userCharge", faUserApi.UserCharge)                 // 充值FaUser
		faUserRouter.GET("userTeam", faUserApi.UserTeam)                      // 团队FaUser
	}
	{
		faUserRouterWithoutRecord.GET("findFaUser", faUserApi.FindFaUser)       // 根据ID获取FaUser
		faUserRouterWithoutRecord.GET("getFaUserList", faUserApi.GetFaUserList) // 获取FaUser列表
	}
}
