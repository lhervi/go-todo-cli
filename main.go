package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	add := flag.Bool("add", false, "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	taskContent := flag.String("task", "", "The content of the task to add")
	completeID := flag.Int("complete", 0, "Complete a task by its ID")
	deleteID := flag.Int("delete", 0, "Delete a task by its ID")

	flag.Parse()

	tasks, err := LoadTasks("tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}

	switch {
	case *add:
		if *taskContent == "" {
			fmt.Println("Error: No task content provided.")
			os.Exit(1)
		}
		addTask(&tasks, *taskContent)
		if err := SaveTasks("tasks.json", tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			os.Exit(1)
		}
		fmt.Println("Task added successfully.")
	case *list:
		listTasks(&tasks)
	case *completeID > 0:
		completeTask(&tasks, *completeID)
		if err := SaveTasks("tasks.json", tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			os.Exit(1)
		}
		fmt.Printf("Task %d completed successfully.\n", *completeID)
	case *deleteID > 0:
		deleteTask(&tasks, *deleteID)
		if err := SaveTasks("tasks.json", tasks); err != nil {
			fmt.Println("Error saving tasks:", err)
			os.Exit(1)
		}
		fmt.Printf("Task %d deleted successfully.\n", *deleteID)
	default:
		fmt.Println("No command specified. Use -add, -list, -complete or -delete.")
		os.Exit(1)
	}
}
