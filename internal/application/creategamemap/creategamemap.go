package creategamemap

import (
	"infiniteconquer/internal/domains/gamemap"
)

type (
	GameRepository interface {
		Save(gameMap *gamemap.GameMap) error
	}

	CreateGameMap struct {
		gameRepository GameRepository
	}

	Command struct {
		Name          string
		Width         uint64
		Lenght        uint64
		CityNumbers   uint64
		TerritorySeed float32
	}
)

func NewCreateGameMap(rep GameRepository) CreateGameMap {
	return CreateGameMap{
		gameRepository: rep,
	}
}

func (c CreateGameMap) Execute(cmd Command) (string, error) {

	gm := gamemap.StartNewGameMap(cmd.Name)

	err := gm.GenerateTerritories(cmd.Width, cmd.Lenght, cmd.TerritorySeed)
	if err != nil {
		return "", err
	}

	err = gm.BuildCities(cmd.CityNumbers)

	if err != nil {
		return "", err
	}

	return string(gm.ID), c.gameRepository.Save(gm)

}
