package usecase

import (
	"crudzoo-go/domain"
	"github.com/stretchr/testify/mock"
)

type TasksInputPort interface {
	Find() ([]domain.Task, error)
}

type MockTasksInputPort struct {
	mock.Mock
}

func (m MockTasksInputPort) Find() ([]domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]domain.Task), args.Error(1)
}
