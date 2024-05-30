package flatlandtype_test

import (
	"testing"

	"infiniteconquer/internal/domains/territories/flatlandtype"
)

func TestFlatLandTerritory(t *testing.T) {

	t.Run("should flatland type must  allow build city", func(t *testing.T) {
		buildCity := flatlandtype.New().CanBuildCity()

		if !buildCity {
			t.Errorf("flat land can build a city")
		}

	})

	t.Run("should has flatland name", func(t *testing.T) {
		if flatlandtype.New().Name() != "flat_land" {
			t.Error("the name must be flat_land")
		}
	})

}
