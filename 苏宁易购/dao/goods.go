package dao

import (
	"fmt"
	"main.go/model"
)

func SelectGoods() (g []model.Goods, err error) {
	sqlStr := "select gid,gname,price,sid,sales,comment from goods where gid>?"
	err = DB.Select(&g, sqlStr, 0)
	if err != nil {
		fmt.Printf("获取所有商品失败，err:%v", err)
	}
	return g, err
}

func SelectGoodInformationByGid(gid int) (g model.Goods, err error) {
	sqlStr := "select gid,gname,price,sid,sales,comment from goods where gid=?"
	err = DB.Get(&g, sqlStr, gid)
	if err != nil {
		fmt.Printf("获取商品详细信息失败，err:%v", err)
	}
	return g, err
}

func InsertCart(uid int, gid int, gname string, price float64) (err error) {
	sqlStr := "insert into cart (uid,gid,gname,price) values (?,?,?,?)"
	_, err = DB.Exec(sqlStr, uid, gid, gname, price)
	if err != nil {
		fmt.Printf("插入购物车失败，err:%v", err)
	}
	return err
}
func SelectGoodInCart(uid int, gid int) (cart model.Cart, err error) {
	sqlStr := "select uid,gid,gname,price,cnumber from cart where uid=? and gid=?"
	err = DB.Get(&cart, sqlStr, uid, gid)
	if err != nil {
		fmt.Printf("获取购物车单个商品信息失败，err:%v", err)
	}
	return cart, err
}
func SelectCart(uid int) (cart []model.Cart, err error) {
	sqlStr := "select uid,gid,gname,price,cnumber from cart where uid=?"
	err = DB.Select(&cart, sqlStr, uid)
	if err != nil {
		fmt.Printf("获取购物车所有商品信息失败，err:%v", err)
	}
	return cart, err
}
func DelGoodInCart(uid int, gid int) (err error) {
	sqlStr := "delete from cart where uid=? and gid=?"
	_, err = DB.Exec(sqlStr, uid, gid)
	if err != nil {
		fmt.Printf("删除购物车内商品失败,err:%v", err)
	}
	return err
}

func UpdateGoodNumber(uid int, gid int, number int) (err error) {
	sqlStr := "update cart set cnumber=? where uid=? and gid=?"
	_, err = DB.Exec(sqlStr, number, uid, gid)
	if err != nil {
		fmt.Printf("修改购物车内商品数目失败,err:%v", err)
	}
	return err
}

func SelectStoreBySid(sid int) (s model.Store, err error) {
	sqlStr := "select sid,name,announce,address from store where sid=?"
	err = DB.Get(&s, sqlStr, sid)
	return s, err
}
