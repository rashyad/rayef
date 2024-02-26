package http

import (
	"fmt"
	"log"
	"net/http"
	c "rayef/controller"
	"time"
)

type Server struct {
	Host string
	Port string
}

func (s Server) StartServer() {

	myHandler := c.RegisterController()
	server := &http.Server{
		Addr:           s.Host + ":" + s.Port,
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Print("Server is now listening at port ", s.Port, "\n")
	log.Fatal(server.ListenAndServe())
}
