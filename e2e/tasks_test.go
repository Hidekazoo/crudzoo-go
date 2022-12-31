package e2e

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const BaseUrl = "http://localhost:8080"

type Task struct {
	Id      string `json:"id"`
	Subject string `json:"subject"`
	Link    string `json:"link"`
	Body    string `json:"body"`
}

type Re struct {
	Data []Task `json:"data"`
}

func TestTasks(t *testing.T) {
	r, err := http.Get(BaseUrl + "/tasks")
	if err != nil {
		t.Errorf("http get err should be nil: %v", err)
	}
	defer r.Body.Close()
	//var j map[string]any
	var j Re
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		t.Errorf("json decode err should be nil: %v", err)
	}
	expected := Re{
		Data: []Task{
			{
				Id:      "id1",
				Subject: "subject1",
				Body:    "body1",
				Link:    "https://example.com",
			},
			{
				Id:      "id2",
				Subject: "subject2",
				Body:    "body2",
				Link:    "https://example.com",
			},
			{
				Id:      "id3",
				Subject: "subject3",
				Body:    "body3",
				Link:    "https://example.com",
			},
		},
	}
	assert.Equal(t, j, expected)
}
