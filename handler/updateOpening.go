package handler

import (
	"net/http"

	"github.com/Miguelluiddev/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validate error: %v", err.Error())
		senError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Query("id")
	if id == "" {
		senError(ctx, http.StatusBadRequest , errParamIsRequired("id" , "queryParameter").Error())
	}
	opening := schemas.Opening{}

	if err := db.First(&opening).Error; err != nil {
		senError(ctx , http.StatusNotFound , "opening not found")
	}
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.ErrorF("error updating opening: %v" , err.Error())
		senError(ctx , http.StatusInternalServerError, "error updating")
	}
	sendSucess(ctx , "update-opening" , opening)
}
