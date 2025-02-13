package todo

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	ID   int
	Task string
}

var todos []Todo
var currentID int
var csvFile = "todos.csv"

func StartTodo() {
	loadTodosFromCSV()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nTodo Application")
		fmt.Println("1. Add Todo")
		fmt.Println("2. View Todos")
		fmt.Println("3. Update Todo")
		fmt.Println("4. Delete Todo")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			addTodo(reader)
		case "2":
			viewTodos()
		case "3":
			updateTodo(reader)
		case "4":
			deleteTodo(reader)
		case "5":
			fmt.Println("Exiting...")
			saveTodosToCSV()
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func loadTodosFromCSV() {
	file, err := os.Open(csvFile)
	if err != nil {
		// If the file doesn't exist, create it
		if os.IsNotExist(err) {
			file, _ = os.Create(csvFile)
			file.Close()
			return
		}
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		task := record[1]
		todos = append(todos, Todo{ID: id, Task: task})
		if id > currentID {
			currentID = id
		}
	}
}

func saveTodosToCSV() {
	file, err := os.Create(csvFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, todo := range todos {
		record := []string{strconv.Itoa(todo.ID), todo.Task}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}
	}
}

func addTodo(reader *bufio.Reader) {
	fmt.Print("Enter task: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	currentID++
	todo := Todo{
		ID:   currentID,
		Task: task,
	}
	todos = append(todos, todo)
	fmt.Println("Todo added successfully!")
	saveTodosToCSV()
}

func viewTodos() {
	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	fmt.Println("\nYour Todos:")
	for _, todo := range todos {
		fmt.Printf("%d: %s\n", todo.ID, todo.Task)
	}
}

func updateTodo(reader *bufio.Reader) {
	viewTodos()
	fmt.Print("Enter the ID of the todo to update: ")
	idInput, _ := reader.ReadString('\n')
	idInput = strings.TrimSpace(idInput)

	var id int
	_, err := fmt.Sscanf(idInput, "%d", &id)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			fmt.Print("Enter new task: ")
			newTask, _ := reader.ReadString('\n')
			newTask = strings.TrimSpace(newTask)

			todos[i].Task = newTask
			fmt.Println("Todo updated successfully!")
			saveTodosToCSV()
			return
		}
	}

	fmt.Println("Todo not found.")
}

func deleteTodo(reader *bufio.Reader) {
	viewTodos()
	fmt.Print("Enter the ID of the todo to delete: ")
	idInput, _ := reader.ReadString('\n')
	idInput = strings.TrimSpace(idInput)

	var id int
	_, err := fmt.Sscanf(idInput, "%d", &id)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Println("Todo deleted successfully!")
			saveTodosToCSV()
			return
		}
	}

	fmt.Println("Todo not found.")
}
