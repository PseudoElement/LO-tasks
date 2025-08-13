package db

import (
	"github.com/pseudoelement/lo-tasks/src/core/models/interfaces"
	"github.com/pseudoelement/lo-tasks/src/core/models/structs"
)

type StoreSchema struct {
	tasks map[string]structs.Task
}

func (s *StoreSchema) Tasks() map[string]structs.Task {
	return s.tasks
}

type StoreActions struct {
	TasksActions interfaces.ITaskActions
}
