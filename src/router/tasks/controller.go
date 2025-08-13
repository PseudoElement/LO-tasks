package tasks

import (
	"net/http"
)

func (this *TasksModule) SetRoutes() {
	http.HandleFunc("/tasks", this.tasksHandler)
	http.HandleFunc("/tasks/", this.taskHandler)
}
