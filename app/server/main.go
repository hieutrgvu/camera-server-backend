package main

import (
	"camera-server-backend/controller"
	"camera-server-backend/handler"
)

func main() {
	c := controller.NewController()
	s := handler.NewServer(c)
	s.Run()
}
