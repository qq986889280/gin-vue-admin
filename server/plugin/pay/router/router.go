package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pay/api"
	"github.com/gin-gonic/gin"
)

type PayRouter struct {
}

func (s *PayRouter) InitPayRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	PayApi := api.ApiGroupApp.PayApi
	{
		plugRouter.POST("order", PayApi.Order)
	}
}
