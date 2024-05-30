package territories

import (
	"errors"

	"infiniteconquer/internal/domains/city"
	"infiniteconquer/internal/domains/territories/flatlandtype"
	"infiniteconquer/internal/domains/territories/mountaintype"
	"infiniteconquer/internal/domains/values/position"
	"math/rand"

	"github.com/google/uuid"
)

type (
	TerritoryID string
	Territory   struct {
		ID      TerritoryID
		Type    TerritoryType
		p       position.Value
		events  []event
		occuped bool
	}

	TerritoryType interface {
		Name() string
		CanBuildCity() bool
	}

	eventType string

	eventValue = struct {
		ID   string
		Type string
	}

	event struct {
		t eventType
		v eventValue
	}
)

func (t Territory) Position() position.Value {
	return t.p
}

func (v event) Type() string {
	return string(v.t)
}

func (v event) Value() any {
	return any(v.v)
}

const created eventType = "territory_created"

var ErrCannotBuild = errors.New("terriotory: cannot build resource in this territory")

func (t *Territory) addEvent(tp eventType, v eventValue) {
	t.events = append(t.events, event{t: tp, v: v})
}

func (t *Territory) Events() []event {
	return t.events
}

func (t *Territory) BuildCity(city *city.City) error {
	if t.occuped || !t.Type.CanBuildCity() {
		return ErrCannotBuild
	}

	city.ChangeTerritory(string(t.ID))
	t.occuped = true
	return nil
}

func NewTerritory(ID string, p position.Value, t TerritoryType) *Territory {
	return &Territory{ID: TerritoryID(ID), Type: t, p: p}
}

func create(p position.Value, tt TerritoryType) *Territory {
	iD := uuid.NewString()
	t := NewTerritory(iD, p, tt)
	t.addEvent(created, eventValue{ID: iD, Type: t.Type.Name()})
	return t
}

func RamdomTerritoryType(mountainPossibility float32) string {
	sortedNumber := rand.Float32()

	isMountain := sortedNumber < mountainPossibility

	if isMountain {
		return "mountain"
	}

	return "flat_land"
}

func CreateTerritory(p position.Value, territoryType string) (Territory, error) {

	switch territoryType {
	case "mountain":
		return *create(p, mountaintype.New()), nil
	case "flat_land":
		return *create(p, flatlandtype.New()), nil
	}

	return Territory{}, errors.New("territory type not found")
}
