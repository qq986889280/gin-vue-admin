package pay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pay/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pay/router"
	"github.com/gin-gonic/gin"
)

type PayPlugin struct {
}

func CreatePayPlug(Title string) *PayPlugin {
	global.GlobalConfig.Title = Title

	return &PayPlugin{}
}

func (*PayPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitPayRouter(group)
}

func (*PayPlugin) RouterPath() string {
	return "pay"
}
