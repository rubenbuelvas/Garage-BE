package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	}
}
