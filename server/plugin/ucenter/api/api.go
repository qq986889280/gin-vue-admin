package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ucenter/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ucenter/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UcenterApi struct{}

// @Tags Ucenter
// @Summary 登录
// @Produce  application/json
// @Param    data  body model.Request true  "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /Ucenter/Login [post]
func (p *UcenterApi) Login(c *gin.Context) {

	var input model.Request
	_ = c.ShouldBindJSON(&input)
	u := &user.FaUser{Username: input.Username, Password: input.Password}
	if user, err := service.ServiceGroupApp.UserLogin(u); err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		if user.Status != 1 {
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		p.TokenNext(c, *user)
		// response.OkWithDetailed(user, "成功", c)
	}
}

// TokenNext 登录以后签发jwt
func (p *UcenterApi) TokenNext(c *gin.Context, user user.FaUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		ID:       user.ID,
		NickName: user.Nickname,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	// if !global.GVA_CONFIG.System.UseMultipoint {
	response.OkWithDetailed(model.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
	// return
	// }
}

// @Tags Ucenter
// @Summary 用户信息
// @Produce  application/json
// @Param    data  body model.Request true  "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /Ucenter/Login [post]
func (p *UcenterApi) Info(c *gin.Context) {
	uid := utils.GetUserID(c)
	var user user.FaUser
	if err := global.GVA_DB.Where("id = ?", uid).First(&user).Error; err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}
