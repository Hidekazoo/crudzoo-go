package main

import (
	"crudzoo-go/handler"
	_ "github.com/jackc/pgx/v4/stdlib"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", handler.Tasks)
	_ = http.ListenAndServe(":8080", nil)
}
