package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cc/api"
	"github.com/gin-gonic/gin"
)

type CcRouter struct {
}

func (s *CcRouter) InitCcRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.CcApi
	{
		plugRouter.GET("cai", plugApi.Cai)
	}
}
