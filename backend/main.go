package main

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"
)

type healthResponse struct {
	Status string `json:"status"`
}

type Task struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	Completed bool      `json:"completed"`
}

var tasks = []Task{
	{ID: 1, UserID: 1, Title: "Buy groceries", Date: time.Now(), Completed: false},
	{ID: 2, UserID: 1, Title: "Finish the project", Date: time.Now(), Completed: false},
	{ID: 3, UserID: 2, Title: "Call the plumber", Date: time.Now(), Completed: false},
	{ID: 4, UserID: 2, Title: "Buy a new phone", Date: time.Now(), Completed: false},
	{ID: 5, UserID: 3, Title: "Finish the report", Date: time.Now(), Completed: false},
	
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := healthResponse {
		Status: "OK",
	}
	json.NewEncoder(w).Encode(response)
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/tasks", getTasksHandler)
	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}