package main

import (
	server "rayef/http"
)

func main() {

	var s = server.Server{
		Host: "localhost",
		Port: "8080",
	}

	s.InitServer()

}
