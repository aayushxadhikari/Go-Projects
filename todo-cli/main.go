package main

import (
	"fmt"
	"os"
)

func main(){

	// checking for the argument and parsing the argument
	if len(os.Args) <2 {
		fmt.Println("Usage: todo [add|list|done|delete]")
		return 
	}
	cmd := os.Args[1]

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
	case "list":
		if err := listTask(); err!=nil{
			fmt.Println("Error:", err)
		}
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
	

	}
}
