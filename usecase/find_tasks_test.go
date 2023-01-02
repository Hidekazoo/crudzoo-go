package usecase

import (
	"crudzoo-go/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTasks(t *testing.T) {
	tasksInputPort := new(MockTasksInputPort)
	tasks := []domain.Task{{}}
	tasksInputPort.On("Find").Return(tasks, nil)
	target, _ := FindTasks(tasksInputPort)
	assert.Equal(t, target, tasks)
}
