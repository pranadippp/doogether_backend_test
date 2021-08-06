package mst_user

import (
	"go/internal/pkg/api/app/request"
	"go/internal/pkg/config"
	"go/internal/pkg/service"

	"gorm.io/gorm"
)

type UserUsecase interface {
	CreateUseCase(req request.UserRequest) (interface{}, error)
	Login(req request.UserRequest) (interface{}, error)
}

//private struct
type userUsecase struct {
	db       *gorm.DB
	userRepo service.Users
	conf     *config.AppConfig
}

func NewUserUseCase(db *gorm.DB, userRepo service.Users, conf *config.AppConfig) UserUsecase {
	return &userUsecase{
		db:       db,
		userRepo: userRepo,
		conf:     conf,
	}
}
