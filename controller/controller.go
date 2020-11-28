package controller

import (
	"camera-server-backend/model"

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
