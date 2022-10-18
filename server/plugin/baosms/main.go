package baosms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/baosms/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/baosms/router"
	"github.com/gin-gonic/gin"
)

type BaoSmsPlugin struct {
}

func CreateBaoSmsPlug(username string, password string, sign string) *BaoSmsPlugin {
	global.GlobalConfig.Username = username
	global.GlobalConfig.Password = password
	global.GlobalConfig.Sign = sign

	return &BaoSmsPlugin{}
}

func (*BaoSmsPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitBaoSmsRouter(group)
}

func (*BaoSmsPlugin) RouterPath() string {
	return "baosms"
}
