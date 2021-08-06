package session

import (
	"fmt"
	"go/internal/pkg/api/app/request"
)

func (uc *sessionUsecase) FindDetailUseCase(req request.SessionRequest) (interface{}, error) {
	criteria := map[string]interface{}{
		"id": req.ID,
	}
	session, err := uc.sessionRepo.FindOneBy(criteria)
	if err != nil {
		return nil, fmt.Errorf("get data failed: %v", err)
	}
	if session.ID == 0 {
		return nil, fmt.Errorf("session data doesn't exist")
	}

	return session, nil
}
