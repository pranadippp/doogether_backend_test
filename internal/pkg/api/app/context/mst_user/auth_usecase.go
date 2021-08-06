package mst_user

import (
	"fmt"
	"go/internal/pkg/api/app/request"
	"go/internal/pkg/helper"

	"golang.org/x/crypto/bcrypt"
)

func IsPasswordMatch(hashed, plain string) bool {
	byteHash := []byte(hashed)
	bytePlain := []byte(plain)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false
	}
	return true
}

func (uc *userUsecase) Login(req request.UserRequest) (interface{}, error) {
	jwtSecret := uc.conf.GetString("app.signature")

	user, err := uc.userRepo.FindOneBy(map[string]interface{}{
		"email": req.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("login failed: %v", err)
	}

	if req.Email != user.Email {
		return nil, fmt.Errorf("user doesn't exist")
	}

	if !IsPasswordMatch(user.Password, req.Password) {
		return nil, fmt.Errorf("email or password doesn't match")
	}

	return helper.GenerateToken(jwtSecret, user)
}
