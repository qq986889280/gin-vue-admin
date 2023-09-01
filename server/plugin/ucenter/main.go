package ucenter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ucenter/router"
	"github.com/gin-gonic/gin"
)

type UcenterPlugin struct {
}

func CreateUcenterPlug()*UcenterPlugin {
	return &UcenterPlugin{}
}

func (*UcenterPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitUcenterRouter(group)
}

func (*UcenterPlugin) RouterPath() string {
	return "Ucenter"
}
