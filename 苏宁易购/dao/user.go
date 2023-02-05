package dao

import (
	"fmt"
	"main.go/model"
)

func SearchUserById(id int) (u model.User, err error) {
	sqlStr := "select id,username,phone,password from user where id=?"
	err = DB.Get(&u, sqlStr, id)
	if err != nil {
		fmt.Printf("用户查询失败，err:%v", err)
		fmt.Printf("id:%v,username:%v,phone:%v,password:%v", u.Id, u.Username, u.Phone, u.Password)
		return u, err
	}
	return u, err
}
func SearchUserByUsername(username string) (u model.User, err error) {
	sqlStr := "select id,username,phone,password from user where username=?"
	err = DB.Get(&u, sqlStr, username)
	if err != nil {
		fmt.Printf("用户查询失败，err:%v", err)
		fmt.Printf("id:%v,username:%v,phone:%v,password:%v", u.Id, u.Username, u.Phone, u.Password)
		return u, err
	}
	return u, err
}

func InsertUser(username string, password string, phone string) (err error) {
	sqlStr := "insert into user(username,password,phone) values (?,?,?)"
	_, err = DB.Exec(sqlStr, username, password, phone)
	if err != nil {
		fmt.Printf("用户插入失败，err:%v", err)
		return err
	}
	return err
}
func UpdatePassword(username string, newPass string) (err error) {
	sqlStr := "update user set password=? where username=?"
	_, err = DB.Exec(sqlStr, newPass, username)
	if err != nil {
		fmt.Printf("修改密码失败，err：%v", err)
	}
	return err
}

func UpdatePhone(id int, phone int) (err error) {
	sqlStr := "update user set phone=? where id=?"
	_, err = DB.Exec(sqlStr, phone, id)
	if err != nil {
		fmt.Printf("修改手机号失败，err：%v", err)
	}
	return err
}
