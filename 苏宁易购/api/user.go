package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
	"main.go/service"
	"main.go/utils"
	"strconv"
)

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	//进行简单的数据判断，一定程度上防止sql注入
	if username == "" || password == "" {
		utils.RespFail(c, "用户名或密码为空")
		return
	}
	if phone == "" {
		utils.RespFail(c, "手机号为空")
		return
	}
	if len(phone) != 11 {
		utils.RespFail(c, "手机号格式错误")
		return
	}
	if service.CheckUsername(username) {
		utils.RespFail(c, "该用户名已存在")
		return
	}
	err := service.CreateUser(username, password, phone)
	if err != nil {
		utils.RespFail(c, "创建用户失败")
		return
	}

	id := service.GetUser(username).Id

	err = service.CreateWallet(id)
	if err != nil {
		utils.RespFail(c, "用户创建钱包失败")
		return
	}
	err = service.CreateUserInformation(id, phone)
	if err != nil {
		utils.RespFail(c, "用户创建账户信息失败")
		return
	}
	utils.RespSuccess(c, "注册成功")
}
func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	//进行简单的数据判断，一定程度上防止sql注入
	if username == "" || password == "" {
		utils.RespFail(c, "用户名或密码为空")
		return
	}
	if !service.CheckUsername(username) {
		utils.RespFail(c, "用户名不存在")
		return
	}
	if !service.VerifyPassword(username, password) {
		utils.RespFail(c, "密码错误")
		return
	}
	uid := service.GetUser(username).Id
	uidString := strconv.Itoa(uid)
	c.SetCookie("user", uidString, 3600, "/", "", false, true)
	utils.RespSuccess(c, "登录成功")

}
func change(c *gin.Context) {
	username := c.PostForm("username")
	phone := c.PostForm("phone")
	newpass := c.PostForm("newpassword")
	//进行简单的数据判断，一定程度上防止sql注入
	if username == "" || newpass == "" {
		utils.RespFail(c, "用户名或新密码为空")
		return
	}
	if phone == "" {
		utils.RespFail(c, "手机号为空")
		return
	}
	if len(phone) != 11 {
		utils.RespFail(c, "手机号格式错误")
		return
	}
	if !service.CheckUsername(username) {
		utils.RespFail(c, "用户名不存在")
		return
	}
	if !service.CheckPhone(username, phone) {
		utils.RespFail(c, "手机号不匹配，无法修改密码")
	}
	err := service.ChangePassword(username, newpass)
	if err != nil {
		utils.RespFail(c, "修改密码失败")
		return
	}
	utils.RespSuccess(c, "修改密码成功")
}
func logout(c *gin.Context) {
	_, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	c.SetCookie("user", "", -1, "/", "", false, true)
	utils.RespSuccess(c, "注销成功")
}
