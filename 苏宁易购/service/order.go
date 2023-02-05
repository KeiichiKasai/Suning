package service

import (
	"main.go/dao"
	"main.go/model"
)

func CreateOrder(uid int, gid int, name string, shipment string, receipt string, money float64) (err error) {
	err = dao.InsertOrder(uid, gid, name, shipment, receipt, money)
	return err
}

func SearchOrder(uid int) (o []model.Order, err error) {
	o, err = dao.SelectOrder(uid)
	return o, err
}

func CancelOrder(uid int, oid int) (err error) {
	err = dao.DelOrder(uid, oid)
	return err
}

func ChangeOrder(uid int, oid int) (err error) {
	err = dao.UpdateOrder(uid, oid)
	return err
}
