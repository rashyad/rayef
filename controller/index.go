package controller

import (
	"net/http"
)

// how can i register controller for each one?
func RegisterController() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", HomePage)

	fileServer := http.FileServer(http.Dir("static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
