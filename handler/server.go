package handler

import "github.com/gin-gonic/gin"

// NewServer ...
func NewServer() *gin.Engine {
	e := gin.Default()

	e.GET("/ping", ping())
	v1 := e.Group("/v1")

	v1.POST(PostEventsRoute, ping())
	v1.GET(GetEventsRoute, ping())

	return e
}
