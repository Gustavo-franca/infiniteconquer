package inmemory

import (
	"encoding/json"
	"errors"

	"infiniteconquer/internal/application/findgamemap"
	"infiniteconquer/internal/domains/gamemap"
	"infiniteconquer/internal/infra/inmemory/internal"
)

type (
	Event interface {
		Type() string
		Value() any
	}
)

type InMemoryGameMap struct {
	g map[string][]byte
}

func NewGameMapRepository() *InMemoryGameMap {
	return &InMemoryGameMap{
		g: make(map[string][]byte),
	}
}

func (i *InMemoryGameMap) Save(g *gamemap.GameMap) error {
	bytes, err := json.Marshal(internal.ToMap(g))

	i.g[string(g.ID)] = bytes
	return err

}

func (i InMemoryGameMap) Find(id string) (findgamemap.Map, error) {

	var gameMap internal.Map

	bytes, ok := i.g[id]
	if !ok {
		return findgamemap.Map{}, errors.New("not_found gamemap")
	}

	err := json.Unmarshal(bytes, &gameMap)

	if err != nil || len(i.g) == 0 {
		return findgamemap.Map{}, errors.New("not_found gamemap")
	}

	terTransformed := make([]findgamemap.Territory, 0, len(gameMap.Territories))
	for _, t := range gameMap.Territories {
		x, y := t.PositionX, t.PositionY
		terTransformed = append(terTransformed, findgamemap.Territory{
			PositionX: int64(x),
			PositionY: int64(y),
			Type:      string(t.Type),
		})

	}

	citiesTransformed := make([]findgamemap.City, 0, len(gameMap.Cities))
	for _, t := range gameMap.Cities {
		x, y := t.PositionX, t.PositionY
		citiesTransformed = append(citiesTransformed, findgamemap.City{
			PositionX: int64(x),
			PositionY: int64(y),
			CityName:  t.CityName,
		})

	}

	return findgamemap.Map{
		ID:          string(gameMap.ID),
		Territories: terTransformed,
		Cities:      citiesTransformed,
		MapName:     gameMap.MapName,
		SizeX:       gameMap.SizeX,
		SizeY:       gameMap.SizeY,
	}, nil
}
