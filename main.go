package main

import (
	"log"
	"net/http"

	"github.com/pseudoelement/lo-tasks/src/core/db"
	"github.com/pseudoelement/lo-tasks/src/core/logger"
	"github.com/pseudoelement/lo-tasks/src/router/tasks"
)

func main() {
	appLogger := logger.NewLogger()
	db := db.NewDB()

	tasksModule := tasks.NewTasksModule(db, appLogger)

	go appLogger.Listen()

	tasksModule.SetRoutes()

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
