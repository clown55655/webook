package web

import (
	"net/http"
	"webook/internal/domain"
	"webook/internal/service"

	regexp "github.com/dlclark/regexp2"

	"github.com/gin-gonic/gin"
)

// UserHandler 我准备在它上面定义跟用户有关的路由
type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	const (
		emailRegexPattern    = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
		passwordRefexPattern = `^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])[a-zA-Z0-9!@#$%^&*]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRefexPattern, regexp.None)
	return &UserHandler{emailExp, passwordExp}
}

type ArticleHandler struct {
}

func (u *UserHandler) RegisterRoutesV1(ug *gin.RouterGroup) {
	ug.POST("/signup", u.Signup)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")

	ug.POST("/signup", u.Signup)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
}

func (u *UserHandler) Signup(context *gin.Context) {
	type SignUpReq struct {
		Email           string
		Password        string
		ComfirmPassword string
	}

	//bind获取请求
	var req SignUpReq
	if err := context.Bind(&req); err != nil {
		return
	}

	//邮件校验
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		context.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		context.String(http.StatusOK, "邮箱格式不正确")
		return
	}

	//密码两次不一致
	if req.Password != req.Password {
		context.String(http.StatusOK, "两次密码不一致")
		return
	}

	//密码校验
	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		context.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		context.String(http.StatusOK, "密码格式不正确")
		return
	}

	//调用service
	err = u.svc.SignUp(context, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {

	}
	context.String(http.StatusOK, "注册成功")
}

func (u *UserHandler) Login(context *gin.Context) {

}

func (u *UserHandler) Edit(context *gin.Context) {

}

func (u *UserHandler) Profile(context *gin.Context) {

}
