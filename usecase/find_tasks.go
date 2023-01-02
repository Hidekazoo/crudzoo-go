package usecase

import (
	"crudzoo-go/domain"
)

func FindTasks(tasksInputPort TasksInputPort) ([]domain.Task, error) {
	return tasksInputPort.Find()
}
