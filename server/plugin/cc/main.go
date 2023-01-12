package cc

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cc/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cc/router"
	"github.com/gin-gonic/gin"
)

type CcPlugin struct {
}

func CreateCcPlug(curl string) *CcPlugin {
	global.GlobalConfig.Curl = curl

	return &CcPlugin{}
}

func (*CcPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitCcRouter(group)
}

func (*CcPlugin) RouterPath() string {
	return "cc"
}
