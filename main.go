package main

import (
	"webook/internal/web"

	"github.com/gin-gonic/gin"
)

func mian() {
	//server是一个web服务器的实例
	server := gin.Default()
	
	u := &web.UserHandler{}
	u.RegisterRoutesV1(server.Group("/user"))
	//u.RegisterRoutes(server)
	server.Run(":8080")
}
