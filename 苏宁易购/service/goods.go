package service

import (
	"main.go/dao"
	"main.go/model"
)

func GetAllGoods() (g []model.Goods, err error) {
	g, err = dao.SelectGoods()
	return g, err
}

func AddProductInCart(uid int, gid int) (err error) {
	m, err := dao.SelectGoodInCart(uid, gid)
	if err == nil {
		before := m.CNumber
		now := before + 1
		err = dao.UpdateGoodNumber(uid, gid, now)
		if err != nil {
			return err
		}
		return err
	}
	var g model.Goods
	g, err = dao.SelectGoodInformationByGid(gid)
	if err != nil {
		return err
	}
	gname := g.GName
	price := g.Price
	err = dao.InsertCart(uid, gid, gname, price)
	return err
}
func DelProductInCart(uid int, gid int) (err error) {
	cart, err := dao.SelectGoodInCart(uid, gid)
	if err != nil {
		return err
	}
	number := cart.CNumber
	now := number - 1
	if now == 0 {
		err = dao.DelGoodInCart(uid, gid)
		if err != nil {
			return err
		}
		return err
	}
	err = dao.UpdateGoodNumber(uid, gid, now)
	return err
}

func GetAllProductInCart(uid int) (cart []model.Cart, err error) {
	cart, err = dao.SelectCart(uid)
	return cart, err
}

func SearchGoodByGid(gid int) (g model.Goods, err error) {
	g, err = dao.SelectGoodInformationByGid(gid)
	return g, err
}

func SearchStoreBySid(sid int) (s model.Store, err error) {
	s, err = dao.SelectStoreBySid(sid)
	return s, err
}
