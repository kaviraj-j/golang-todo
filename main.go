package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	todoTitle   string
	isCompleted bool
}

func main() {
	var tasks []Todo
	var choice int

	for {
		fmt.Println(`
		What do you want to do?
		1. Print all todos
		2. Add new todo
		3. Exit
		`)

		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			printAllTasks(tasks)

		case 2:
			var todoTitle = getNewTaskFromUser()
			tasks = append(tasks, Todo{todoTitle: todoTitle, isCompleted: false})
			fmt.Println("Todo added!")

		case 3:
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid choice, please enter 1, 2, or 3.")
		}
	}
}

func printAllTasks(tasks []Todo) {
	if len(tasks) == 0 {
		fmt.Println("No todos yet.")
	} else {
		fmt.Println("Your todos:")
		for i, task := range tasks {
			status := "Incomplete"
			if task.isCompleted {
				status = "Complete"
			}
			fmt.Printf("%d. %s [%s]\n", i+1, task.todoTitle, status)
		}
	}
}

func getNewTaskFromUser() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter todo title: ")
	todoTitle, _ := reader.ReadString('\n')
	trimmedTodoTitle := strings.TrimSpace(todoTitle)
	if len(trimmedTodoTitle) != 0 {
		return trimmedTodoTitle

	}
	fmt.Print("Don't leave an empty title\n")
	return getNewTaskFromUser()
}
