package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/kelseyhightower/envconfig"
	"log"
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
	//tasks := []Task{
	//	{Id: "id1", Subject: "subject1", Link: "https://example.com", Body: "body1"},
	//	{Id: "id2", Subject: "subject2", Link: "https://example.com", Body: "body2"},
	//	{Id: "id3", Subject: "subject3", Link: "https://example.com", Body: "body3"},
	//}
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatalf("env error %v", err)
	}

	db, err := sql.Open("pgx", c.DBSetting)
	if nil != err {
		log.Fatal("infra connected error")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("infra ping error")
	}
	type TaskDB struct {
		TaskID  uuid.UUID
		Subject string
		Link    string
		Body    string
	}

	rows, err := db.Query(`SELECT task_id, subject, link, body FROM tasks ORDER BY created_at desc;`)
	if err != nil {
		log.Fatalf("query all tasks: %v", err)
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		var (
			taskID              uuid.UUID
			subject, link, body string
		)
		if err := rows.Scan(&taskID, &subject, &link, &body); err != nil {
			log.Fatalf("scan the tasks : %v", err)
		}
		tasks = append(tasks, &Task{
			Id:      taskID,
			Subject: subject,
			Link:    link,
			Body:    body,
		})
	}
	if err := rows.Close(); err != nil {
		log.Fatalf("rows close: %v", err)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("scan the tasks : %v", err)
	}
	res := map[string][]*Task{
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
