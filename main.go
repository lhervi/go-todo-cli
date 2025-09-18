package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// Define los "flags" para los comandos.
	add := flag.Bool("add", false, "Add new task")
	list := flag.Bool("list", false, "List all task")
	taskContent := flag.String("task", "", "The content of the task of add")	

	flag.Parse()

	//Cargar las tareas existentes
	tasks, err := LoadTasks("tasks.json")

	if err != nil {
		fmt.Println("Error loading tasks: ", err)
		os.Exit(1)
	}

	//Lógica de los comandos
	switch {
	case *add:
		if *taskContent == "" {
			fmt.Println("Error: no task content provided.")
			os.Exit(1)
		}

		addTask (&tasks, *taskContent)
		if err := SaveTasks("tasks.json", tasks); err != nil {
			fmt.Prinln("Error saving tasks:", err)
			os.Exit(1)
		}
		fmt.Println("Task added successfully.")
	case *list:
		listTasks(&tasks)		
	default:
		fmt.Println("No command specified. Use -aa or -list.")
		os.Exit(1)
	}

	func addTask(tasks *[]Task, content string) {
		id:= 1
		if len(*tasks) > 0 {
			id = (*tasks)[len(*tasks)-1].ID + 1
		}
		newTask := Task {
			ID: id,
			Content: content,
			Completed: false,
			CreatedAt: time.Now(),

		}
		*tasks = append(*tasks, bewTask)
	}

	// listTasks imprime la lista de tareas en la terminal.
	func listTasks(tasks *[]Task) {
		if len(*tasks) == 0 {
			fmt.Println("No tasks to display.")
			return
		}
		fmt.Println("Tasls:")
		for _, task := range *tasks {
			status := "[ ]"
			if task.Completed {
				status = "[x]"
			}
			fmt.Printf("%s %d: %s\n", status, task.ID, task.Content)
		}
	}


}