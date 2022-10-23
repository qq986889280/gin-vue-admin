// 自动生成模板FaCharge
package charge

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// FaCharge 结构体
type FaCharge struct {
	global.GVA_MODEL
	UserId      *int     `json:"userId" form:"userId" gorm:"column:user_id;comment:会员id;size:10;"`
	ShangId     int      `json:"shangId" form:"shangId" gorm:"column:shang_id;comment:商户id;size:10;"`
	From        string   `json:"from" form:"from" gorm:"column:from;comment:转账地址;size:255;"`
	To          string   `json:"to" form:"to" gorm:"column:to;comment:收账地址;size:255;"`
	Number      float64  `json:"number" form:"number" gorm:"column:number;comment:金额;size:22;"`
	Fee         *float64 `json:"fee" form:"fee" gorm:"column:fee;comment:转账费用;size:22;"`
	Txid        string   `json:"txid" form:"txid" gorm:"column:txid;comment:区块交易id;size:255;"`
	CreateTime  *int     `json:"createTime" form:"createTime" gorm:"column:create_time;comment:交易时间;size:10;"`
	Status      int      `json:"status" form:"status" gorm:"column:status;comment:状态:0=交易中,1=成功,2=失败;size:11;"`
	Coin        string   `json:"coin" form:"coin" gorm:"column:coin;comment:币种;size:255;"`
	Tradelog    string   `json:"tradelog" form:"tradelog" gorm:"column:tradelog;comment:日志;size:255;"`
	OrderSn     string   `json:"orderSn" form:"orderSn" gorm:"column:order_sn;comment:app订单号;size:255;"`
	Contract    string   `json:"contract" form:"contract" gorm:"column:contract;comment:;size:255;"`
	Cycle       *int     `json:"cycle" form:"cycle" gorm:"column:cycle;comment:质押周期;size:10;"`
	EndTime     *int     `json:"endTime" form:"endTime" gorm:"column:end_time;comment:到期时间;size:10;"`
	Notice      int      `json:"notice" form:"notice" gorm:"column:notice;comment:通知状态:0=待通知,1=成功,2=失败;size:11;"`
	Noticetimes int      `json:"noticetimes" form:"noticetimes" gorm:"column:noticetimes;comment:通知次数;size:11;default:0"`
	Bind        string   `json:"bind" form:"bind" gorm:"column:bind;comment:bind;size:255;"`
	FaShangId   int      `json:"fashangId" form:"fashangId" gorm:"column:fashang_id;comment:充值类型;size:11;"`
	FaShang     FaShang
}

// TableName FaCharge 表名
func (FaCharge) TableName() string {
	return "fa_charge"
}
