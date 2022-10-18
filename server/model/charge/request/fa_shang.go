package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/charge"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type FaShangSearch struct{
    charge.FaShang
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
