package main

import (
	"main.go/api"
	"main.go/api/middleware"
	"main.go/dao"
)

func main() {
	middleware.ViperSetup()
	dao.InitDatabase()
	api.InitRouter()
}
