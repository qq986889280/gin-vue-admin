package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/baosms/api"
	"github.com/gin-gonic/gin"
)

type BaoSmsRouter struct {
}

func (s *BaoSmsRouter) InitBaoSmsRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	plugApi := api.ApiGroupApp.BaoSmsApi
	{
		plugRouter.POST("send", plugApi.Send)
	}
}
