package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
	"main.go/service"
	"main.go/utils"
	"strconv"
)

func setting(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	nickname := c.PostForm("nickname")
	phone := c.PostForm("phone")
	gender := c.PostForm("gender")
	age := c.PostForm("age")
	intAge, _ := strconv.Atoi(age)
	intPhone, _ := strconv.Atoi(phone)
	if gender == "" {
		gender = "未填写"
	}
	if age == "" {
		intAge = 0
	}
	err = service.ReviseInformation(id, nickname, intPhone, gender, intAge)
	if err != nil {
		utils.RespFail(c, "修改个人信息失败")
		return
	}
	err = service.ChangePhone(id, intPhone)
	if err != nil {
		utils.RespFail(c, "修改用户信息失败")
		return
	}
	utils.RespSuccess(c, "修改成功")
}

func information(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	i, err := service.SearchUserInformation(id)
	if err != nil {
		utils.RespFail(c, "查询失败")
		return
	}
	c.JSON(200, i)
}
