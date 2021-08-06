package session

import (
	"fmt"
	"go/internal/pkg/api/app/request"
	"go/internal/pkg/model"
	"time"
)

func (uc *sessionUsecase) CreateUseCase(req request.SessionRequest) (interface{}, error) {

	user, err := uc.userRepo.FindOneBy(map[string]interface{}{
		"email": req.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("login failed: %v", err)
	}

	m := model.Session{
		UserID:      user.ID,
		Name:        req.Name,
		Description: req.Description,
		Start:       time.Now().Format("2006-01-02 15:4:5"),
		Duration:    req.Duration,
		Created:     time.Now().Format("2006-01-02 15:4:5"),
		Updated:     time.Now().Format("2006-01-02 15:4:5"),
	}

	tx := uc.db.Begin()
	defer tx.Rollback()

	session, err := uc.sessionRepo.Create(&m, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	tx.Commit()

	return session, nil
}
