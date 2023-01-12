package model

type RequestPayCheck struct {
	ShangId int    `json:"shangId" form:"shangId" gorm:"column:shang_id;comment:商户id;size:10;"`
	OrderSn string `json:"orderSn" form:"orderSn" gorm:"column:order_sn;comment:app订单号;size:255;"`
}

type Response struct {
	Ptype        string // ptype
	Contractaddr string // 合约地址
}

type ResponseInfo struct {
	ShangId      int
	Ptype        string
	Contractaddr string
	To           string
	Net          int
}

type NoticeResponse struct {
	OrderSn  string
	Number   float64
	From     string
	To       string
	Contract string
	Txid     string
	Sign     string
}
