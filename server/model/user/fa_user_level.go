// 自动生成模板FaUserLevel
package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// FaUserLevel 结构体
type FaUserLevel struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:等级名称;size:100;"`
      Ztnum  *int `json:"ztnum" form:"ztnum" gorm:"column:ztnum;comment:直推;size:10;"`
      Tdnum  *int `json:"tdnum" form:"tdnum" gorm:"column:tdnum;comment:团队;size:10;"`
      N1  *int `json:"n1" form:"n1" gorm:"column:n1;comment:分红代数;size:10;"`
      Rate  *float64 `json:"rate" form:"rate" gorm:"column:rate;comment:分红比率;size:19;"`
}


// TableName FaUserLevel 表名
func (FaUserLevel) TableName() string {
  return "fa_user_level"
}

