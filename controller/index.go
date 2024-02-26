package controller

import "net/http"

// how can i register controller for each one?
func RegisterController() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", HomePage)

	return mux
}
