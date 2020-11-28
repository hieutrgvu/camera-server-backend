package handler

import (
	"camera-server-backend/controller"
	"github.com/gin-gonic/gin"
)

// NewServer ...
func NewServer(ctrl controller.IController) *gin.Engine {
	e := gin.Default()

	e.GET("/ping", ping())
	v1 := e.Group("/v1")

	v1.POST(PostEventsRoute, postEvents(ctrl))
	v1.GET(GetEventsRoute, getEvents(ctrl))
	v1.GET(GetImagesRoute, getImages(ctrl))

	return e
}
