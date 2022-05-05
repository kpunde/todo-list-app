package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sampleAppDemo/service"
)

var (
	loginService = service.NewLoginService()
	jwtService   = service.NewJWTService()
)

type credentials struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func Login(ctx *gin.Context) {
	var cred credentials
	err := ctx.ShouldBind(&cred)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
	}
	isAuthenticated := loginService.Login(cred.Email, cred.Password)
	if isAuthenticated {
		ctx.JSON(http.StatusOK, gin.H{
			"token": jwtService.GenerateToken(cred.Email, true),
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, nil)
	}

}
