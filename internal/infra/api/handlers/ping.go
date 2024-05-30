package handlers

import (
	"net/http"
)

func PingHandler(server *http.ServeMux) {
	server.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
}
