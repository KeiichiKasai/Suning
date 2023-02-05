package service

import (
	"fmt"
	"main.go/dao"
	"main.go/model"
)

func CreateUserInformation(id int, phone string) (err error) {
	err = dao.InsertUserInformation(id, phone)
	if err != nil {
		fmt.Printf("问题出在这里，err：%v", err)
	}
	return err
}
func ReviseInformation(id int, nickname string, phone int, gender string, age int) (err error) {
	err = dao.UpdateUserInformation(id, nickname, phone, gender, age)
	return err
}
func SearchUserInformation(id int) (u model.UserInformation, err error) {
	u, err = dao.SelectUserInformation(id)
	return u, err
}
