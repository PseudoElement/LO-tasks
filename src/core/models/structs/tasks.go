package structs

type BasicTask struct {
	Creator   string `json:"creator"`
	Info      string `json:"info"`
	Completed bool   `json:"completed"`
}

type Task struct {
	BasicTask

	CreatedAt string `json:"created_at"`
	Id        string `json:"id"`
}
