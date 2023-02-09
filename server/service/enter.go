package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/charge"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/shang"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/user"
	"github.com/flipped-aurora/gin-vue-admin/server/service/video"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	ChargeServiceGroup  charge.ServiceGroup
	ShangServiceGroup   shang.ServiceGroup
	VideoServiceGroup   video.ServiceGroup
	UserServiceGroup    user.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
