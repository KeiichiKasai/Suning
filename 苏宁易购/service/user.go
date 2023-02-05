package service

import (
	"fmt"
	"main.go/dao"
	"main.go/model"
	"strconv"
)

func CheckUsername(username string) (flag bool) {
	flag = true
	_, err := dao.SearchUserByUsername(username)
	if err != nil {
		flag = false
		return flag
	}
	return flag
}
func CreateUser(username string, password string, phone string) (err error) {
	err = dao.InsertUser(username, password, phone)
	return err
}
func VerifyPassword(username string, password string) (flag bool) {
	flag = true
	u, err := dao.SearchUserByUsername(username)
	if err != nil {
		fmt.Printf("在验证密码时未找到用户，重大错误:%v", err)
		return
	}
	if u.Password != password {
		flag = false
		return flag
	}
	return flag
}
func CheckPhone(username string, phone string) (flag bool) {
	flag = true
	u, err := dao.SearchUserByUsername(username)
	if err != nil {
		fmt.Printf("在验证手机号时未找到用户，重大错误:%v", err)
		return
	}
	intPhone, _ := strconv.Atoi(phone)
	if u.Phone != intPhone {
		flag = false
		return flag
	}
	return flag
}
func ChangePassword(username string, newPass string) (err error) {
	err = dao.UpdatePassword(username, newPass)
	return err
}
func ChangePhone(id int, phone int) (err error) {
	err = dao.UpdatePhone(id, phone)
	return err
}
func GetUser(username string) (u model.User) {
	u, err := dao.SearchUserByUsername(username)
	if err != nil {
		fmt.Printf("在获取时未找到用户，重大错误:%v", err)
		return
	}
	return u
}
func GetUserById(id int) (u model.User) {
	u, err := dao.SearchUserById(id)
	if err != nil {
		fmt.Printf("在获取时未找到用户，重大错误:%v", err)
		return
	}
	return u
}
func CreateWallet(id int) (err error) {
	err = dao.InsertWallet(id)
	return err
}
