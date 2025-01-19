package handler

import "github.com/gin-gonic/gin"

func DeleteOpeningHandler(etx *gin.Context) {
	etx.JSON(200, gin.H{
		"message": "Delete opening",
	})
}