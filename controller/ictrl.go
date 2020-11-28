package controller

import "camera-server-backend/model"

// IController ...
type IController interface {
	SaveEvents(events []*model.Event) error
}