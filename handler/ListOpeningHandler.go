package handler

import "github.com/gin-gonic/gin"

func ListOpeningHandler(etx *gin.Context) {
	etx.JSON(200, gin.H{
		"message": "Get opening list",
	})

}