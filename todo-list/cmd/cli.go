package cmd

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aayushxadhikari/todo-list/task"
)

func RunCLI(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== TODO LIST (CLI) ===")
		fmt.Println("1. Add task")
		fmt.Println("2. View tasks")
		fmt.Println("3. Delete task")
		fmt.Println("4. Exit")
		fmt.Print("Choose option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("❌ Invalid input.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter task: ")
			taskText, _ := reader.ReadString('\n')
			taskText = strings.TrimSpace(taskText)
			task.AddTask(db, taskText)
		case 2:
			task.ListTask(db)
		case 3:
			fmt.Print("Enter task ID to delete: ")
			idText, _ := reader.ReadString('\n')
			idText = strings.TrimSpace(idText)
			id, err := strconv.Atoi(idText)
			if err != nil {
				fmt.Println("❌ Invalid ID.")
				continue
			}
			task.DeleteTask(db, id)
		case 4:
			fmt.Println("👋 Exiting CLI...")
			return
		default:
			fmt.Println("❌ Invalid choice.")
		}
	}
}
