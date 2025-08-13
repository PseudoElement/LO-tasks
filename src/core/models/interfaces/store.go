package interfaces

import "github.com/pseudoelement/lo-tasks/src/core/models/structs"

type IStoreSchema interface {
	Tasks() map[string]structs.Task
}
