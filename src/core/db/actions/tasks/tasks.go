package a_tasks

import (
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/pseudoelement/lo-tasks/src/core/models/interfaces"
	"github.com/pseudoelement/lo-tasks/src/core/models/structs"
)

type TaskActions struct {
	mu    sync.RWMutex
	store interfaces.IStoreSchema
}

func NewTaskActions(store interfaces.IStoreSchema) *TaskActions {
	return &TaskActions{store: store}
}

func (this *TaskActions) CreateTask(basicTask structs.BasicTask) structs.Task {
	this.mu.Lock()
	defer this.mu.Unlock()

	id := uuid.New()
	tasks := this.store.Tasks()
	task := structs.Task{
		BasicTask: basicTask,
		CreatedAt: time.Now().Format("02-01-2006 15:04"),
		Id:        id.String(),
	}

	tasks[id.String()] = task

	return task
}

func (this *TaskActions) RemoveTask(id string) bool {
	startLen := len(this.store.Tasks())

	this.mu.Lock()
	defer this.mu.Unlock()

	delete(this.store.Tasks(), id)

	newLen := len(this.store.Tasks())

	return startLen != newLen
}

func (this *TaskActions) GetTask(id string) (structs.Task, bool) {
	this.mu.RLock()
	defer this.mu.RUnlock()

	tasks := this.store.Tasks()
	val, ok := tasks[id]

	return val, ok
}

func (this *TaskActions) GetTasks(completedFirst bool) []structs.Task {
	var tasks []structs.Task = make([]structs.Task, len(this.store.Tasks()))
	var idx int
	for _, task := range this.store.Tasks() {
		tasks[idx] = task
		idx++
	}

	sort.Slice(tasks, func(i, j int) bool {
		if completedFirst {
			return tasks[i].Completed
		} else {
			return !tasks[i].Completed
		}
	})

	return tasks
}

var _ interfaces.ITaskActions = (*TaskActions)(nil)
