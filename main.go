package main

import (
	"crudzoo-go/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", handler.Tasks)
	http.ListenAndServe(":8888", nil)

}
