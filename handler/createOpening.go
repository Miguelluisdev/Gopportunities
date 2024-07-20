package handler

import (
	"net/http"

	"github.com/Miguelluiddev/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		senError(ctx , http.StatusBadRequest , err.Error())
	}

	opening := schemas.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("error creating opening: %v", err.Error())
		senError(ctx , http.StatusInternalServerError ,"error creating opening on database")
		return
	}

	sendSucess(ctx , "create opening" ,opening)
}
