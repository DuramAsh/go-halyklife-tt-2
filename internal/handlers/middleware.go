package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckMethodType(ctx *gin.Context) {
	allowedMethods := []string{"GET", "POST", "PUT", "DELETE"}
	proceed := false
	for _, el := range allowedMethods {
		if ctx.Request.Method == el {
			proceed = true
			break
		}
	}
	if !proceed {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{
			"message": "this method is not allowed",
		})
		ctx.Abort()
		return
	}
	ctx.Next()
}
