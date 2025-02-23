package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/charge"
	chargeReq "github.com/flipped-aurora/gin-vue-admin/server/model/charge/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pay/model"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PayService struct{}

// 提交充值订单
func (e *PayService) Order(faCharge charge.FaCharge) (err error) {
	return nil
	// var ss string = faCharge.From + faCharge.To + strconv.FormatFloat(faCharge.Number, 'f', -1, 64) + faCharge.Txid + strconv.Itoa(faCharge.ShangId) + faCharge.OrderSn + faCharge.Contract + "wowowo"

	// if faCharge.Bind != utils.Hash(ss, "md5", false) {
	// 	global.GVA_LOG.Error("参数有误!")
	// 	err = errors.New("参数有误")
	// 	return
	// }
	// faCharge.Status = 0
	// faCharge.Notice = 0
	// err = global.GVA_DB.Create(&faCharge).Error
	// if err != nil {
	// 	return err
	// }
	// return nil
}

// 获取充值币种信息 PAY/INFO
func (faShangService *PayService) GetFaShangInfoList(info chargeReq.FaShangSearch) (list []charge.FaShangApi, err error) {
	// 创建db
	db := global.GVA_DB.Model(&charge.FaShang{})
	// 如果有条件搜索 下方会自动创建搜索语句
	// if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
	// 	db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	// }
	if info.ShangId > 0 {
		db = db.Where("shang_id = ?", info.ShangId)
	}
	if info.Net > 0 {
		db = db.Where("net = ?", info.Net)
	}
	err = db.Where("end_time > ?", time.Now()).Find(&list).Error
	for _, v := range list {
		if v.Fun != utils.Hash(strconv.Itoa(v.ShangId)+v.Ptype+v.Contractaddr+v.To+strconv.Itoa(v.Net)+"fvbexop", "sha1", false) {
			err = errors.New("参数有误")
			return
		}
	}
	return list, err
}

