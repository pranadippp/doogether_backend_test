package helper

import (
	"fmt"
	"go/internal/pkg/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(jwtSecret string, user *model.Users) (map[string]interface{}, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "doogether-api",                                 // issuer token
		"id":   user.ID,                                         //token user
		"name": user.Name,                                       //token name user
		"mail": user.Email,                                      //token email user
		"iat":  time.Now().Unix(),                               //token time
		"exp":  time.Now().Add(time.Hour * 24 * 30 * 12).Unix(), //expire token time
		"type": "access_token",
	})
	t, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"mail": user.Email,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"type": "refresh_token",
	})

	rt, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return map[string]interface{}{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}

func GetHeaderBearerToken(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func VerifyToken(jwtSecret string, ctx *gin.Context) (*jwt.Token, error) {

	tokenString := GetHeaderBearerToken(ctx)
	if tokenString == "" {
		return nil, fmt.Errorf("An authorization header is required")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		//Make sure that the token method confirm to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected Signing Method: %v", token.Header["alg"])

		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func TokenValidation(jwtSecret string, ctx *gin.Context) error {
	token, err := VerifyToken(jwtSecret, ctx)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(jwtSecret string, ctx *gin.Context) (interface{}, error) {
	token, err := VerifyToken(jwtSecret, ctx)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid auth token")
	}
}

//check if token expired before expired time
// func (c *MyClaims) VerifyIssuedAt(cmp int64, req bool) bool {
// 	return true
// }
