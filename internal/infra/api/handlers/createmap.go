package handlers

import (
	"encoding/json"
	"fmt"
	"infiniteconquer/internal/application/creategamemap"
	"log"
	"net/http"
)

type Command interface {
	Execute(cmd creategamemap.Command) (string, error)
}

func CreateMap(s *http.ServeMux, c Command) {

	s.HandleFunc("POST /map", func(w http.ResponseWriter, r *http.Request) {

		var payload struct {
			Name        string  `json:"map_name"`
			SizeX       uint64  `json:"size_x"`
			SizeY       uint64  `json:"size_y"`
			CityNumbers uint64  `json:"city_numbers"`
			TerrainSeed float32 `json:"terrain_seed"`
		}

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(400)
			log.Print(err)
			return
		}

		id, err := c.Execute(creategamemap.Command{
			Name:          payload.Name,
			Width:         payload.SizeX,
			Lenght:        payload.SizeY,
			CityNumbers:   payload.CityNumbers,
			TerritorySeed: payload.TerrainSeed,
		})

		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(201)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"id": "%s"}`, id)))

	})
}
