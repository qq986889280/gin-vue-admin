package main

import (
	"fmt"
	"net"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	var startcat bool = false
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				if ipnet.IP.String() == "192.168.3.6" {
					startcat = true
				}
			}
		}
	}
	if startcat {
		global.GVA_VP = core.Viper() // 初始化Viper
		initialize.OtherInit()
		global.GVA_LOG = core.Zap() // 初始化zap日志库
		zap.ReplaceGlobals(global.GVA_LOG)
		global.GVA_DB = initialize.Gorm() // gorm连接数据库
		initialize.Timer()
		initialize.DBList()
		if global.GVA_DB != nil {
			initialize.RegisterTables(global.GVA_DB) // 初始化表
			// 程序结束前关闭数据库链接
			db, _ := global.GVA_DB.DB()
			defer db.Close()
		}
		core.RunWindowsServer()
	}

}
