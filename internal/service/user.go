package service

import (
	"webook/internal/domain"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func (svc *UserService) SignUp(context *gin.Context, u domain.User) error {

	return nil
}
