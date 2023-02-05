package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
	"main.go/service"
	"main.go/utils"
	"strconv"
)

func like(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	gidStr := c.PostForm("gid")
	gid, err := strconv.Atoi(gidStr)
	err = service.AddLike(id, gid)
	if err != nil {
		utils.RespFail(c, "收藏失败")
		return
	}
	utils.RespSuccess(c, "收藏成功")
}

func unlike(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	gidStr := c.PostForm("gid")
	gid, err := strconv.Atoi(gidStr)
	err = service.MinusLike(id, gid)
	if err != nil {
		utils.RespFail(c, "取消收藏失败")
		return
	}
	utils.RespSuccess(c, "已取消")
}
