package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CheckLogin 检测是否存在cookie
func CheckLogin(c *gin.Context) (id int, err error) {
	var strid string
	strid, err = c.Cookie("user")
	id, _ = strconv.Atoi(strid)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	return id, err
}
