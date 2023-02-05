package model

type Address struct {
	Aid   int    `json:"aid"`   //地址id
	Uid   int    `json:"uid"`   //用户id
	Place string `json:"place"` //收货地址
	PName string `json:"pname"` //收货人
	Phone int    `json:"phone"` //手机号
}
