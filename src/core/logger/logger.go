package logger

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/pseudoelement/lo-tasks/src/core/models/structs"
)

type LoggerMsg struct {
	action string
	data   []byte
}

type InvalidReqData struct {
	Url   string
	Query string
	Body  string
}

type GetAllTasksData struct {
	CompletedFirst bool
}

type Logger struct {
	loggerChan chan LoggerMsg
}

func NewLogger() *Logger {
	n := &Logger{
		loggerChan: make(chan LoggerMsg),
	}

	return n
}

func (l *Logger) Chan() <-chan LoggerMsg {
	return l.loggerChan
}

func (l *Logger) Listen() {
	for msg := range l.Chan() {
		log.Printf("[logger %s] data - %+v\n", msg.action, string(msg.data))
	}
}

func (l *Logger) LogInvalidRequest(req *http.Request) {
	bodyBytes, _ := io.ReadAll(req.Body)
	b, _ := json.Marshal(InvalidReqData{
		Url:   req.URL.Host + req.URL.Path,
		Query: req.URL.RawQuery,
		Body:  string(bodyBytes),
	})
	l.loggerChan <- LoggerMsg{action: "InvalidRequest", data: b}
}

func (l *Logger) LogCreateTask(task structs.Task) {
	b, _ := json.Marshal(task)
	l.loggerChan <- LoggerMsg{action: "CreateTask", data: b}
}

func (l *Logger) LogGetTaskByID(task structs.Task) {
	b, _ := json.Marshal(task)
	l.loggerChan <- LoggerMsg{action: "GetTaskByID", data: b}
}

func (l *Logger) LogGetAllTasks(completedFirst bool) {
	b, _ := json.Marshal(GetAllTasksData{CompletedFirst: completedFirst})
	l.loggerChan <- LoggerMsg{action: "GetAllTasks", data: b}
}
