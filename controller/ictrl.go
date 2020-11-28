package controller

import "camera-server-backend/model"

// IController ...
type IController interface {
	SaveEvents(events []model.Event) error
	GetEvents() ([]model.Event, error)
	GetImages([]string) ([]*model.ImageInfo, error)
}
