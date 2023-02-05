package model

import "time"

type Order struct {
	Oid      int       `json:"oid"`      //订单id
	Otime    time.Time `json:"otime"`    //订单生成时间
	Status   string    `json:"status"`   //状态
	Uid      int       `json:"uid"`      //用户id
	Gid      int       `json:"gid"`      //商品id
	GName    string    `json:"gname"`    //商品名称
	Shipment string    `json:"shipment"` //发货地
	Receipt  string    `json:"receipt"`  //收货地
	Money    float64   `json:"money"`    //订单金额
}
