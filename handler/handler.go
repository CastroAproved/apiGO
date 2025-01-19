package handler

import (
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(etx *gin.Context){
	etx.JSON(200, gin.H{
		"message": "show Opening",
	})
}

func CreateOpeningHandler(etx *gin.Context){
	etx.JSON(200, gin.H{
		"message": "Create opening",
	})
}

	func DeleteOpeningHandler(etx *gin.Context){
		etx.JSON(200, gin.H{
			"message": "Delete opening",
		})
}


func UpdateOpeningHandler(etx *gin.Context){
	etx.JSON(200, gin.H{
		"message": "Update opening",
	})
}

func ListOpeningHandler(etx *gin.Context){
	etx.JSON(200, gin.H{
		"message": "Get opening list",
	})


}









