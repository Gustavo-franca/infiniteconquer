package api_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {

	baseUrl := runAPI(t)

	request, err := http.NewRequest("GET", fmt.Sprint(baseUrl, "/ping"), nil)
	if err != nil {
		t.Error(err)
	}

	client := http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status ok instead received %s - %d", resp.Status, resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if string(bytes) != "pong" {
		t.Errorf("expected body must be pong instead received %s", string(bytes))
	}

}
