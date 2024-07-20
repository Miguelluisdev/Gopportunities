package handler

import (
	"net/http"

	"github.com/Miguelluiddev/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		senError(ctx, http.StatusInternalServerError , "error listing openings")
		return
	}
	sendSucess(ctx , "list openings" , openings)
}
