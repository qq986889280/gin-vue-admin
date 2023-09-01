package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ucenter/api"
	"github.com/gin-gonic/gin"
)

type UcenterRouter struct {
}

func (s *UcenterRouter) InitUcenterRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	ucenter := api.ApiGroupApp.UcenterApi
	{
		plugRouter.POST("Login", ucenter.Login)
		plugRouter.GET("Info", ucenter.Info)
	}
}
