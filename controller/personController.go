package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sampleAppDemo/entity"
	"sampleAppDemo/service"
	"sampleAppDemo/utility"
	"strconv"
)

var (
	personService = service.NewPersonService()
)

type PersonController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindItems(ctx *gin.Context)
}

type personController struct {
}

func (i *personController) FindItems(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 0)
	ctx.JSON(200, itemService.FindByUserId(id))
}

func (i *personController) FindAll(ctx *gin.Context) {
	ctx.JSON(200, personService.FindAll())
}

func (i *personController) Save(ctx *gin.Context) {
	var person entity.User
	err := ctx.ShouldBindJSON(&person)
	if err != nil {
		utility.Log(zap.ErrorLevel, nil)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		person.PasswordHash = utility.HashAndSalt(person.PasswordHash)
		ctx.JSON(201, personService.Save(person))
	}
}

func (i *personController) Update(ctx *gin.Context) {
	var person entity.User
	err := ctx.ShouldBindJSON(&person)
	if err != nil {
		utility.Log(zap.ErrorLevel, nil)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		id, _ := strconv.ParseInt(ctx.Param("id"), 0, 0)
		person.Id = id
		person.PasswordHash = utility.HashAndSalt(person.PasswordHash)
		ctx.JSON(200, personService.Update(person))
	}
}

func (i *personController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		utility.Log(zap.ErrorLevel, err)
	}
	person := personService.FindById(id)
	personService.Delete(person)
	ctx.JSON(200, nil)
}

func NewPersonController() PersonController {
	return &personController{}
}
