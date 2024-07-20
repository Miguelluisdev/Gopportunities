package handler

import (
	"fmt"
	"net/http"

	"github.com/Miguelluiddev/Gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		senError(ctx, http.StatusBadRequest , errParamIsRequired("id" , "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}

	 if err := db.First(&opening , id).Error; err != nil {
		senError(ctx, http.StatusNotFound , fmt.Sprintf("openning with id: %s not found" , id))
		return
	 }

	 if err := db.Delete(&opening).Error; err != nil {
		senError(ctx , http.StatusInternalServerError , fmt.Sprintf("error deleting opening with id: %s" , id))
		return
	 }
	 sendSucess(ctx , "delete-opening" , opening)
}
