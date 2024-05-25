package main

import (
	"challegi/database"
	"challegi/handlers/cms"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.Init(); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Server successfully started."})
	})

	r.GET("/questions", cms.GetQuestions)
	r.POST("/questions", cms.CreateQuestion)

	r.Run(":8080")
}
