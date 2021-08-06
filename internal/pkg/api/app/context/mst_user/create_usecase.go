package mst_user

import (
	"fmt"
	"go/internal/pkg/api/app/request"
	"go/internal/pkg/helper"
	"go/internal/pkg/model"
	"log"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func IsEmailMatch(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}

func HashPassword(Password string) (string, error) {
	b := []byte(Password)
	password, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		return "failed to create hashed password: %v", err
	}
	return string(password), nil
}

func (uc *userUsecase) CreateUseCase(req request.UserRequest) (interface{}, error) {
	emailisempty := helper.IsEmpty(req.Email)
	if emailisempty {
		return nil, fmt.Errorf("nama is empty")
	}

	namaisempty := helper.IsEmpty(req.Name)
	if namaisempty {
		return nil, fmt.Errorf("nama is empty")
	}

	passisempty := helper.IsEmpty(req.Password)
	if passisempty {
		return nil, fmt.Errorf("password is empty")
	}

	user, _ := uc.userRepo.FindOneBy(map[string]interface{}{
		"email": req.Email,
	})
	if req.Email == user.Email {
		return nil, fmt.Errorf("email already exist")
	}

	if !IsEmailMatch(req.Email) {
		return nil, fmt.Errorf("your email format wrong")
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		log.Printf("unable to generate hashed password: %v", err)
		return nil, fmt.Errorf("registeration failed")
	}

	nameisalphabet := helper.IsAlphabet(req.Name)
	if !nameisalphabet {
		return nil, fmt.Errorf("account name is only alphabet")
	}

	m := model.Users{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Created:  time.Now().Format("2006-01-02 15:4:5"),
		Updated:  time.Now().Format("2006-01-02 15:4:5"),
	}

	tx := uc.db.Begin()
	defer tx.Rollback()

	user, err = uc.userRepo.Create(&m, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	tx.Commit()

	return user, nil
}
