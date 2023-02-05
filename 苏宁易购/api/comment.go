package api

import (
	"github.com/gin-gonic/gin"
	"main.go/api/middleware"
	"main.go/service"
	"main.go/utils"
	"strconv"
)

func comment(c *gin.Context) {
	id, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	strGid := c.PostForm("gid")
	content := c.PostForm("content")
	gid, _ := strconv.Atoi(strGid)
	username := service.GetUserById(id).Username
	err = service.CreateComment(gid, username, content)
	if err != nil {
		utils.RespFail(c, "评论失败")
	}
	err = service.AddComment(gid)
	if err != nil {
		utils.RespFail(c, "商品评论数增加失败")
	}
	utils.RespSuccess(c, "评论成功")
}
func viewComment(c *gin.Context) {
	_, err := middleware.CheckLogin(c)
	if err != nil {
		utils.RespUnauthorized(c)
		return
	}
	gidStr := c.PostForm("gid")
	gid, _ := strconv.Atoi(gidStr)
	comments, err := service.GetComments(gid)
	if err != nil {
		utils.RespFail(c, "未查询到相关评论")
		return
	}
	c.JSON(200, comments)
}
