package handlers

import (
	"infiniteconquer/internal/application/creategamemap"
	"infiniteconquer/internal/application/findgamemap"
	"infiniteconquer/internal/infra/inmemory"
	"net/http"
)

func NewHandlers() http.Handler {
	s := http.NewServeMux()

	// Health Check
	PingHandler(s)

	// MAP
	gamemapRepository := inmemory.NewGameMapRepository()
	createGameMapExecutor := creategamemap.NewCreateGameMap(gamemapRepository)

	CreateMap(s, createGameMapExecutor)

	finder := findgamemap.NewMapFinder(gamemapRepository)

	FindMap(s, finder)

	return s
}
