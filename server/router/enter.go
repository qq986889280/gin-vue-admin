package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/charge"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/user"
	"github.com/flipped-aurora/gin-vue-admin/server/router/video"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Charge  charge.RouterGroup
	Shang   charge.RouterGroup
	Video   video.RouterGroup
	User    user.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
