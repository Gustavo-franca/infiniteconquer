package city_test

import (
	"strings"
	"testing"

	"infiniteconquer/internal/domains/city"
)

func TestRandomName(t *testing.T) {

	name := city.RandomName()

	if !strings.Contains(name, "city_") {
		t.Error("random name must contain city preffix")
	}

}

func TestCity(t *testing.T) {

	t.Run("should create a city correctly", func(t *testing.T) {
		c := city.Create("City 1")

		if c.Name != "City 1" {
			t.Errorf("city %v don't contains City 1", c)
		}

		events := c.GetEvents()

		if events[0].Type() != "city_created" {
			t.Errorf("city %v don't contains city created event", c)
		}

		if v, ok := events[0].Value().(struct {
			ID string
		}); !ok || v.ID == "" {
			t.Errorf("city %v don't contains city created event value %v", c, v)
		}

		if c.ID == "" {
			t.Errorf("city %v don't contains city id", c)
		}

	})

	t.Run("should get territory id", func(t *testing.T) {
		c := city.Create("City 1")

		c.ChangeTerritory("10")

		if c.Territory() != "10" {
			t.Error("should territory id be 10")
		}

	})

}
