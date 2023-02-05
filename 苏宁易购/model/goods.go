package model

type Goods struct {
	Gid     int     `json:"gid"`     //商品id
	GName   string  `json:"gname"`   //商品名称
	Price   float64 `json:"price"`   //价格
	Sid     int     `json:"sid"`     //所属店铺id
	Sales   int     `json:"sales"`   //销量
	Comment int     `json:"comment"` //评论数
}
type Store struct {
	Sid      int    `json:"sid"`      //店铺id
	Name     string `json:"name"`     //店铺名称
	Announce string `json:"announce"` //店铺公告
	Address  string `json:"address"`  //发货地址
}
