package model

type UserInformation struct {
	Id       int    `json:"id"`       //用户id
	Nickname string `json:"nickname"` //用户昵称
	Phone    int    `json:"phone"`    //手机号
	Gender   string `json:"gender"`   //性别
	Age      string `json:"age"`      //年龄
}
