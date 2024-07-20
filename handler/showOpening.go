package handler

import (
	"net/http"

	"github.com/Miguelluiddev/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		senError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		senError(ctx , http.StatusNotFound , "opening not found")
		return
	}

	sendSucess(ctx , "show-opening" , opening)

}
