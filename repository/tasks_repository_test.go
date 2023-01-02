package repository

import (
	"crudzoo-go/domain"
	"crudzoo-go/infra"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTasksRepository_Find(t *testing.T) {
	mockDB := new(infra.MockDB)
	id, _ := uuid.Parse("f3ea7216-ed1c-4351-b723-cf0ccb8a7bcd")
	mockDB.On("GetTasks").Return([]infra.TaskDB{{
		TaskID:    id,
		Subject:   "subject",
		Link:      "link",
		Body:      "body",
		CreatedAt: time.Time{},
	}}, nil)
	target := TasksRepository{
		DB: mockDB,
	}
	actual, _ := target.Find()

	expected := []domain.Task{{
		Id: "f3ea7216-ed1c-4351-b723-cf0ccb8a7bcd",
		Content: domain.TaskContent{
			Subject: "subject",
			Link:    "link",
			Body:    "body",
		},
	}}
	assert.Equal(t, actual, expected)
}
