package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
	"main.go/service"
	"main.go/utils"
)

func addAddress(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	place := c.PostForm("address")
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	if place == "" {
		utils.RespFail(c, "收货地址不能为空")
		return
	}
	if name == "" {
		utils.RespFail(c, "收货人不能为空")
		return
	}
	if phone == "" {
		utils.RespFail(c, "手机号不能为空")
		return
	}
	err = service.CreateAddress(id, place, name, phone)
	if err != nil {
		utils.RespFail(c, "创建地址失败")
		return
	}
	utils.RespSuccess(c, "创建地址成功")
}

func delAddress(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	aid := c.PostForm("aid")
	if aid == "" {
		utils.RespFail(c, "aid不能为空")
		return
	}
	err = service.DelAddress(id, aid)
	if err != nil {
		utils.RespFail(c, "删除地址失败")
		return
	}
	utils.RespSuccess(c, "删除地址成功")
}

func viewAddress(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	address, err := service.GetAddress(id)
	if err != nil {
		utils.RespFail(c, "未查询到地址")
		return
	}
	c.JSON(200, address)
}
