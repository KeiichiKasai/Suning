package dao

import (
	"fmt"
	"main.go/model"
)

func InsertWallet(id int) (err error) {
	sqlStr := "insert into wallet(id,money) values (?,?)"
	_, err = DB.Exec(sqlStr, id, 0.00)
	if err != nil {
		fmt.Printf("钱包插入失败，err:%v", err)
		return err
	}
	return err
}

func SearchWalletById(id int) (w model.Wallet, err error) {
	sqlStr := "select id,money from wallet where id=?"
	err = DB.Get(&w, sqlStr, id)
	if err != nil {
		fmt.Printf("钱包查询失败，err:%v", err)
	}
	return w, err
}

func UpdateWallet(id int, money float64) (err error) {
	sqlStr := "update wallet set money=? where id=?"
	_, err = DB.Exec(sqlStr, money, id)
	if err != nil {
		fmt.Printf("修改金额失败，err:%v", err)
	}
	return err
}
