package handler

import "github.com/gin-gonic/gin"

func ShowOpeningHandler(etx *gin.Context) {
	etx.JSON(200, gin.H{
		"message": "show Opening",
	})
}