package main

import (
	"fmt"
	"time"
)

// addTask añade una nueva tarea a la lista.
func addTask(tasks *[]Task, content string) {
	id := 1
	if len(*tasks) > 0 {
		id = (*tasks)[len(*tasks)-1].ID + 1
	}
	newTask := Task{
		ID:        id,
		Content:   content,
		Completed: false,
		CreatedAt: time.Now(),
	}
	*tasks = append(*tasks, newTask)
}

// listTasks imprime la lista de tareas en la terminal.
func listTasks(tasks *[]Task) {
	if len(*tasks) == 0 {
		fmt.Println("No tasks to display.")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range *tasks {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		fmt.Printf("%s %d: %s\n", status, task.ID, task.Content)
	}
}

// completeTask marca una tarea como completada.
func completeTask(tasks *[]Task, id int) {
	for i, task := range *tasks {
		if task.ID == id {
			(*tasks)[i].Completed = true
			return
		}
	}
}

// deleteTask elimina una tarea de la lista.
func deleteTask(tasks *[]Task, id int) {
	for i, task := range *tasks {
		if task.ID == id {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			return
		}
	}
}
