package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	cfg "github.com/flipped-aurora/gin-vue-admin/server/plugin/cc/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/cc/model"
)

type CcService struct{}

func (e *CcService) Do(input model.Request) (res model.TkzyRe, err error) {
	// 写你的CC业务逻辑
	// global.GVA_LOG.Info("CcService.Do at" + time.Now().Format("2006-01-02 15:04:05.000"))
	defer func() {
		if err := recover(); err != nil {
			// global.GVA_DB.Model(&model.VideoDetail{}).Where("id = ?", v.ID).Update("noticetimes", gorm.Expr("noticetimes+ ?", 1))
			fmt.Printf("Runtime panic caught: %v\n", err)
		}
	}()
	url := cfg.GlobalConfig.Curl
	if url == "" {
		url = "https://api.tiankongapi.com/api.php/provide/vod/?ac=detail"
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	re, _ := http.DefaultClient.Do(req)
	defer re.Body.Close()
	body, _ := ioutil.ReadAll(re.Body)
	fmt.Println(url)
	// fmt.Println(string(body))
	var data model.TkzyRe
	json.Unmarshal(body, &data)
	err = global.GVA_DB.Create(&data.List).Error
	if err != nil {
		fmt.Println(err.Error())
	}
	// global.GVA_LOG.Info(string(body))
	return data, nil
}
