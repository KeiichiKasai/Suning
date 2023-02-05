package model

import "time"

type Comment struct {
	Cid      int       `json:"cid"`      //评论id
	Gid      int       `json:"gid"`      //评论的商品
	Username string    `json:"username"` //用户昵称
	Ctime    time.Time `json:"ctime"`    //评论时间
	Content  string    `json:"content"`  //内容
}
