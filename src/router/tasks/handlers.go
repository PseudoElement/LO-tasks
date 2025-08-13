package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pseudoelement/lo-tasks/src/core/models/structs"
	"github.com/pseudoelement/lo-tasks/src/utils"
)

func (this *TasksModule) tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		this.getAllTasks(w, r)
	case http.MethodPost:
		this.createTask(w, r)
	default:
		utils.FailResponse(w, "Method not allowed.", http.StatusBadRequest)
	}
}

func (this *TasksModule) taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		this.getTaskByID(w, r)
	default:
		utils.FailResponse(w, "Method not allowed.", http.StatusBadRequest)
	}
}

func (this *TasksModule) getAllTasks(w http.ResponseWriter, r *http.Request) {
	completedStr := r.URL.Query().Get("completedFirst")
	completedFirst, _ := strconv.ParseBool(completedStr)

	tasks := this.db.Actions().TasksActions.GetTasks(completedFirst)
	go this.logger.LogGetAllTasks(completedFirst)

	utils.SuccessResponse(w, tasks, http.StatusCreated)
}

func (this *TasksModule) getTaskByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/tasks/"):]
	task, found := this.db.Actions().TasksActions.GetTask(id)

	if !found {
		msg := fmt.Sprintf("Task with id %s not found.", id)
		go this.logger.LogInvalidRequest(r)
		utils.FailResponse(w, msg, http.StatusBadRequest)

		return
	}

	go this.logger.LogGetTaskByID(task)

	utils.SuccessResponse(w, task, http.StatusOK)
}

func (this *TasksModule) createTask(w http.ResponseWriter, r *http.Request) {
	var basicTask structs.BasicTask
	err := json.NewDecoder(r.Body).Decode(&basicTask)
	if err != nil {
		go this.logger.LogInvalidRequest(r)
		utils.FailResponse(w, "Invalid request body", http.StatusBadRequest)

		return
	}

	if basicTask.Creator == "" || basicTask.Info == "" {
		go this.logger.LogInvalidRequest(r)
		utils.FailResponse(w, "Creator and Info are required.", http.StatusBadRequest)

		return
	}

	task := this.db.Actions().TasksActions.CreateTask(basicTask)
	go this.logger.LogCreateTask(task)

	utils.SuccessResponse(w, task, http.StatusCreated)
}
