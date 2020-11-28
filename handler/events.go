package handler

import (
	"camera-server-backend/controller"
	"camera-server-backend/model"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func postEvents(ctrl controller.IController) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			QueryImage string   `json:"query_image" binding:"required"`
			Timestamp  string   `json:"timestamp" binding:"required"`
			CameraID   string   `json:"camera_id" binding:"required"`
			VectorIDs  []string `json:"vector_ids" binding:"required"`
		}

		if err := c.ShouldBind(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		log.Infof("Event from camera %s: %v\n", body.CameraID, body.VectorIDs)

		ctrl.SaveEvents([]*model.Event{
			{
				QueryImage: body.QueryImage,
				Timestamp:  body.Timestamp,
				CameraID:   body.CameraID,
				VectorIDs:  strings.Join(body.VectorIDs, ","),
			},
		})
		c.AbortWithStatus(http.StatusOK)
	}
}

// getEvents ...
func getEvents(ctrl controller.IController) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := ctrl.GetEvents()
		if err != nil {
			log.Error("getEvents: err call controller = ", err)
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, response)
	}
}

// getEvents ...
func getImages(ctrl controller.IController) gin.HandlerFunc {
	return func(c *gin.Context) {
		vectorIDs := c.QueryArray("id")
		log.Info("image id = ", vectorIDs)
		response, err := ctrl.GetImages(vectorIDs)
		if err != nil {
			log.Error("getEvents: err call controller = ", err)
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, response)
	}
}
