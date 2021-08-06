package session

import (
	"fmt"
	"go/internal/pkg/api/app/request"
	"log"
	"time"
)

func (uc *sessionUsecase) UpdateUseCase(req request.SessionRequest) (interface{}, error) {
	criteria := map[string]interface{}{
		"id": req.ID,
	}
	session, err := uc.sessionRepo.FindOneBy(criteria)
	log.Printf("isi session: %+v", session)
	if err != nil {
		return nil, fmt.Errorf("get data failed: %v", err)
	}
	if session.ID == 0 {
		return nil, fmt.Errorf("session data doesn't exist")
	}

	//find user
	user, err := uc.userRepo.FindOneBy(map[string]interface{}{
		"email": req.Email,
	})
	log.Printf("isi user: %+v", user)
	if err != nil {
		return nil, fmt.Errorf("login failed: %v", err)
	}
	if session.UserID != user.ID {
		return nil, fmt.Errorf("you can't update session")
	}

	session.Description = req.Description
	session.Duration = req.Duration
	session.Updated = time.Now().Format("2006-01-02 15:4:5")

	tx := uc.db.Begin()
	defer tx.Rollback()

	err = uc.sessionRepo.Update(session, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to delete data: %v", err)
	}

	tx.Commit()

	return session, nil
}
