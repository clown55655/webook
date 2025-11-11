package main

import (
	"webook/internal/web"

	"github.com/gin-gonic/gin"
)

func main() {
	//server是一个web服务器的实例
	server := gin.Default()

	u := web.NewUserHandler()
	u.RegisterRoutesV1(server.Group("/user"))
	//u.RegisterRoutes(server)
	server.Run(":8080")
}
