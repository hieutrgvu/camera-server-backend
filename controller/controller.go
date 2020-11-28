package controller

import (
	"camera-server-backend/model"
	"fmt"

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
	storageEvents = make([]*model.Event, 0, 1000)
)

// SaveEvents ...
func (c *Controller) SaveEvents(events []*model.Event) error {
	storageEvents = append(storageEvents, events...)
	log.Info("Save Events Done")
	return nil
}

// GetEvents ...
func (c *Controller) GetEvents() ([]*model.Event, error) {
	response := storageEvents
	log.Info("Load Events Done")
	return response, nil
}

// GetEvents ...
func (c *Controller) GetImages(vectorID string) (string, error) {
	response := "Image Default"

	if len(vectorID) == 0 {
		return "", fmt.Errorf("Not found")
	}

	log.Info("Load Images Done")
	return response, nil
}
