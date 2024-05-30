package mountaintype_test

import (
	"infiniteconquer/internal/domains/territories/mountaintype"
	"testing"
)

func TestMountainTerritory(t *testing.T) {

	t.Run("should flatland type must  allow build city", func(t *testing.T) {
		buildCity := mountaintype.New().CanBuildCity()

		if buildCity {
			t.Errorf("flat land cannot build a city")
		}

	})

	t.Run("should has flatland name", func(t *testing.T) {
		if mountaintype.New().Name() != "mountain" {
			t.Error("the name must be mountain")
		}
	})

}
