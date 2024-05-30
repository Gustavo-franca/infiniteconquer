package main

import (
	"infiniteconquer/internal/infra/api"
	"infiniteconquer/internal/infra/api/handlers"
	"log"
	"os"
	"strconv"
)

func main() {

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	if port == 0 || err != nil {
		port = 8080
	}
	err = api.ListenAndServe(port, handlers.NewHandlers())
	if err != nil {
		log.Fatal(err)
	}
}
