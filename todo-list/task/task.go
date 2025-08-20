package task

import (
	"database/sql"
	"fmt"
	"log"
)

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func GetAllTasks(db *sql.DB) []Task{
	rows, _:= db.Query("SELECT id, task FROM todos")
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Text)
		tasks = append(tasks, t)
	}
	return tasks
}

func AddTask(db *sql.DB, task string) {
	stmt, err := db.Prepare("INSERT INTO todos(task) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(task)
	if err != nil {
		log.Fatal()
	}
	fmt.Println("✅Task Added.")
}

func ListTask(db *sql.DB) {
	rows, err := db.Query("SELECT id, task FROM todos ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\n📓Your Tasks:")
	for rows.Next() {
		var id int
		var task string
		err := rows.Scan(&id, &task)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d. %s\n", id, task)
	}
}

func DeleteTask(db *sql.DB, id int) {
	stmt, err := db.Prepare("DELETE FROM todos where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	count, _ := res.RowsAffected()
	if count > 0 {
		fmt.Println("🗑️ Task deleted.")
	} else {
		fmt.Println("⚠️ Task not found.")
	}
}
