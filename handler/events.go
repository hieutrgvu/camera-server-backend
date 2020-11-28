package handler

import (
	"camera-server-backend/controller"
	"camera-server-backend/model"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func postEvents(ctrl controller.IController) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body model.Event
		if err := c.ShouldBind(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		log.Infof("Event from camera %s: %v\n", body.CameraID, body.VectorIds)

		ctrl.SaveEvents([]*model.Event{&body})
		c.AbortWithStatus(http.StatusOK)
	}
}
