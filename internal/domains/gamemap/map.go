package gamemap

import (
	"errors"

	"github.com/google/uuid"

	"infiniteconquer/internal/domains/city"
	"infiniteconquer/internal/domains/territories"
	"infiniteconquer/internal/domains/values/position"
)

type (
	GameMapEvent struct {
		t eventType
		v eventValue
	}

	eventValue struct {
		ID string
	}

	eventType string

	GameMapID string

	GameMap struct {
		Name        string
		SizeX       int64
		SizeY       int64
		ID          GameMapID
		territories []territories.Territory
		cities      []city.City
		events      []GameMapEvent
	}
)

const (
	created eventType = "game_map_created"
)

var (
	ErrLenght = errors.New("lenght should be grather than 0 and less than 1000")
	ErrWidth  = errors.New("width should be greather than 0 and less than 1000")
	ErrSeed   = errors.New("seed should be betwen 0 and 1")
)

func (g GameMapEvent) Type() string {
	return string(g.t)
}

func (g GameMapEvent) Value() any {
	return g.v
}

func NewGameMap(
	id string,
	name string,
	sizeX int64,
	sizeY int64,
	territories []territories.Territory,
	cities []city.City) *GameMap {
	return &GameMap{
		ID:          GameMapID(id),
		Name:        name,
		territories: territories,
		cities:      cities,
		SizeX:       sizeX,
		SizeY:       sizeY,
	}
}

func StartNewGameMap(name string) *GameMap {
	iD := GameMapID(uuid.NewString())
	return &GameMap{
		Name:        name,
		ID:          iD,
		territories: nil,
		events: []GameMapEvent{{
			t: created,
			v: eventValue{
				ID: string(iD),
			},
		}},
	}
}

func (gm GameMap) Territories() []territories.Territory {
	return gm.territories
}

func (gm GameMap) Cities() []city.City {
	return gm.cities
}

func (gm GameMap) FindTerritory(territoryID string) territories.Territory {
	for _, t := range gm.territories {
		if string(t.ID) == territoryID {
			return t
		}

	}
	return territories.Territory{}
}

func (gm GameMap) Events() []GameMapEvent {
	return gm.events
}

func (gm *GameMap) buildCity(name string) (city.City, error) {

	city := city.Create(name)

	for _, t := range gm.territories {
		if err := t.BuildCity(city); err == nil {
			return *city, nil
		}
	}

	return *city, errors.New("all territories ocupped")

}

func (gm *GameMap) BuildCities(quantity uint64) error {
	gm.cities = make([]city.City, 0, quantity)

	for range quantity {
		city, err := gm.buildCity(city.RandomName())
		if err != nil {
			return err
		}
		gm.cities = append(gm.cities, city)
	}
	return nil
}

func (gm *GameMap) GenerateTerritories(width, lenght uint64, seed float32) error {
	if width == 0 || width > 1000 {
		return ErrWidth
	}

	if lenght == 0 || lenght > 1000 {
		return ErrLenght
	}

	if seed < 0 || seed > 1 {
		return ErrSeed
	}

	gm.SizeX = int64(width)
	gm.SizeY = int64(lenght)

	for x := 0; x < int(width); x++ {
		for y := 0; y < int(lenght); y++ {
			p := position.New(int64(x), int64(y))
			tr, _ := territories.CreateTerritory(p, territories.RamdomTerritoryType(seed)) // Random Type always generate a valid type
			gm.territories = append(gm.territories, tr)
		}
	}

	return nil
}
