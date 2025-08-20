package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aayushxadhikari/todo-list/task"
)

func StartServer(db *sql.DB){
	http.HandleFunc("/tasks", func(w http.ResponseWriter,r *http.Request){
		switch r.Method{
		case "GET":
			tasks:= task.GetAllTasks(db)
			json.NewEncoder(w).Encode(tasks)
		case "POST":
			var t task.Task
			err := json.NewDecoder(r.Body).Decode(&t)
			if err != nil{
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
			task.AddTask(db, t.Text)
			w.WriteHeader(http.StatusCreated)
		}
	})
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request){
		idStr := r.URL.Path[len("/tasks/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil{
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if r.Method == "DELETE"{
			task.DeleteTask(db, id)
			w.WriteHeader(http.StatusNoContent)
		}
	})
	http.ListenAndServe(":8080", nil)
}