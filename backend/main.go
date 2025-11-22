package main

import "net/http"
import "encoding/json"

// Task represents one item in the list
type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var tasks = []Task{} // in-memory store

func main() {
	http.HandleFunc("/", taskHandler)
	http.ListenAndServe(":5000", nil)
}


// taskHandler handles getting the task list and adding new tasks
func taskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(tasks)
		return
	}

	if r.Method == http.MethodPost {
		var newTask Task
		json.NewDecoder(r.Body).Decode(&newTask)

		if newTask.Title == "" || newTask.Description == "" {
			http.Error(w, "json is missing the title or description", http.StatusBadRequest)
			return
		}

		tasks = append(tasks, newTask)
		json.NewEncoder(w).Encode(newTask)
		return
	}

	http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
}
