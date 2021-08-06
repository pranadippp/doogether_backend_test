package session

import (
	"fmt"
	"go/internal/pkg/api/app/request"
	"go/internal/pkg/response"
	"net/http"
)

func (uc *sessionUsecase) FindAllUseCase(req request.SessionRequest) (interface{}, error) {

	session, err := uc.sessionRepo.FindAll(req.Page, req.Size, req.Order)
	if err != nil {
		return nil, fmt.Errorf("get data failed: %v", err)
	}
	if len(session) == 0 {
		return nil, fmt.Errorf("session data not found")
	}

	resp := response.PagedResponse{
		Ok:     true,
		Status: "success",
		Code:   http.StatusOK,
		Data:   session,
		Total:  uc.sessionRepo.Count(nil),
	}

	return resp, nil
}
