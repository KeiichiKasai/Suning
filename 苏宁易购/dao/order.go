package dao

import "main.go/model"

func InsertOrder(uid int, gid int, name string, shipment string, receipt string, money float64) (err error) {
	sqlStr := "insert into order (status,uid,gid,gname,shipment,receipt,money) values (?,?,?,?,?,?,?)"
	_, err = DB.Exec(sqlStr, "商品派送中...", uid, gid, name, shipment, receipt, money)
	return err
}

func SelectOrder(uid int) (o []model.Order, err error) {
	sqlStr := "select oid,otime,status,uid,gid,gname,shipment,receipt,money from order where uid=?"
	err = DB.Get(&o, sqlStr, uid)
	return o, err
}

func DelOrder(uid int, oid int) (err error) {
	sqlStr := "delete from order where uid=? and oid=?"
	_, err = DB.Exec(sqlStr, uid, oid)
	return err
}

func UpdateOrder(uid int, oid int) (err error) {
	sqlStr := "update order set status=? where oid=? and uid=?"
	_, err = DB.Exec(sqlStr, "已收货", oid, uid)
	return err
}
