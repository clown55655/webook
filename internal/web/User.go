package web

import (
	"github.com/gin-gonic/gin"
)

// UserHandler 我准备在它上面定义跟用户有关的路由
type UserHandler struct {
}

type ArticleHandler struct {
}

func (u *UserHandler) RegisterRoutesV1(ug *gin.RouterGroup) {
	ug.POST("/users/signup", u.Signup)
	ug.POST("/users/login", u.Login)
	ug.POST("/users/edit", u.Edit)
	ug.GET("users/profile", u.Profile)
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")

	ug.POST("/users/signup", u.Signup)
	ug.POST("/users/login", u.Login)
	ug.POST("/users/edit", u.Edit)
	ug.GET("users/profile", u.Profile)
}

/*function (u *UserHandler)  Re{}*/
func (u *UserHandler) Signup(context *gin.Context) {

}

func (u *UserHandler) Login(context *gin.Context) {

}

func (u *UserHandler) Edit(context *gin.Context) {

}

func (u *UserHandler) Profile(context *gin.Context) {

}
