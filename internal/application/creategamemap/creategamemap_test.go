package creategamemap_test

import (
	"infiniteconquer/internal/application/creategamemap"
	"infiniteconquer/internal/infra/inmemory"
	"testing"
)

func TestCreateNewMap(t *testing.T) {
	t.Run("should create game map", func(t *testing.T) {

		cmd := creategamemap.Command{
			Name:          "map_1",
			Width:         10,
			Lenght:        10,
			CityNumbers:   10,
			TerritorySeed: 0.25,
		}

		repo := inmemory.NewGameMapRepository()
		game := creategamemap.NewCreateGameMap(repo)

		id, err := game.Execute(cmd)

		if err != nil {
			t.Errorf("must create map succefully, but receive error: %v", err)
		}

		if id == "" {
			t.Fatal("must have an id")
		}

	})

	t.Run("should create game map with territories and cities", func(t *testing.T) {

		cmd := creategamemap.Command{
			Name:          "map_1",
			Width:         10,
			Lenght:        10,
			CityNumbers:   10,
			TerritorySeed: 0.25,
		}

		repo := inmemory.NewGameMapRepository()
		game := creategamemap.NewCreateGameMap(repo)

		id, err := game.Execute(cmd)

		if err != nil {
			t.Errorf("must create map succefully, but receive error: %v", err)
		}

		if id == "" {
			t.Fatal("must have an id")
		}

		gmFinded, err := repo.Find(id)
		if err != nil {
			t.Errorf("must find the gm by id %s", id)
		}

		if len(gmFinded.Cities) != 10 {
			t.Errorf("must have 10 cities but receive %d", len(gmFinded.Cities))
		}

		if len(gmFinded.Territories) != 100 {
			t.Errorf("must have 100 cities but receive %d", len(gmFinded.Territories))
		}

	})

}
