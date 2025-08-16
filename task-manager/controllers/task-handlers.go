package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/waltertaya/dev-snippets/task-manager/db"
	"github.com/waltertaya/dev-snippets/task-manager/models"
)

func GetAndAddTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		rows, err := db.DB.Query("SELECT id, title, done FROM tasks")
		defer rows.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks := []models.Task{}
		for rows.Next() {
			var task models.Task
			if err := rows.Scan(&task.ID, &task.Title, &task.Done); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			tasks = append(tasks, task)
		}
		response := map[string]interface{}{"tasks": tasks}
		json.NewEncoder(w).Encode(response)
	}

	if r.Method == "POST" {
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := db.DB.Exec("INSERT INTO tasks (title, done) VALUES (?, ?)", task.Title, task.Done)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, _ := res.LastInsertId()
		task.ID = int(id)

		json.NewEncoder(w).Encode(task)
		return
	}

}

func UpdateOrDeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/tasks/"):] // extract task ID

	switch r.Method {
	case "PUT":
		var task models.Task
		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := db.DB.Exec("UPDATE tasks SET title=?, done=? WHERE id=?", task.Title, task.Done, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"updated": id})

	case "DELETE":
		_, err := db.DB.Exec("DELETE FROM tasks WHERE id=?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"deleted": id})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
