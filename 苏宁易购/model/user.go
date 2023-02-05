package model

type User struct {
	Id       int    `json:"id"`       //用户id
	Username string `json:"username"` //用户名
	Password string `json:"password"` //密码
	Phone    int    `json:"phone"`    //手机号
}
