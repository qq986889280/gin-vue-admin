package model
type Request struct {
        From string // from 
        To string // to 
        Coin string // coin 
        Number string // 金额 
        Txid string // hash 
        Shang_id string // 商家id 
        Order_sn string // 订单号 
        Contract string // 合约地址 
        Bind string // bind 
}

type Response struct {
        Ptype string // ptype 
        Contractaddr string // 合约地址 
}
