package utils

import "github.com/gin-gonic/gin"

func RespSuccess(c *gin.Context, message string) {
	c.JSON(200, gin.H{
		"status":  "200",
		"message": message,
	})
}

func RespFail(c *gin.Context, message string) {
	c.JSON(500, gin.H{
		"status":  "500",
		"message": message,
	})
}

func RespUnauthorized(c *gin.Context) {
	c.JSON(401, gin.H{
		"status":  "401",
		"message": "未登录",
	})
}
