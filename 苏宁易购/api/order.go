package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
	"main.go/service"
	"main.go/utils"
	"strconv"
)

func browse(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	o, err := service.SearchOrder(id)
	if err != nil {
		utils.RespFail(c, "未查询到订单")
		return
	}
	c.JSON(200, o)
}
func cancelOrder(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	oidStr := c.PostForm("oid")
	oid, _ := strconv.Atoi(oidStr)
	err = service.CancelOrder(id, oid)
	if err != nil {
		utils.RespFail(c, "取消订单失败")
		return
	}
	utils.RespSuccess(c, "取消订单成功")
}

func confirmOrder(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	oidStr := c.PostForm("oid")
	oid, _ := strconv.Atoi(oidStr)
	err = service.ChangeOrder(id, oid)
	if err != nil {
		utils.RespFail(c, "确认订单状态失败")
		return
	}
	utils.RespSuccess(c, "已签收")
}
