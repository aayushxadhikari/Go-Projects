package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){

	// checking for the argument and parsing the argument
	if len(os.Args) <2 {
		fmt.Println("Usage: todo [add|list|done|delete]")
		return 
	}
	// extracting the command 
	cmd := os.Args[1]

	// sending the data to add the task 
	switch cmd{
	case "add":
		if len(os.Args)<3{
			fmt.Println("Usage: todo add <task name>")
			return 
		}
		task := os.Args[2]
		if err := addTask(task); err != nil{
			fmt.Println("Error:",err)
		}else{
			fmt.Println("Task Added.")
		}
	// listing all teh list that have been added	
	case "list":
		if err := listTask(); err!=nil{
			fmt.Println("Error:", err)
		}
	// converts the task id to integer and marks it as "done"
	case "done":
		if len(os.Args)<3{
			fmt.Println("Usage: todo done <task id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil{
			fmt.Println("Invalid task ID.")
			return
		}
		if err := markDone(id); err!= nil{
			fmt.Println("Error:",err)
		}else{
			fmt.Println("Marked as Done!")
		}
	// converting the string to integer and then deleting the task
	case "delete":
		if len(os.Args)<3{
			fmt.Println("Usage: todo delete <task id>")
		}
		id, err:= strconv.Atoi(os.Args[2])
		if err != nil{
			fmt.Println("Invalid task ID")
			return
		}
		if err := deleteTask(id); err!= nil{
			fmt.Println("Error:", err)
		} else{
			fmt.Println("Task deleted")
		}
	default:
		fmt.Println("Unknown command:", cmd)
	}
}
