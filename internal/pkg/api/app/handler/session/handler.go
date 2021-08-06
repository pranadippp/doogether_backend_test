package session

import (
	"go/internal/pkg/api/app/context/session"

	"github.com/gin-gonic/gin"
)

type SessionHandler interface {
	CreateHandler(ctx *gin.Context)
	UpdateHandler(ctx *gin.Context)
	DeleteHandler(ctx *gin.Context)
	FindDetailHandler(ctx *gin.Context)
	FindAllHandler(ctx *gin.Context)
}

//private struct
type sessionHandler struct {
	uc session.SessionUsecase
}

func NewSessionHandler(uc session.SessionUsecase) SessionHandler {
	return &sessionHandler{uc: uc}
}
