package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"infiniteconquer/internal/infra/api"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func runAPI(t *testing.T) string {

	s, err := api.StartServer(8080)

	if err != nil {
		t.Error(err)
	}

	t.Cleanup(func() {
		s.Close()
	})

	return "http://localhost:8080"
}

func TestCreateNewMap(t *testing.T) {

	baseUrl := runAPI(t)

	body := bytes.NewBufferString(`
	{
		"map_name": "name",
		"size_x" : 100,
		"size_y" : 100,
		"city_numbers" : 10,
		"terrain_seed" : 0.25,
		"neutral_cities" : 10
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

	_, err = uuid.Parse(p.ID)
	if p.ID == "" || err != nil {
		t.Errorf("expected a valid uuid %v", err)
	}

}
