package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pay/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pay/service"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type PayRouter struct {
}

func (s *PayRouter) InitPayRouter(Router *gin.RouterGroup) {
	plugRouter := Router
	PayApi := api.ApiGroupApp.PayApi
	{
		plugRouter.POST("order", PayApi.Order)
		plugRouter.POST("info", PayApi.GetFaShangList)
		plugRouter.POST("check", PayApi.PayCheck)
		plugRouter.POST("chargecheck", PayApi.ChargeCheck)
	}
	var option []cron.Option
	if global.GVA_CONFIG.Timer.WithSeconds {
		option = append(option, cron.WithSeconds())
	}
	global.GVA_Timer.AddTaskByFunc("TestTimer", "@every 30s", service.ServiceGroupApp.ChargeCheck, option...)
	global.GVA_Timer.AddTaskByFunc("TestTimer", "@every 10s", service.ServiceGroupApp.ChargeNotice, option...)

}
