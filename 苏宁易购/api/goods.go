package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
	"main.go/service"
	"main.go/utils"
	"strconv"
)

func allProduct(c *gin.Context) {
	_, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	g, err := service.GetAllGoods()
	if err != nil {
		utils.RespFail(c, "未能获取到商品数据")
		return
	}
	c.JSON(200, g)
}

func addInCart(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	gidStr := c.PostForm("gid")
	gid, _ := strconv.Atoi(gidStr)
	err = service.AddProductInCart(id, gid)
	if err != nil {
		utils.RespFail(c, "加入购物车失败")
		return
	}
	utils.RespSuccess(c, "加入购物车成功")
}

func delInCart(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	gidStr := c.PostForm("gid")
	gid, _ := strconv.Atoi(gidStr)
	err = service.DelProductInCart(id, gid)
	if err != nil {
		utils.RespFail(c, "删除该商品失败")
		return
	}
	utils.RespSuccess(c, "删除该商品成功")
}
func WatchCart(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	cart, err := service.GetAllProductInCart(id)
	if err != nil {
		utils.RespFail(c, "未获取到购物车内商品")
		return
	}
	c.JSON(200, cart)
}
