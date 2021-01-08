package main

import (
	"camera-server-backend/controller"
	"camera-server-backend/handler"
	"camera-server-backend/helper"
	"camera-server-backend/pkgs/store"
	"fmt"
)

const serverPort = 8080

func main() {
	helper.InitDataImagePath()
	s, err := store.NewStore()
	if err != nil {
		panic(err)
	}

	c := controller.NewController(s)
	h := handler.NewHandler(c)
	h.Run(fmt.Sprintf(":%d", serverPort))
}
