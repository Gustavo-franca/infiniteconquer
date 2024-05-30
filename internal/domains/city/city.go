package city

import (
	"github.com/google/uuid"
)

type (
	Coordinates interface {
		X() int
		Y() int
		Z() int
	}

	CityID string

	City struct {
		ID        CityID
		Name      string
		territory *string
		events    []cityEvent
	}

	cityEvent struct {
		t cityTypeEvent
		v cityEventValue
	}

	cityEventValue = struct {
		ID string
	}
	cityTypeEvent string
)

const (
	created cityTypeEvent = "city_created"
)

func (c cityEvent) Type() string {
	return string(c.t)
}

func (c cityEvent) Value() any {
	return any(c.v)
}

func RandomName() string {
	return "city_" + uuid.NewString()[0:8]
}

func Create(name string) *City {
	id := uuid.NewString()
	city := NewCity(id, name, nil)
	city.addEvent(created, cityEventValue{
		ID: id,
	})
	return city
}

func NewCity(id string, name string, t *string) *City {
	return &City{
		ID:        CityID(id),
		Name:      name,
		territory: t,
	}
}

func (c *City) ChangeTerritory(t string) {
	c.territory = &t
}

func (c *City) Territory() string {
	if c.territory == nil {
		return ""
	}
	return *c.territory
}

func (c *City) addEvent(t cityTypeEvent, v cityEventValue) {
	c.events = append(c.events, cityEvent{
		t: t,
		v: v,
	})
}

func (c *City) GetEvents() []cityEvent {
	return c.events
}
