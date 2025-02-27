// 自动生成模板FaShang
package charge

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// FaShang 结构体
type FaShang struct {
	global.GVA_MODEL
	ShangId      int       `json:"shangId" form:"shangId" gorm:"column:shang_id;comment:商户id;size:10;"`
	Image        string    `json:"image" form:"image" gorm:"column:image;comment:logo;size:200;"`
	Ptype        string    `json:"ptype" form:"ptype" gorm:"column:ptype;comment:支付方式;size:100;"`
	Contractaddr string    `json:"contractaddr" form:"contractaddr" gorm:"column:contractaddr;comment:合约地址;size:200;"`
	To           string    `json:"to" form:"to" gorm:"column:to;comment:收款地址;size:200;"`
	Url          string    `json:"url" form:"url" gorm:"column:url;comment:回调url;size:200;"`
	Net          int       `json:"net" form:"net" gorm:"column:net;comment:区块链;size:3;"`
	Fun          string    `json:"fun" form:"fun" gorm:"column:fun;comment:fun;size:200;"`
	EndTime      time.Time // 到期时间
	GetTime      time.Time // 同步时间
}

type FaShangApi struct {
	ID           uint   `gorm:"primarykey"` // 主键ID
	ShangId      int    `json:"shangId" form:"shangId" gorm:"column:shang_id;comment:商户id;size:10;"`
	Image        string `json:"image" form:"image" gorm:"column:image;comment:logo;size:200;"`
	Ptype        string `json:"ptype" form:"ptype" gorm:"column:ptype;comment:支付方式;size:100;"`
	Contractaddr string `json:"contractaddr" form:"contractaddr" gorm:"column:contractaddr;comment:合约地址;size:200;"`
	To           string `json:"to" form:"to" gorm:"column:to;comment:收款地址;size:200;"`
	Net          int    `json:"net" form:"net" gorm:"column:net;comment:区块链;size:3;"`
	Fun          string `json:"fun" form:"fun" gorm:"column:fun;comment:fun;size:200;"`
}

// TableName FaShang 表名
func (FaShang) TableName() string {
	return "fa_shang"
}
