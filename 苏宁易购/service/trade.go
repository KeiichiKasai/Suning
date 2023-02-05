package service

import (
	"main.go/dao"
	"main.go/model"
	"strconv"
)

func AddMoney(id int, money float64) (err error) {
	var w model.Wallet
	w, err = dao.SearchWalletById(id)
	if err != nil {
		return err
	}
	before := w.Money
	now := before + money
	err = dao.UpdateWallet(id, now)
	if err != nil {
		return err
	}
	return err
}
func MinusMoney(id int, money float64) (flag bool, err error) {
	flag = false
	var w model.Wallet
	w, err = dao.SearchWalletById(id)
	if err != nil {
		return flag, err
	}
	before := w.Money
	now := before - money
	if now < 0 {
		flag = true
		return flag, nil
	}
	err = dao.UpdateWallet(id, now)
	if err != nil {
		return flag, err
	}
	return flag, err
}
func GetWallet(id int) (w model.Wallet, err error) {
	w, err = dao.SearchWalletById(id)
	return w, err
}

// ComputePrice 计算购物车内商品金额总和
func ComputePrice(uid int, gid []int) (price float64, err error) {
	price = 0.00
	for i := 1; i < len(gid); i++ {
		c, err := dao.SelectGoodInCart(uid, gid[i])
		if err != nil {
			return price, err
		}
		count, err := strconv.ParseFloat(strconv.Itoa(c.CNumber), 64)
		price += c.Price * count
	}
	return price, err
}

// ComputeOrderPrice 计算购物车内一种物品金额
func ComputeOrderPrice(uid int, gid int) (price float64, err error) {
	g, err := dao.SelectGoodInCart(uid, gid)
	if err != nil {
		return 0, err
	}
	price = 0.00
	count, err := strconv.ParseFloat(strconv.Itoa(g.CNumber), 64)
	price = g.Price * count
	return price, err
}

// ComputeShopPrice 计算商城内物品金额
func ComputeShopPrice(gid int, number string) (price float64, err error) {
	g, err := dao.SelectGoodInformationByGid(gid)
	if err != nil {
		return 0, err
	}
	price = 0.00
	count, err := strconv.ParseFloat(number, 64)
	price = g.Price * count
	return price, err
}
