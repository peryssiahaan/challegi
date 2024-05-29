package cms

import (
	"challegi/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAnswer(ctx *gin.Context) {
	var answers database.Answer
	if err := ctx.BindJSON(&answers); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.CreateAnswer(&answers); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, answers)
}

func GetAnswers(ctx *gin.Context) {
	answers, err := database.GetAnswers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, answers)
}
