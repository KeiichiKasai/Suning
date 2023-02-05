package dao

import "main.go/model"

func InsertUserInformation(id int, phone string) (err error) {
	sqlStr := "insert into userinformation (id,nickname,phone,gender,age) values (?,?,?,?,?)"
	_, err = DB.Exec(sqlStr, id, "未填写", phone, "未填写", 0)
	return err
}

func UpdateUserInformation(id int, nickname string, phone int, gender string, age int) (err error) {
	sqlStr := "update userinformation set nickname=?,phone=?,gender=?,age=? where id=?"
	_, err = DB.Exec(sqlStr, nickname, phone, gender, age, id)
	return err
}

func SelectUserInformation(id int) (u model.UserInformation, err error) {
	sqlStr := "select id,nickname,phone,gender,age from userinformation where id=?"
	err = DB.Get(&u, sqlStr, id)
	if err != nil {
		return model.UserInformation{}, err
	}
	return u, err
}
