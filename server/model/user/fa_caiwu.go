// 自动生成模板FaCaiwu
package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// FaCaiwu 结构体
type FaCaiwu struct {
      global.GVA_MODEL
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`
      Account  string `json:"account" form:"account" gorm:"column:account;comment:用户名;size:255;"`
      Yprice  *float64 `json:"yprice" form:"yprice" gorm:"column:yprice;comment:原金额;"`
      Price  *float64 `json:"price" form:"price" gorm:"column:price;comment:操作金额;"`
      Nprice  *float64 `json:"nprice" form:"nprice" gorm:"column:nprice;comment:当前金额;"`
      Memo  string `json:"memo" form:"memo" gorm:"column:memo;comment:操作说明;size:200;"`
      Type  string `json:"type" form:"type" gorm:"column:type;comment:财务类型;size:5;"`
      Ptype  string `json:"ptype" form:"ptype" gorm:"column:ptype;comment:钱包类型;size:100;"`
      OrderId  string `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单号;size:100;"`
}


// TableName FaCaiwu 表名
func (FaCaiwu) TableName() string {
  return "fa_caiwu"
}

