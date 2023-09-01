package service

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type UcenterService struct{}

func (e *UcenterService) UserLogin(u *user.FaUser) (res *user.FaUser, err error) {
	// 写你的业务逻辑
	var user user.FaUser
	err = global.GVA_DB.Where("username = ?", u.Username).Preload("FaUserLevel").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}
