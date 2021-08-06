package service

import (
	"go/internal/pkg/model"

	"gorm.io/gorm"
)

type Session interface {
	Create(m *model.Session, tx *gorm.DB) (*model.Session, error)
	Update(m *model.Session, tx *gorm.DB) error
	Delete(m *model.Session, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.Session, error)
	FindAll(page, size int, order string) ([]*model.Session, error)
	Count(criteria map[string]interface{}) int
}

//private struct
type sessionServices struct {
	db *gorm.DB
}

//constructor
func NewSessionServices(db *gorm.DB) Session {
	return &sessionServices{db: db}
}

func (svc *sessionServices) Create(m *model.Session, tx *gorm.DB) (*model.Session, error) {
	err := tx.Create(&m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (svc *sessionServices) Update(m *model.Session, tx *gorm.DB) error {
	err := tx.Model(&m).Where("ID=?", m.ID).Updates(&m).Error
	if err != nil {
		return err
	}

	return nil
}

func (svc *sessionServices) Delete(m *model.Session, tx *gorm.DB) error {
	err := tx.Model(&m).Where("ID=?", m.ID).Delete(&m).Error
	if err != nil {
		return err
	}

	return nil
}

func (svc *sessionServices) FindBy(criteria map[string]interface{}) ([]*model.Session, error) {
	var data []*model.Session

	err := svc.db.Where(criteria).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *sessionServices) FindOneBy(criteria map[string]interface{}) (*model.Session, error) {
	var m model.Session

	err := svc.db.Where(criteria).Find(&m).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (svc *sessionServices) FindAll(page, size int, order string) ([]*model.Session, error) {
	var data []*model.Session

	if page == 0 || size == 0 {
		page, size = 1, 10
	}

	offset := (page - 1) * size

	if order == "" {
		order = "created desc"
	}

	err := svc.db.Limit(size).Offset(offset).Order(order).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (svc *sessionServices) Count(criteria map[string]interface{}) int {
	var result int64

	err := svc.db.Model(model.Session{}).Count(&result).Error
	if err != nil {
		return 0
	}
	if len(criteria) >= 1 {
		err := svc.db.Model(model.Session{}).Where(criteria).Count(&result).Error
		if err != nil {
			return 0
		}
	}

	return int(result)
}
