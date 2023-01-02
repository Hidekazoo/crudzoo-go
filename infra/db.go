package infra

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/mock"
	"log"
	"time"
)

type TaskDB struct {
	TaskID    uuid.UUID
	Subject   string
	Link      string
	Body      string
	CreatedAt time.Time
}

type DB interface {
	GetTasks() ([]TaskDB, error)
}

type DBImple struct{}
type Config struct {
	DBSetting string `envconfig:"DB_SETTING" required:"true"`
}

func (d *DBImple) GetTasks() ([]TaskDB, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		log.Fatalf("env error %v", err)
	}

	db, err := sql.Open("pgx", c.DBSetting)
	if nil != err {
		log.Fatal("infra connected error")
	}

	rows, err := db.Query(`SELECT task_id, subject, link, body FROM tasks ORDER BY created_at desc;`)
	if err != nil {
		log.Fatalf("query all tasks: %v", err)
	}
	defer rows.Close()

	var tasks []TaskDB
	for rows.Next() {
		var (
			taskID              uuid.UUID
			subject, link, body string
		)
		if err := rows.Scan(&taskID, &subject, &link, &body); err != nil {
			log.Fatalf("scan the tasks : %v", err)
		}
		tasks = append(tasks, TaskDB{
			TaskID:  taskID,
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
	return tasks, nil
}

type MockDB struct {
	mock.Mock
}

func (m MockDB) GetTasks() ([]TaskDB, error) {
	args := m.Called()
	return args.Get(0).([]TaskDB), args.Error(1)
}
