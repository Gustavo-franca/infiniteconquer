package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type Map struct {
	ID          string      `json:"id"`
	MapName     string      `json:"map_name"`
	SizeX       int64       `json:"size_x"`
	SizeY       int64       `json:"size_y"`
	Cities      []City      `json:"cities"`
	Territories []Territory `json:"territories"`
}

type City struct {
	CityName  string `json:"city Name"`
	PositionX int64  `json:"position_X"`
	PositionY int64  `json:"position_Y"`
}

type Territory struct {
	Type      string `json:"type"`
	PositionX int64  `json:"position_X"`
	PositionY int64  `json:"position_Y"`
}

func TestCreateAndFindMap(t *testing.T) {

	baseUrl := runAPI(t)

	id := CreateSeedMap(t, baseUrl)

	request, err := http.NewRequest("GET", fmt.Sprint(baseUrl, fmt.Sprint("/map/", id)), nil)
	if err != nil {
		t.Error(err)
	}

	client := http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected created status instead received %s - %d", resp.Status, resp.StatusCode)
	}

	var p Map
	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		t.Errorf("expected a valid payload %v", err)
	}

	if p.MapName != "name" {
		t.Error("map name must be name")
	}

	if p.SizeX != 10 {
		t.Error("size x must 10")
	}

	if p.SizeY != 10 {
		t.Error("size y must 10")
	}

	if len(p.Territories) != 100 {
		t.Errorf("must contain 100 territories - contains %d", len(p.Territories))

	}

	if p.Territories[0].PositionX != 0 && p.Territories[0].PositionY != 0 {
		t.Error("the first territory must have position 0")
	}

	lastTerroryIndex := len(p.Territories) - 1
	if p.Territories[lastTerroryIndex].PositionX != 9 && p.Territories[lastTerroryIndex].PositionY != 9 {
		t.Error("the last territory must have position 9")
	}

	if p.ID != id {
		t.Error("map must constain the same id")
	}

	if len(p.Cities) != 10 {
		t.Error("must contain 10 cities")
	}

}

func CreateSeedMap(t *testing.T, baseUrl string) string {
	t.Helper()

	body := bytes.NewBufferString(`
	{
		"map_name": "name",
		"size_x" : 10,
		"size_y" : 10,
		"city_numbers" : 10,
		"terrain_seed" : 0.25
	}
	`)

	request, err := http.NewRequest("POST", fmt.Sprint(baseUrl, "/map"), body)
	if err != nil {
		t.Error(err)
	}

	client := http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected created status instead received %s - %d", resp.Status, resp.StatusCode)
	}

	var p struct {
		ID string `json:"id"`
	}

	err = json.NewDecoder(resp.Body).Decode(&p)
	if err != nil {
		t.Errorf("expected a valid payload %v", err)

	}

	return p.ID
}
