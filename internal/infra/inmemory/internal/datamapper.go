package internal

import (
	"infiniteconquer/internal/domains/city"
	"infiniteconquer/internal/domains/gamemap"
	"infiniteconquer/internal/domains/territories"
	"infiniteconquer/internal/domains/territories/flatlandtype"
	"infiniteconquer/internal/domains/territories/mountaintype"
	"infiniteconquer/internal/domains/values/position"
)

type (
	Map struct {
		ID          string      `json:"id"`
		MapName     string      `json:"map_name"`
		SizeX       int64       `json:"size_x"`
		SizeY       int64       `json:"size_y"`
		Cities      Cities      `json:"cities"`
		Territories Territories `json:"territories"`
	}

	City struct {
		ID          string
		CityName    string `json:"city Name"`
		PositionX   int64  `json:"position_X"`
		PositionY   int64  `json:"position_Y"`
		TerritoryID string
	}

	Cities []City

	Territories []Territory

	Territory struct {
		ID        string
		Type      Type  `json:"type"`
		PositionX int64 `json:"position_X"`
		PositionY int64 `json:"position_Y"`
	}

	Type string
)

func ToMap(gm *gamemap.GameMap) Map {

	territories := gm.Territories()

	terTransformed := make([]Territory, 0, len(territories))
	for _, t := range territories {
		x, y := t.Position().Get()
		terTransformed = append(terTransformed, Territory{
			ID:        string(t.ID),
			PositionX: int64(x),
			PositionY: int64(y),
			Type:      Type(t.Type.Name()),
		})

	}

	cities := gm.Cities()

	citiesTransformed := make([]City, 0, len(cities))
	for _, t := range cities {
		x, y := gm.FindTerritory(t.Territory()).Position().Get()
		citiesTransformed = append(citiesTransformed, City{
			ID:          string(t.ID),
			PositionX:   int64(x),
			PositionY:   int64(y),
			CityName:    t.Name,
			TerritoryID: t.Territory(),
		})

	}

	return Map{
		ID:          string(gm.ID),
		Territories: terTransformed,
		Cities:      citiesTransformed,
		MapName:     gm.Name,
		SizeX:       gm.SizeX,
		SizeY:       gm.SizeY,
	}
}

func (t Cities) ToEntity() []city.City {
	cities := make([]city.City, len(t))
	for _, ter := range t {
		cities = append(cities, ter.ToEntity())
	}

	return cities
}

func (t City) ToEntity() city.City {
	return *city.NewCity(
		t.ID,
		t.CityName,
		&t.TerritoryID,
	)
}

func (t Type) ToEntity() territories.TerritoryType {
	switch t {
	case "mountain":
		return mountaintype.New()
	case "flat_land":
		return flatlandtype.New()
	}
	return nil
}

func (t Territory) ToEntity() territories.Territory {
	return *territories.NewTerritory(
		t.ID,
		position.New(t.PositionX, t.PositionY),
		t.Type.ToEntity(),
	)
}

func (t Territories) ToEntity() []territories.Territory {
	ters := make([]territories.Territory, len(t))
	for _, ter := range t {
		ters = append(ters, ter.ToEntity())
	}

	return ters

}
func (m Map) ToEntity() *gamemap.GameMap {

	terr := m.Territories.ToEntity()
	cities := m.Cities.ToEntity()

	return gamemap.NewGameMap(
		m.ID,
		m.MapName,
		m.SizeX,
		m.SizeY,
		terr,
		cities,
	)
}
