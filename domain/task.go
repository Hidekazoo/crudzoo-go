package domain

type Task struct {
	Id      string
	Content TaskContent
}

type TaskContent struct {
	Subject string
	Link    string
	Body    string
}
