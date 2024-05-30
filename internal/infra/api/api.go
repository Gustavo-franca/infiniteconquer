package api

import (
	"infiniteconquer/internal/infra/api/handlers"
	"net/http"
	"strconv"
)

func ListenAndServe(port int64, handlers http.Handler) error {
	return http.ListenAndServe(":"+strconv.FormatInt(port, 10), handlers)
}

func StartHttpServer(port int64, handlers http.Handler) *http.Server {
	server := &http.Server{Addr: ":" + strconv.FormatInt(port, 10), Handler: handlers}

	go func() {
		_ = server.ListenAndServe()
	}()
	return server
}

func StartServer(port int64) (*http.Server, error) {
	return StartHttpServer(port, handlers.NewHandlers()), nil
}
