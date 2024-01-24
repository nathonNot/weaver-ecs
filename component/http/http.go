package http

import (
	"fmt"
	"net/http"
)

func GetHandler() *http.ServeMux {
	handle := http.NewServeMux()
	handle.HandleFunc("/", Hello)
	return handle
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}
