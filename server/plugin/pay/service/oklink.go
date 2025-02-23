package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

// 定义请求参数的结构体
type RequestParams struct {
	ChainShortName       string `json:"chainShortName"`
	Address              string `json:"address"`
	ProtocolType         string `json:"protocolType"`
	TokenContractAddress string `json:"tokenContractAddress"`
	StartBlockHeight     string `json:"startBlockHeight"`
	EndBlockHeight       string `json:"endBlockHeight"`
	IsFromOrTo           string `json:"isFromOrTo"`
	Page                 string `json:"page"`
	Limit                string `json:"limit"`
}

type ResponseData struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Limit           string `json:"limit"`
		Page            string `json:"page"`
		TotalPage       string `json:"totalPage"`
		TransactionList []struct {
			TxID                 string `json:"txId"`
			BlockHash            string `json:"blockHash"`
			Height               string `json:"height"`
			TransactionTime      string `json:"transactionTime"`
			From                 string `json:"from"`
			To                   string `json:"to"`
			TokenContractAddress string `json:"tokenContractAddress"`
			TokenID              string `json:"tokenId"`
			Amount               string `json:"amount"`
			Symbol               string `json:"symbol"`
			IsFromContract       bool   `json:"isFromContract"`
			IsToContract         bool   `json:"isToContract"`
		} `json:"transactionList"`
	} `json:"data"`
}

// 定义返回数据的结构体（根据实际返回数据格式调整）
// type ResponseData struct {
// 	Code string      `json:"code"`
// 	Msg  string      `json:"msg"`
// 	Data interface{} `json:"data"`
// }

// 发送GET请求并获取响应
func TokenTransactions(params RequestParams) (*ResponseData, error) {
	// 构造请求URL
	baseURL := "https://www.oklink.com/api/v5/explorer/address/token-transaction-list"
	queryParams := url.Values{}
	queryParams.Add("chainShortName", params.ChainShortName)
	queryParams.Add("address", params.Address)
	queryParams.Add("protocolType", params.ProtocolType)
	queryParams.Add("tokenContractAddress", params.TokenContractAddress)
	queryParams.Add("startBlockHeight", params.StartBlockHeight)
	// queryParams.Add("endBlockHeight", params.EndBlockHeight)
	// queryParams.Add("isFromOrTo", params.IsFromOrTo)
	queryParams.Add("page", params.Page)
	queryParams.Add("limit", params.Limit)

	// 拼接完整URL
	url := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())

	// url := "https://www.oklink.com/api/v5/explorer/blockchain/summary?chainShortName=ETH"
	method := "GET"
	apiKey := "79a6bd3a-8b5b-430b-8e6c-98b87bed9924"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to parse response JSON: %v", err)
	}
	req.Header.Add("Ok-Access-Key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response JSON: %v", err)
	}

	// 发送GET请求
	// resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// 解析JSON响应
	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response JSON: %v", err)
	}

	return &responseData, nil
}

func TestAdd(t *testing.T) {

	// 设置请求参数
	params := RequestParams{
		ChainShortName:       "tron",
		Address:              "TXohHjqnizydh2YuRizKHjcWfNWnG3JPxN",
		ProtocolType:         "token_20",
		TokenContractAddress: "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
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

	// 打印返回结果
	fmt.Printf("Response Code: %s\n", response.Code)
	fmt.Printf("Response Message: %s\n", response.Msg)
	fmt.Printf("Response Data: %v\n", response.Data[0].TransactionList)
}
