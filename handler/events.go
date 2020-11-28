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
		var payload model.Event
		if err := c.ShouldBind(&payload); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		log.Infof("Event recevied: %+v\n", payload)

		ctrl.SaveEvents([]*model.Event{&payload})
		c.AbortWithStatus(http.StatusOK)
	}
}
