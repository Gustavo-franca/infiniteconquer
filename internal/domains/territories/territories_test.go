package territories_test

import (
	"infiniteconquer/internal/domains/city"
	"infiniteconquer/internal/domains/territories"
	"infiniteconquer/internal/domains/territories/mountaintype"
	"infiniteconquer/internal/domains/values/position"
	"testing"
)

func TestTerritory(t *testing.T) {

	t.Run("should create a new territory", func(t *testing.T) {
		p := position.New(10, 10)
		tr := territories.NewTerritory("id", p, mountaintype.New())

		if tr.ID != "id" {
			t.Error("id is invalid")
		}

		if tr.Type.Name() != "mountain" {
			t.Error("type is invalid")
		}

		if len(tr.Events()) > 0 {
			t.Error("events must by empty")
		}

		if tr.Position() != p {
			t.Error("position must be equal")
		}

	})

	t.Run("should create a mountain territory", func(t *testing.T) {

		p := position.New(10, 10)
		tr, err := territories.CreateTerritory(p, "mountain")
		if err != nil {
			t.Error(err)
		}

		if tr.Type.Name() != "mountain" {
			t.Error("type name must be mountain")
		}
	})

	t.Run("should create a flat land territory", func(t *testing.T) {

		p := position.New(10, 10)
		tr, err := territories.CreateTerritory(p, "flat_land")
		if err != nil {
			t.Error(err)
		}

		if tr.Type.Name() != "flat_land" {
			t.Error("type name must be flat_land")
		}
	})

	t.Run("should return error when create a invalid type territory", func(t *testing.T) {

		p := position.New(10, 10)
		_, err := territories.CreateTerritory(p, "invalid")
		if err == nil {
			t.Fatal("must receive a error")
		}

		if err.Error() != "territory type not found" {
			t.Error("the error string must be correct")
		}
	})

	t.Run("should build a city in territory when is a flatland and isn't occuped", func(t *testing.T) {
		p := position.New(10, 10)
		tr, err := territories.CreateTerritory(p, "flat_land")
		if err != nil {
			t.Fatal(err)
		}

		city := city.Create("test")

		err = tr.BuildCity(city)
		if err != nil {
			t.Fatal(err)
		}

		err = tr.BuildCity(city)
		if err == nil {
			t.Error("territory must allow only a city")
		}

		if city.Territory() != string(tr.ID) {
			t.Error("territory id must be the same vinculed in the city")
		}

	})

	t.Run("should not build a city in territory when is a mountain", func(t *testing.T) {
		p := position.New(10, 10)
		tr, err := territories.CreateTerritory(p, "mountain")
		if err != nil {
			t.Fatal(err)
		}

		city := city.Create("test")

		err = tr.BuildCity(city)
		if err == nil {
			t.Error("territory must allow only a city")
		}

		if city.Territory() != "" {
			t.Error("city must not be vinculed with mountain territory")
		}

	})

}
