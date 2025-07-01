package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var taskFile = "tasks.json"

func loadTasks() ([]Task, error) {
	var tasks []Task
	file, err := os.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func saveTasks(tasks []Task) error{
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil{
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}

func addTask(title string) error{
	tasks, err := loadTasks()
	if err != nil{
		return err
	}
	id := 1
	if len(tasks)>0{
		id = tasks[len(tasks)-1].ID +1
	}
	tasks = append(tasks, Task{ID:id, Title:title, Done:false})
	return saveTasks(tasks)
}

func listTask() error{
	tasks, err := loadTasks()
	if err!= nil{
		return err
	}
	if len(tasks) == 0{
		fmt.Println("No tasks yet.")
	}
	for _, t := range tasks{
		status :=""
		if t.Done{
			status = "x"
		}
		fmt.Println("[%s] %d: %s\n", status, t.ID, t.Title)
	}
	return nil
}

func markDone(id int)error{
	tasks, err := loadTasks()
	if err != nil{
		return err
	}
	for i, t := range tasks{
		if t.ID == id{
			tasks[i].Done = true
			return saveTasks(tasks)
		}
	}
	return fmt.Errorf("Task %d not found", id)
}

func deleteTask(id int) error{
	tasks, err := loadTasks()
	if err!= nil{
		return err
	}
	var updated []Task
	found := false
	for _, t:= range tasks{
		if t.ID!= id{
			updated = append(updated, t)
		}else{
			found = true
		}
	}
	if !found {
		return fmt.Errorf("Tasks %d not found", id)
	}
	return saveTasks(updated)
}