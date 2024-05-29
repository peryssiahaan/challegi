package cms

import (
	"challegi/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGroup(ctx *gin.Context) {
	var group database.Group
	if err := ctx.BindJSON(&group); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.CreateGroup(&group); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, group)
}

func GetGroups(ctx *gin.Context) {
	groups, err := database.GetGroups()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, groups)
}
