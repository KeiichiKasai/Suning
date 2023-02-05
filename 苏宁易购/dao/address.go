package dao

import (
	"fmt"
	"main.go/model"
)

func InsertAddress(id int, place string, name string, phone int) (err error) {
	sqlStr := "insert into address (uid,place,pname,phone) values (?,?,?,?)"
	_, err = DB.Exec(sqlStr, id, place, name, phone)
	return err
}

func DeleteAddress(id int, aid int) (err error) {
	sqlStr := "delete from address where uid=? and aid=?"
	_, err = DB.Exec(sqlStr, id, aid)
	return err
}

func SelectAddress(id int) (a []model.Address, err error) {
	sqlStr := "select aid,uid,place,pname,phone from address where uid=?"
	err = DB.Select(&a, sqlStr, id)
	if err != nil {
		fmt.Printf("查询地址失败，err:%v", err)
		return nil, err
	}
	return a, err
}

func SelectAddressByAid(aid int) (a model.Address, err error) {
	sqlStr := "select aid,uid,place,pname,phone from address where aid=?"
	err = DB.Select(&a, sqlStr, aid)
	return a, err
}
