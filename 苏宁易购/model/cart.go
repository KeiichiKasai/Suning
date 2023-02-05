package model

type Cart struct {
	Uid     int     `json:"uid"`     //用户id
	Gid     int     `json:"gid"`     //商品id
	GName   string  `json:"gname"`   //商品名称
	Price   float64 `json:"price"`   //商品单价
	CNumber int     `json:"cnumber"` //数量
}
