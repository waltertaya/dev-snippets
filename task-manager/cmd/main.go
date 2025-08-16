package main

import (
	"net/http"

	"github.com/waltertaya/dev-snippets/task-manager/controllers"
)

func main() {
	http.HandleFunc("/tasks", controllers.GetAndAddTasks)
	http.HandleFunc("/tasks/", controllers.UpdateOrDeleteTask)

	http.ListenAndServe(":8080", nil)
}
