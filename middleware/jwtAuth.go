package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"net/http"
	"sampleAppDemo/repository"
	"sampleAppDemo/service"
	"sampleAppDemo/utility"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			errorWithAbort(context, "Header not found")
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.NewJWTService().ValidateToken(tokenString)
		if token.Valid {
			repo := repository.NewPersonRepository()
			claims := token.Claims.(jwt.MapClaims)
			email := fmt.Sprintf("%v", claims["email"])
			result, err := repo.FindByEmail(email)
			if err != nil {
				errorWithAbort(context, err)
			}

			context.Set("current-user", result.Email)

		} else {
			errorWithAbort(context, err)
		}
	}
}

func errorWithAbort(context *gin.Context, err interface{}) {
	utility.Log(zap.ErrorLevel, err)
	context.AbortWithStatus(http.StatusUnauthorized)
}
