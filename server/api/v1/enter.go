package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/charge"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/video"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	ChargeApiGroup  charge.ApiGroup
	ShangApiGroup   charge.ApiGroup
	VideoApiGroup   video.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
