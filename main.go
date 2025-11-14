package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"webook/internal/web"
)

func main() {
	//server是一个web服务器的实例
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PATCH"},
		AllowHeaders:     []string{"authorization", "Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "http://localhost")
		},
		MaxAge: 12 * time.Hour,
	}))

	u := web.NewUserHandler()
	u.RegisterRoutesV1(server.Group("/user"))
	//u.RegisterRoutes(server)
	server.Run(":8080")
}
