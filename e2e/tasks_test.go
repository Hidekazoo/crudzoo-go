package e2e

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const BaseUrl = "http://localhost:8080"

func TestTasks(t *testing.T) {
	r, err := http.Get(BaseUrl + "/tasks")
	if err != nil {
		t.Errorf("http get err should be nil: %v", err)
	}
	defer r.Body.Close()
	var j []map[string]string
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		t.Errorf("json decode err should be nil: %v", err)
	}
	expected := []map[string]string{
		{
			"subject": "subject1",
			"body":    "body1",
			"link":    "https://example.com",
		},
		{
			"subject": "subject2",
			"body":    "body2",
			"link":    "https://example.com",
		},
		{
			"subject": "subject3",
			"body":    "body3",
			"link":    "https://example.com",
		},
	}
	assert.Equal(t, j, expected)
}
