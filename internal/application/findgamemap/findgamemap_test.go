package findgamemap_test

import (
	"infiniteconquer/internal/application/creategamemap"
	"infiniteconquer/internal/application/findgamemap"
	"infiniteconquer/internal/infra/inmemory"
	"testing"
)

func mockGameMap(t *testing.T) (string, findgamemap.GameRepository) {
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
		t.Fatalf("must create map succefully, but receive error: %v", err)
	}

	if id == "" {
		t.Fatal("must have an id")
	}

	return id, repo

}

func TestFindGameMap(t *testing.T) {
	t.Run("should create game map with territories and cities", func(t *testing.T) {

		id, repo := mockGameMap(t)

		finder := findgamemap.NewMapFinder(repo)

		gmFinded, err := finder.FindMap(id)
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
