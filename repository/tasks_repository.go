package repository

import (
	"crudzoo-go/domain"
	"crudzoo-go/infra"
)

type TasksRepository struct {
	DB infra.DB
}

func (t *TasksRepository) Find() ([]domain.Task, error) {
	data, _ := t.DB.GetTasks()
	var tasks []domain.Task
	for _, v := range data {
		tasks = append(tasks, domain.Task{
			Id: v.TaskID.String(),
			Content: domain.TaskContent{
				Subject: v.Subject,
				Link:    v.Link,
				Body:    v.Body,
			},
		})
	}
	return tasks, nil
}
