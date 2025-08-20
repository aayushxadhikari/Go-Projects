package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/aayushxadhikari/todo-list/api"
	"github.com/aayushxadhikari/todo-list/cmd"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Exec(`CREATE TABLE IF NOT EXISTS todos(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	task TEXT NOT NULL
	);`)

	// Decide Mode: CLI or API
	if len(os.Args) > 1&& os.Args[1] == "api"{
		log.Println("Starting the REST API on http://localhost:8080")
		api.StartServer(db)
	}else{
		cmd.RunCLI(db)
	}

	// // Create table if it doesnot exists
	// createTable := `
	// CREATE TABLE IF NOT EXISTS todos(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	task TEXT NOT NULL
	// );`
	// _, err = db.Exec(createTable)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// reader := bufio.NewReader(os.Stdin)
	// // CLI-Like interaction
	// for {
	// 	fmt.Println("\n===TODO-LIST===")
	// 	fmt.Println("1. Add list")
	// 	fmt.Println("2. View list")
	// 	fmt.Println("3. Delete list")
	// 	fmt.Println("4. Exit")
	// 	fmt.Println("Choose Option:")

	// 	// Read line and convert to integer
	// 	input, _ := reader.ReadString('\n')
	// 	input = strings.TrimSpace(input)
	// 	choice, err := strconv.Atoi(input)
	// 	if err != nil {
	// 		fmt.Println("❌ Invalid input. Please enter a number.")
	// 		continue
	// 	}

	// 	switch choice {
	// 	case 1:
	// 		fmt.Println("Enter Task:")
	// 		taskInput, _ := reader.ReadString('\n')
	// 		taskInput = strings.TrimSpace(taskInput)
	// 		task.AddTask(db, taskInput)
	// 	case 2:
	// 		task.ListTask(db)
	// 	case 3:
	// 		fmt.Println("Enter task id to delete:")
	// 		idInput, _ := reader.ReadString('\n')
	// 		idInput = strings.TrimSpace(idInput)
	// 		id, err := strconv.Atoi(idInput)
	// 		if err != nil {
	// 			fmt.Println("❌ Invalid ID. Please enter a number.")
	// 			continue
	// 		}
	// 		task.DeleteTask(db, id)
	// 	case 4:
	// 		fmt.Println("Goodbye!")
	// 		os.Exit(0)
	// 	default:
	// 		fmt.Println("Invalid Choice. Please Try Again.")
	// 	}
	// }

}
