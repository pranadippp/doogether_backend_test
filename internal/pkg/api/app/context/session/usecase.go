package session

import (
	"go/internal/pkg/api/app/request"
	"go/internal/pkg/config"
	"go/internal/pkg/service"

	"gorm.io/gorm"
)

type SessionUsecase interface {
	CreateUseCase(req request.SessionRequest) (interface{}, error)
	UpdateUseCase(req request.SessionRequest) (interface{}, error)
	DeleteUseCase(req request.SessionRequest) (interface{}, error)
	FindDetailUseCase(req request.SessionRequest) (interface{}, error)
	FindAllUseCase(req request.SessionRequest) (interface{}, error)
}

//private struct
type sessionUsecase struct {
	db          *gorm.DB
	sessionRepo service.Session
	userRepo    service.Users
	conf        *config.AppConfig
}

func NewSessionUseCase(db *gorm.DB,
	sessionRepo service.Session,
	userRepo service.Users,
	conf *config.AppConfig) SessionUsecase {
	return &sessionUsecase{
		db:          db,
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
		conf:        conf,
	}
}
