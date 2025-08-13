package db

import (
	a_tasks "github.com/pseudoelement/lo-tasks/src/core/db/actions/tasks"
	"github.com/pseudoelement/lo-tasks/src/core/models/structs"
)

type DB struct {
	store   *StoreSchema
	actions StoreActions
}

func NewDB() *DB {
	store := &StoreSchema{tasks: make(map[string]structs.Task)}
	return &DB{store: store, actions: StoreActions{TasksActions: a_tasks.NewTaskActions(store)}}
}

func (this *DB) Actions() StoreActions {
	return this.actions
}
