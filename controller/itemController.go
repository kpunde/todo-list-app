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
	itemService = service.NewItemService()
)

type ItemController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type itemController struct {
}

func (i *itemController) FindAll(ctx *gin.Context) {
	ctx.JSON(200, itemService.FindAll())
}

func (i *itemController) Save(ctx *gin.Context) {
	var item entity.Item
	err := ctx.ShouldBindJSON(&item)
	if err != nil {
		utility.Log(zap.ErrorLevel, nil)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		ctx.JSON(201, itemService.Save(item))
	}
}

func (i *itemController) Update(ctx *gin.Context) {
	var item entity.Item
	currentUserEmail := ctx.GetString("current-user")

	err := ctx.ShouldBindJSON(&item)
	if err != nil {
		utility.Log(zap.ErrorLevel, nil)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		id, _ := strconv.ParseInt(ctx.Param("id"), 0, 0)
		item.Id = id

		if item.Author.Email != currentUserEmail {
			ctx.JSON(http.StatusUnauthorized, nil)
			return
		}

		ctx.JSON(200, itemService.Update(item))
	}
}

func (i *itemController) Delete(ctx *gin.Context) {
	currentUserEmail := ctx.GetString("current-user")
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		utility.Log(zap.ErrorLevel, err)
	}
	item := itemService.FindById(id)
	if item.Author.Email != currentUserEmail {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}
	itemService.Delete(item)
	ctx.JSON(200, nil)
}

func NewItemController() ItemController {
	return &itemController{}
}
