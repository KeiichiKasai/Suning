package model

type Wallet struct {
	Id    int     `json:"id"`    //钱包id
	Money float64 `json:"money"` //余额
}
