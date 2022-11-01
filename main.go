package main

import (
	"crudzoo-go/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", handler.Tasks)
	_ = http.ListenAndServe(":8080", nil)
}
