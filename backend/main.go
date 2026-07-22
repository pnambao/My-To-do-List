package main

import (
	_ "github.com/lib/pq"
	"database/sql"
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

type CreateTaskRequest struct {
	UserID int       `json:"user_id"`
	Title  string    `json:"title"`
	Date   time.Time `json:"date"`
}

type UpdateTaskRequest struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Date      time.Time `json:"date"`
	Completed bool      `json:"completed"`
}

var db *sql.DB

/*var tasks = []Task{
	{ID: 1, UserID: 1, Title: "Buy groceries", Date: time.Now(), Completed: false},
	{ID: 2, UserID: 1, Title: "Finish the project", Date: time.Now(), Completed: false},
	{ID: 3, UserID: 2, Title: "Call the plumber", Date: time.Now(), Completed: false},
	{ID: 4, UserID: 2, Title: "Buy a new phone", Date: time.Now(), Completed: false},
	{ID: 5, UserID: 3, Title: "Finish the report", Date: time.Now(), Completed: false},
	
}*/

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

	rows, err := db.Query(`
		SELECT id, user_id, title, date, completed
		FROM tasks
		ORDER BY date;
	`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {

		var task Task

		err := rows.Scan(
			&task.ID,
			&task.UserID,
			&task.Title,
			&task.Date,
			&task.Completed,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)

	}

	json.NewEncoder(w).Encode(tasks)

}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /tasks was called")
	var request CreateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec(`
		INSERT INTO tasks
		(user_id, title, date, completed)
		VALUES ($1,$2,$3,false)
	`,
		request.UserID,
		request.Title,
		request.Date,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	
	var request UpdateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec(`
		UPDATE tasks
			SET title=$1,
   			 date=$2,
   			 completed=$3
			WHERE id=$4;
	`,
		request.Title,
		request.Date,
		request.Completed,
		request.ID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := db.Exec(`
		DELETE FROM tasks
		WHERE id=$1;
	`,
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func main() {

// Connect to PostgreSQL
connStr := "host=localhost port=5432 user=postgres password=tadashi dbname=todo_app sslmode=disable"

var err error

db, err = sql.Open("postgres", connStr)
if err != nil {
	panic(err)
}

err = db.Ping()
if err != nil {
	panic(err)
}

fmt.Println("Connected to PostgreSQL!")

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
	
		case http.MethodGet:
			getTasksHandler(w, r)
	
		case http.MethodPost:
			createTaskHandler(w, r)

		case http.MethodPut:
			updateTaskHandler(w, r)
	
		case http.MethodDelete:
			deleteTaskHandler(w, r)
	
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	
		}
	
	})

	fmt.Println("Server running on http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}