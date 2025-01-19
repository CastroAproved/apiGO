package handler

import "github.com/gin-gonic/gin"

func UpdateOpeningHandler(etx *gin.Context) {
	etx.JSON(200, gin.H{
		"message": "Update opening",
	})
}