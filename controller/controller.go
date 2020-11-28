package controller

import (
	"camera-server-backend/model"
	"fmt"
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// Controller ...
type Controller struct {
}

// NewController ...
func NewController() *Controller {
	return &Controller{}
}

const (
	eventImageDir = "./images"
)

var (
	storageEvents = make([]model.EventRecord, 0, 1000)
)

func init() {
	os.MkdirAll(eventImageDir, os.ModePerm)
}

// SaveEvents ...
func (c *Controller) SaveEvents(events []*model.Event) error {

	for _, e := range events {
		file, _ := e.QueryImage.Open()
		defer file.Close()
		filePath := fmt.Sprintf("./%s/%s_%s_%s",
			eventImageDir,
			e.CameraID,
			time.Now().Format("20060102.150405.000000"),
			e.QueryImage.Filename)
		dst, _ := os.Create(filePath)
		defer dst.Close()
		io.Copy(dst, file)

		storageEvents = append(storageEvents, model.EventRecord{
			QueryImagePath: filePath,
			Timestamp:      e.Timestamp,
			CameraID:       e.CameraID,
			VectorIds:      e.VectorIds,
		})
	}

	log.Info("Save Events Done")
	return nil
}
