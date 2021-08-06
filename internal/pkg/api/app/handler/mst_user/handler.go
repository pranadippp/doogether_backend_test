package mst_user

import (
	"go/internal/pkg/api/app/context/mst_user"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	CreateHandler(ctx *gin.Context)
	LoginHandler(ctx *gin.Context)
}

//private struct
type userHandler struct {
	uc mst_user.UserUsecase
}

func NewUserHandler(uc mst_user.UserUsecase) UserHandler {
	return &userHandler{uc: uc}
}
