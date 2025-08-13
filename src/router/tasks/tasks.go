package tasks

import (
	"github.com/pseudoelement/lo-tasks/src/core/db"
	"github.com/pseudoelement/lo-tasks/src/core/logger"
)

type TasksModule struct {
	db     *db.DB
	logger *logger.Logger
}

func NewTasksModule(db *db.DB, logger *logger.Logger) *TasksModule {
	return &TasksModule{db: db, logger: logger}
}
