package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	Subject string `json:"id"`
	Link    string `json:"link"`
	Body    string `json:"body"`
}

func Tasks(w http.ResponseWriter, r *http.Request) {
	tasks := []Task{
		{Subject: "subject1", Link: "https://example.com", Body: "body1"},
		{Subject: "subject2", Link: "https://example.com", Body: "body2"},
		{Subject: "subject3", Link: "https://example.com", Body: "body3"},
	}
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(tasks); err != nil {
			http.Error(w, fmt.Sprintf(`{"status": "%s"}`, err), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, `{"status": "permits only GET"}`, http.StatusMethodNotAllowed)
	}
}
