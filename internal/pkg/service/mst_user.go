package service

import (
	"go/internal/pkg/model"

	"gorm.io/gorm"
)

type Users interface {
	Create(m *model.Users, tx *gorm.DB) (*model.Users, error)
	FindOneBy(criteria map[string]interface{}) (*model.Users, error)
	Count(criteria map[string]interface{}) int
}

//private struct
type userServices struct {
	db *gorm.DB
}

//constructor
func NewUserServices(db *gorm.DB) Users {
	return &userServices{db: db}
}

func (svc *userServices) Create(m *model.Users, tx *gorm.DB) (*model.Users, error) {
	err := tx.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (svc *userServices) FindOneBy(criteria map[string]interface{}) (*model.Users, error) {
	var m model.Users
	// svc.db.LogMode(true)
	err := svc.db.Where(criteria).Find(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (svc *userServices) Count(criteria map[string]interface{}) int {
	var result int64

	err := svc.db.Where(model.Users{}).Where(criteria).Count(&result).Error
	if err != nil {
		return 0
	}

	return int(result)
}
