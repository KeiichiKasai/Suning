package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
	"main.go/service"
	"main.go/utils"
	"strconv"
	"strings"
)

func buyFromCart(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}

	//前端输入商品gid时需要用逗号间隔
	gidStr := c.PostForm("gid")
	aidStr := c.PostForm("aid")
	aid, _ := strconv.Atoi(aidStr)
	gidSlice := strings.Split(gidStr, ",")
	count := len(gidSlice)
	gid := make([]int, count)
	for i := 0; i < count; i++ {
		a, _ := strconv.Atoi(gidSlice[i])
		gid = append(gid, a)
	}
	need, err := service.ComputePrice(id, gid)
	if err != nil {
		utils.RespFail(c, "计算总金额失败")
		return
	}
	flag, err := service.MinusMoney(id, need)
	if err != nil {
		utils.RespFail(c, "扣除余额失败")
		return
	}
	if flag {
		utils.RespFail(c, "您的余额不足")
		return
	}
	address, err := service.SearchAddressByAid(aid)
	receipt := address.Place
	for i := 0; i < count; i++ {
		good, err := service.SearchGoodByGid(gid[i])
		if err != nil {
			utils.RespFail(c, "获取商品失败")
			return
		}
		gname := good.GName
		sid := good.Sid
		store, err := service.SearchStoreBySid(sid)
		if err != nil {
			utils.RespFail(c, "获取商店信息失败")
			return
		}
		shipment := store.Address
		money, err := service.ComputeOrderPrice(id, gid[i])
		if err != nil {
			utils.RespFail(c, "计算订单金额失败")
			return
		}
		err = service.CreateOrder(id, gid[i], gname, shipment, receipt, money)
		if err != nil {
			utils.RespFail(c, "生成订单失败")
			return
		}
	}
	utils.RespSuccess(c, "购买成功，订单已生成")
}

func buyInShop(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	gidStr := c.PostForm("gid")
	number := c.PostForm("number")
	aidStr := c.PostForm("aid")
	aid, _ := strconv.Atoi(aidStr)
	gid, _ := strconv.Atoi(gidStr)
	g, err := service.SearchGoodByGid(gid)
	if err != nil {
		utils.RespFail(c, "未找到该商品")
		return
	}
	need, err := service.ComputeShopPrice(gid, number)
	if err != nil {
		utils.RespFail(c, "计算金额失败")
		return
	}
	flag, err := service.MinusMoney(id, need)
	if err != nil {
		utils.RespFail(c, "扣除余额失败")
		return
	}
	if flag {
		utils.RespFail(c, "您的余额不足")
		return
	}
	address, err := service.SearchAddressByAid(aid)
	receipt := address.Place
	sid := g.Sid
	gname := g.GName
	s, err := service.SearchStoreBySid(sid)
	if err != nil {
		utils.RespFail(c, "获取商品所属商店信息失败")
	}
	shipment := s.Address
	err = service.CreateOrder(id, gid, gname, shipment, receipt, need)
	if err != nil {
		utils.RespFail(c, "生成订单失败")
		return
	}
	utils.RespSuccess(c, "生成订单成功")
}
