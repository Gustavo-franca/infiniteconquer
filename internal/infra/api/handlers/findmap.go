package handlers

import (
	"encoding/json"
	"infiniteconquer/internal/application/findgamemap"
	"net/http"
)

type MapFinder interface {
	FindMap(id string) (findgamemap.Map, error)
}

func FindMap(s *http.ServeMux, f MapFinder) {

	s.HandleFunc("GET /map/{id}", func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		gameMap, err := f.FindMap(id)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		bytes, err := json.Marshal(gameMap)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(201)
		_, _ = w.Write(bytes)

	})
}
