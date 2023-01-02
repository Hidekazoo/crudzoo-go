package handler

import (
	"crudzoo-go/infra"
	"crudzoo-go/repository"
	"crudzoo-go/usecase"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type Task struct {
	Id      uuid.UUID `json:"id"`
	Subject string    `json:"subject"`
	Link    string    `json:"link"`
	Body    string    `json:"body"`
}

type Config struct {
	DBSetting string `envconfig:"DB_SETTING" required:"true"`
}

func Tasks(w http.ResponseWriter, r *http.Request) {
	dbImpl := infra.DBImple{}
	tasksPort := repository.TasksRepository{DB: &dbImpl}
	tasks, _ := usecase.FindTasks(&tasksPort)

	var tasksRes []*Task
	for _, v := range tasks {
		id, _ := uuid.Parse(v.Id)
		tasksRes = append(tasksRes, &Task{
			Id:      id,
			Subject: v.Content.Subject,
			Link:    v.Content.Link,
			Body:    v.Content.Body,
		})
	}
	res := map[string][]*Task{
		"data": tasksRes,
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
