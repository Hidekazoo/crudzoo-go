package domain

type Task struct {
	Id      string
	content TaskContent
}

type TaskContent struct {
	Subject string
	Link    string
	Body    string
}
