package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	Id      string `json:"id"`
	Subject string `json:"subject"`
	Link    string `json:"link"`
	Body    string `json:"body"`
}

func Tasks(w http.ResponseWriter, r *http.Request) {
	tasks := []Task{
		{Id: "id1", Subject: "subject1", Link: "https://example.com", Body: "body1"},
		{Id: "id2", Subject: "subject2", Link: "https://example.com", Body: "body2"},
		{Id: "id3", Subject: "subject3", Link: "https://example.com", Body: "body3"},
	}
	res := map[string][]Task{
		"data": tasks,
	}
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, fmt.Sprintf(`{"status": "%s"}`, err), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, `{"status": "permits only GET"}`, http.StatusMethodNotAllowed)
	}
}
