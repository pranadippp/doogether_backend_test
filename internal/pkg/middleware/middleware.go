package middleware

import (
	"go/internal/pkg/helper"
	"go/internal/pkg/response"
	"log"
	"net/http"

	"go/internal/pkg/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func InitCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conf := config.NewAppConfig()

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", conf.GetString("cors.allow_origin"))
		ctx.Writer.Header().Set("Access-Control-Allow-Header", conf.GetString("cors.allow_header"))
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", conf.GetString("cors.allow_methods"))

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusOK)
		}

		ctx.Next()
	}
}

func ValidationToken(jwtSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := helper.ExtractTokenMetadata(jwtSecret, ctx)
		if err != nil {
			response.NewErrorResponse(ctx, http.StatusUnauthorized, err, "failed to validate token")
			return
		}

		log.Printf("isi id: %v", claims.(jwt.MapClaims)["id"])
		log.Printf("isi name: %v", claims.(jwt.MapClaims)["name"])

		ctx.Set("userId", claims.(jwt.MapClaims)["id"])
		ctx.Set("userName", claims.(jwt.MapClaims)["name"])
		ctx.Set("userEmail", claims.(jwt.MapClaims)["mail"])

		ctx.Next()
	}
}