func (e *PayService) ChargeCheck() (faShang charge.FaShang, err error) {
	global.GVA_LOG.Info("ChargeCheck at" + time.Now().Format(time.RFC3339))
	// var list []charge.FaCharge
	// err := global.GVA_DB.Model(&charge.FaCharge{}).Where("created_at < ?", time.Unix(time.Now().Unix()-30, 0)).Where("status = ?", 0).Preload(clause.Associations).Find(&list).Error
	// if err != nil {
	// 	panic(err)
	// }
	// if len(list) > 0 {
	// 	for _, v := range list {
	// 		if v.FaShang.Contractaddr == v.Contract && v.FaShang.To == v.To {
	// 			is_success, _ := e.ChargeStatus(v, int16(v.FaShang.Net))
	// 			if is_success {
	// 				global.GVA_DB.Model(&charge.FaCharge{}).Where("id = ?", v.ID).Update("status", 1)
	// 			} else {
	// 				global.GVA_DB.Model(&charge.FaCharge{}).Where("id = ?", v.ID).Update("status", 2)
	// 			}
	// 		} else {
	// 			global.GVA_DB.Model(&charge.FaCharge{}).Where("id = ?", v.ID).Update("status", 2)
	// 		}
	// 	}
	// }

	// 创建db
	db := global.GVA_DB.Model(&charge.FaShang{})
	// var faShang charge.FaShang
	db.Joins("FaCharge").Where("get_time < ?", time.Now().Add(-5)).Where("end_time > ?", time.Now()).First(&faShang)
	return
	if faShang.Fun != utils.Hash(strconv.Itoa(faShang.ShangId)+faShang.Ptype+faShang.Contractaddr+faShang.To+strconv.Itoa(faShang.Net)+"fvbexop", "sha1", false) {
		err = errors.New("参数有误")
		return
	}
	global.GVA_DB.Model(&charge.FaShang{}).Where("id = ?", faShang.ID).Update("get_time", time.Now())
	// 设置请求参数
	params := RequestParams{
		ChainShortName:       "tron",
		Address:              faShang.To,
		ProtocolType:         "token_20",
		TokenContractAddress: faShang.Contractaddr,
		// StartBlockHeight:     "",
		// EndBlockHeight:       "",
		// IsFromOrTo:           "",
		Page:  "1",
		Limit: "10",
	}
	// 调用函数发送请求
	response, err := TokenTransactions(params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if response.Code == "0" {
		for _, v := range response.Data[0].TransactionList {
			if v.IsFromContract && v.TokenID == "1002000" {
				global.GVA_DB.Model(&charge.FaCharge{}).Where("txid = ?", v.TxID).Update("status", 1)
			}
		}
	}
	// return response.Data[0].TransactionList, nil
	return faShang, err
}

// 通知外部app
func (e *PayService) ChargeNotice() {
	global.GVA_LOG.Info("ChargeNotice at" + time.Now().Format("2006-01-02 15:04:05.000"))
	// err2 := global.GVA_DB.Exec("UPDATE `fa_charge` a LEFT JOIN `fa_shang` b ON a.fashang_id=b.id SET a.`stop`=1 WHERE a.notice=0 AND b.end_time < '" + time.Now().Format("2006-01-02 15:04:05.000") + "'").Error
	//  if err2 != nil {
	// 	panic(err2)
	// }
	var list []charge.FaCharge
	err := global.GVA_DB.Model(&charge.FaCharge{}).Where("status = 1 and notice=0 and noticetimes<3 and stop=0").Preload(clause.Associations).Limit(5).Find(&list).Error
	if err != nil {
		panic(err)
	}
	if len(list) > 0 {

		for _, v := range list {
			notice := model.NoticeResponse{
				OrderSn:  v.OrderSn,
				Number:   v.Number,
				From:     v.From,
				To:       v.To,
				Contract: v.Contract,
				Txid:     v.Txid,
				Sign:     utils.Hash(v.OrderSn+strconv.FormatFloat(v.Number, 'f', -1, 64)+v.From+v.To+v.Txid+strconv.Itoa(v.ShangId), "md5", false),
			}
			defer func() {
				if err := recover(); err != nil {
					global.GVA_DB.Model(&charge.FaCharge{}).Where("id = ?", v.ID).Update("noticetimes", gorm.Expr("noticetimes+ ?", 1))
					fmt.Printf("Runtime panic caught: %v\n", err)
				}
			}()
			cc, _ := json.Marshal(notice)
			payload := strings.NewReader(string(cc))
			req, _ := http.NewRequest("POST", v.FaShang.Url, payload)
			req.Header.Add("accept", "application/json")
			req.Header.Add("content-type", "application/json")
			res, _ := http.DefaultClient.Do(req)

			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			fmt.Println(string(body))
			if string(body) == "success" {
				global.GVA_DB.Model(&charge.FaCharge{}).Where("id = ?", v.ID).Update("notice", 1)
			} else {
				global.GVA_DB.Model(&charge.FaCharge{}).Where("id = ?", v.ID).Update("noticetimes", gorm.Expr("noticetimes+ ?", 1))
			}
		}
	}
}

// 检查订单状态
func (e *PayService) ChargeStatus(charge charge.FaCharge, net int16) (success bool, err error) {

	client := http.Client{}
	params := url.Values{}
	success = false
	switch net {
	case 1:
		//tron test
		// url := "https://shastapi.tronscan.org/api/transaction-info?hash=" + charge.Txid
		//tron
		url := "https://api.tronscan.org/api/transaction-info?hash=" + charge.Txid

		// payload := strings.NewReader("{\"value\":\"" + Txid + "\"}")

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("accept", "application/json")
		req.Header.Add("content-type", "application/json")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		// fmt.Println(string(body))
		type TronRe struct {
			Block             int    `json:"block"`
			Hash              string `json:"hash"`
			ContractRet       string `json:"contractRet"`
			Confirmed         bool   `json:"confirmed,omitempty"`
			TokenTransferInfo struct {
				IconURL         string `json:"icon_url"`
				Symbol          string `json:"symbol"`
				Level           string `json:"level"`
				Decimals        int    `json:"decimals"`
				Name            string `json:"name"`
				ToAddress       string `json:"to_address"`
				ContractAddress string `json:"contract_address"`
				Type            string `json:"type"`
				Vip             bool   `json:"vip"`
				TokenType       string `json:"tokenType"`
				FromAddress     string `json:"from_address"`
				AmountStr       string `json:"amount_str"`
			} `json:"tokenTransferInfo"`
		}
		var data TronRe
		json.Unmarshal(body, &data)

		// fmt.Printf("%v\n", data)
		// //  这是处理位数的代码段
		numStr := strconv.FormatFloat(math.Pow10(data.TokenTransferInfo.Decimals)*charge.Number, 'E', -1, 64)
		convertAmount, _ := decimal.NewFromString(numStr)
		// tenDecimal := big.NewFloat(math.Pow(10, float64(data.TokenTransferInfo.Decimals)))
		// _amount := new(big.Float).SetFloat64(charge.Number)
		// convertAmount, _ := new(big.Float).Mul(tenDecimal, _amount).Int(&big.Int{})
		fmt.Println(data.TokenTransferInfo.AmountStr)
		fmt.Println(convertAmount.String())
		global.GVA_LOG.Info("ChargeStatus-TRON",
			zap.String("ContractAddress", data.TokenTransferInfo.ContractAddress),
			zap.String("FromAddress", data.TokenTransferInfo.FromAddress),
			zap.String("ToAddress", data.TokenTransferInfo.ToAddress),
			zap.String("TokenType", data.TokenTransferInfo.TokenType),
			zap.String("AmountStr", data.TokenTransferInfo.AmountStr),
			zap.String("charge.Number", convertAmount.String()),
		)
		if data.ContractRet == "SUCCESS" && charge.From == data.TokenTransferInfo.FromAddress && charge.To == data.TokenTransferInfo.ToAddress && charge.Contract == data.TokenTransferInfo.ContractAddress && data.TokenTransferInfo.AmountStr == convertAmount.String() {
			success = true
			return
		}
	case 4:
		//tron test
		url := "https://shastapi.tronscan.org/api/transaction-info?hash=" + charge.Txid
		//tron
		// url := "https://api.tronscan.org/api/transaction-info?hash=" + charge.Txid

		// payload := strings.NewReader("{\"value\":\"" + Txid + "\"}")

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("accept", "application/json")
		req.Header.Add("content-type", "application/json")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		// fmt.Println(string(body))
		type TronRe struct {
			Block             int    `json:"block"`
			Hash              string `json:"hash"`
			ContractRet       string `json:"contractRet"`
			Confirmed         bool   `json:"confirmed,omitempty"`
			TokenTransferInfo struct {
				IconURL         string `json:"icon_url"`
				Symbol          string `json:"symbol"`
				Level           string `json:"level"`
				Decimals        int    `json:"decimals"`
				Name            string `json:"name"`
				ToAddress       string `json:"to_address"`
				ContractAddress string `json:"contract_address"`
				Type            string `json:"type"`
				Vip             bool   `json:"vip"`
				TokenType       string `json:"tokenType"`
				FromAddress     string `json:"from_address"`
				AmountStr       string `json:"amount_str"`
			} `json:"tokenTransferInfo"`
		}
		var data TronRe
		json.Unmarshal(body, &data)

		// fmt.Printf("%v\n", data)
		// //  这是处理位数的代码段
		numStr := strconv.FormatFloat(math.Pow10(data.TokenTransferInfo.Decimals)*charge.Number, 'E', -1, 64)
		convertAmount, _ := decimal.NewFromString(numStr)
		// tenDecimal := big.NewFloat(math.Pow(10, float64(data.TokenTransferInfo.Decimals)))
		// _amount := new(big.Float).SetFloat64(charge.Number)
		// convertAmount, _ := new(big.Float).Mul(tenDecimal, _amount).Int(&big.Int{})
		fmt.Println(data.TokenTransferInfo.AmountStr)
		fmt.Println(convertAmount.String())
		global.GVA_LOG.Info("ChargeStatus-TRON-TEST",
			zap.String("ContractAddress", data.TokenTransferInfo.ContractAddress),
			zap.String("FromAddress", data.TokenTransferInfo.FromAddress),
			zap.String("ToAddress", data.TokenTransferInfo.ToAddress),
			zap.String("TokenType", data.TokenTransferInfo.TokenType),
			zap.String("AmountStr", data.TokenTransferInfo.AmountStr),
			zap.String("charge.Number", convertAmount.String()),
		)
		if data.ContractRet == "SUCCESS" && charge.From == data.TokenTransferInfo.FromAddress && charge.To == data.TokenTransferInfo.ToAddress && charge.Contract == data.TokenTransferInfo.ContractAddress && data.TokenTransferInfo.AmountStr == convertAmount.String() {
			success = true
			return
		}
	case 2:
		//bsc
		parseURL, _ := url.Parse("https://api.bscscan.com/api")
		params.Set("module", "transaction")
		params.Set("action", "gettxreceiptstatus")
		params.Set("txhash", charge.Txid)
		params.Set("apikey", "129VB3SX26CBVPUMBVVGZV55WXUMEQKPMP")
		//如果参数中有中文参数,这个方法会进行URLEncode
		parseURL.RawQuery = params.Encode()
		urlPathWithParams := parseURL.String()
		req, err := http.NewRequest(http.MethodGet, urlPathWithParams, nil)
		if err != nil {
			return success, err
		}
		// 发送请求
		resp, err := client.Do(req)
		if err != nil {
			return success, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return success, err
		}
		// {"status":"1","message":"OK","result":{"status":""}}
		type Bep20Re struct {
			Status  string `json:"status"`
			Message string `json:"message"`
			Result  struct {
				Status string `json:"status"`
			} `json:"result"`
		}
		var data Bep20Re
		json.Unmarshal(body, &data)
		fmt.Println(string(body))
		if data.Result.Status == "1" {
			success = true
			return success, err
		}
	case 5:
		//bsc testnet
		parseURL, _ := url.Parse("https://api-testnet.bscscan.com/api")
		params.Set("module", "transaction")
		params.Set("action", "gettxreceiptstatus")
		params.Set("txhash", charge.Txid)
		params.Set("apikey", "129VB3SX26CBVPUMBVVGZV55WXUMEQKPMP")
		//如果参数中有中文参数,这个方法会进行URLEncode
		parseURL.RawQuery = params.Encode()
		urlPathWithParams := parseURL.String()
		req, err := http.NewRequest(http.MethodGet, urlPathWithParams, nil)
		if err != nil {
			return success, err
		}
		// 发送请求
		resp, err := client.Do(req)
		if err != nil {
			return success, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return success, err
		}
		// {"status":"1","message":"OK","result":{"status":""}}
		type Bep20Re struct {
			Status  string `json:"status"`
			Message string `json:"message"`
			Result  struct {
				Status string `json:"status"`
			} `json:"result"`
		}
		var data Bep20Re
		json.Unmarshal(body, &data)
		fmt.Println(string(body))
		if data.Result.Status == "1" {
			success = true
			return success, err
		}
	default:
		return success, err
	}
	return success, err
}
