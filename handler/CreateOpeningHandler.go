package handler

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(etx *gin.Context) {
	etx.JSON(200, gin.H{
		"message": "Create opening",
	})
}