package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/model"
	"main.go/service"
	"main.go/utils"
)

func Oauth(ctx *gin.Context) {
	var err error
	// 获取 code
	var code = ctx.Query("code")
	// 通过 code, 获取 token
	fmt.Printf("code:%s\n", code)
	var tokenAuthUrl = service.GetTokenAuthUrl(code) //获取token所在的url
	var token *model.Token
	if token, err = service.GetToken(tokenAuthUrl); err != nil {
		utils.RespFail(ctx, "wrong")
		return
	}
	ctx.JSON(200, *token)
	// 通过token，获取用户信息
	var userInfo map[string]interface{}
	userInfo, err = service.GetUserInfo(token)
	if err != nil {
		utils.RespFail(ctx, "wrong")
		return
	}
	user := userInfo["login"]
	ctx.JSON(200, gin.H{
		"用户名:": user,
	})
	ctx.SetCookie("user", "999", 3600, "/", "", false, true)
}
