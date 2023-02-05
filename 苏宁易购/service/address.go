package service

import (
	"main.go/dao"
	"main.go/model"
	"strconv"
)

func CreateAddress(id int, place string, name string, phone string) (err error) {
	intPhone, _ := strconv.Atoi(phone)
	err = dao.InsertAddress(id, place, name, intPhone)
	return err
}

func DelAddress(id int, aid string) (err error) {
	intAid, _ := strconv.Atoi(aid)
	err = dao.DeleteAddress(id, intAid)
	return err
}

func GetAddress(id int) (a []model.Address, err error) {
	a, err = dao.SelectAddress(id)
	return a, err
}

func SearchAddressByAid(aid int) (a model.Address, err error) {
	a, err = dao.SelectAddressByAid(aid)
	return a, err
}
