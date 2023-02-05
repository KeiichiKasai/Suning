package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
	"main.go/service"
	"main.go/utils"
	"strconv"
)

func charge(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	strMoney := c.PostForm("money")
	money, _ := strconv.ParseFloat(strMoney, 64)
	err = service.AddMoney(id, money)
	if err != nil {
		utils.RespFail(c, "充值失败")
		return
	}
	utils.RespSuccess(c, "充值成功")
}
func view(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	w, err := service.GetWallet(id)
	if err != nil {
		utils.RespFail(c, "未找到钱包")
		return
	}
	money := w.Money
	strMoney := strconv.FormatFloat(money, 'f', 2, 64)
	utils.RespSuccess(c, "钱包余额为:"+strMoney)
}
