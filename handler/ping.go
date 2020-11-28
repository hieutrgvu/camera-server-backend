package handler

import (
	"github.com/gin-gonic/gin"
)

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
