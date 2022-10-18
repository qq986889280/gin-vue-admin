package service

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/baosms/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/baosms/model"
)

type BaoSmsService struct{}

var ECODE = map[string]string{
	"-1": "参数错误",
	"30": "错误密码",
	"40": "账号不存在",
	"41": "余额不足",
	"43": "IP地址限制",
	"50": "内容含有敏感词",
	"51": "手机号码不正确",
}

func (e *BaoSmsService) SendService(Request model.Request) (err error) {
	// 写你的业务逻辑
	mobile := Request.Mobile
	code := Request.Code
	user := global.GlobalConfig.Username
	pass := global.GlobalConfig.Password
	sign := global.GlobalConfig.Sign
	client := http.Client{}
	params := url.Values{}
	parseURL, err := url.Parse("http://api.smsbao.com/sms")
	if err != nil {
		return err
	}
	msg := sign + "您的验证码是" + code + "。如非本人操作，请忽略本短信"
	params.Set("u", user)
	params.Set("p", pass)
	params.Set("m", mobile)
	params.Set("c", msg)
	//如果参数中有中文参数,这个方法会进行URLEncode
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	req, err := http.NewRequest(http.MethodGet, urlPathWithParams, nil)
	if err != nil {
		return err
	}
	// 添加请求头
	req.Header.Add("Content-type", "application/json;charset=utf-8")
	// req.Header.Add("header", "header😂😂")
	// 添加cookie
	// cookie1 := &http.Cookie{
	// 	Name:  "aaa",
	// 	Value: "aaa-value",
	// }
	// req.AddCookie(cookie1)
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(b))
	if string(b) != "0" {
		return errors.New(ECODE[string(b)])
	}
	// fmt.Println(urlPathWithParams)
	return nil
}
