package main

import (
	"camera-server-backend/handler"
)

func main() {
	s := handler.NewServer()
	s.Run()
}
