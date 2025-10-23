package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Task represents a todo item
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

// In-memory storage for tasks
var tasks []Task
var nextID = 1

// CORS middleware
func enableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Max-Age", "86400")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}

// Get all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Create a new task
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newTask.ID = nextID
	newTask.CreatedAt = time.Now()
	nextID++

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

// Update a task
func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Update request: %s %s\n", r.Method, r.URL.Path) // Debug log

	idStr := r.URL.Path[len("/api/tasks/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("Invalid task ID: %s\n", idStr) // Debug log
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		fmt.Printf("JSON decode error: %v\n", err) // Debug log
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Updating task %d with: %+v\n", id, updatedTask) // Debug log

	for i, task := range tasks {
		if task.ID == id {
			// Preserve original ID and CreatedAt
			updatedTask.ID = id
			updatedTask.CreatedAt = task.CreatedAt
			tasks[i] = updatedTask

			fmt.Printf("Task updated successfully: %+v\n", updatedTask) // Debug log
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}

	fmt.Printf("Task not found: %d\n", id) // Debug log
	http.Error(w, "Task not found", http.StatusNotFound)
}

// Delete a task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/tasks/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func main() {
	// Initialize with some sample tasks
	tasks = []Task{
		{ID: 1, Title: "Learn Go", Completed: false, CreatedAt: time.Now()},
		{ID: 2, Title: "Build React app", Completed: true, CreatedAt: time.Now()},
	}
	nextID = 3

	// API routes with CORS middleware
	http.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w, r)
		if r.Method == "OPTIONS" {
			return
		}
		switch r.Method {
		case "GET":
			getTasks(w, r)
		case "POST":
			createTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/tasks/", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w, r)
		if r.Method == "OPTIONS" {
			return
		}
		switch r.Method {
		case "PUT":
			updateTask(w, r)
		case "DELETE":
			deleteTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Serve static files (React app)
	http.Handle("/", http.FileServer(http.Dir("./frontend/build")))

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
