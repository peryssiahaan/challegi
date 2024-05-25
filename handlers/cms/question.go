package cms

import (
	"challegi/database"
	"challegi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateQuestion(ctx *gin.Context) {
	var question models.Question
	if err := ctx.BindJSON(&question); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.CreateQuestion(&question); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, question)
}

func GetQuestions(ctx *gin.Context) {
	questions, err := database.GetQuestions()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, questions)
}
