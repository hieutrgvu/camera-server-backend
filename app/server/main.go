package main

import (
	"camera-server-backend/controller"
	"camera-server-backend/handler"
	"camera-server-backend/pkgs/store"
)

func main() {
	s, err := store.NewStore()
	if err != nil {
		panic(err)
	}

	c := controller.NewController(s)
	h := handler.NewHandler(c)
	h.Run()
}
