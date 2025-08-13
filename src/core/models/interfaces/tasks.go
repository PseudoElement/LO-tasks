package interfaces

import "github.com/pseudoelement/lo-tasks/src/core/models/structs"

type ITaskActions interface {
	CreateTask(value structs.BasicTask) structs.Task
	RemoveTask(id string) bool
	GetTask(id string) (structs.Task, bool)
	GetTasks(completedFirst bool) []structs.Task
}
