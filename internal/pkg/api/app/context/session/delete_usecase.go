package session

import (
	"fmt"
	"go/internal/pkg/api/app/request"
)

func (uc *sessionUsecase) DeleteUseCase(req request.SessionRequest) (interface{}, error) {
	//find session
	criteria := map[string]interface{}{
		"id": req.ID,
	}
	session, err := uc.sessionRepo.FindOneBy(criteria)
	if err != nil {
		return nil, fmt.Errorf("get data failed: %v", err)
	}
	if session.ID == 0 {
		return nil, fmt.Errorf("session doesn't exist")
	}

	//find user
	user, err := uc.userRepo.FindOneBy(map[string]interface{}{
		"email": req.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("login failed: %v", err)
	}
	if session.UserID != user.ID {
		return nil, fmt.Errorf("you can't delete session")
	}

	tx := uc.db.Begin()
	defer tx.Rollback()

	err = uc.sessionRepo.Delete(session, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to delete data: %v", err)
	}

	tx.Commit()

	return session, nil
}
