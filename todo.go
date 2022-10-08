package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// welcome page function
// getting input from the console
// createTodo
// delete a singleTodo
// list all Todos created

type TODO struct {
	Id   string
	Todo string
}

var TODOLIST []TODO
var menuNumber int

func main() {
	welcome()

	mainMenu()

}

func mainMenu() {
	newLine(1)
	fmt.Println("Select operation")
	fmt.Println("1. Create Todo \t")
	fmt.Println("2. List Todo Item \t")
	fmt.Println("3. Edit Todo Item \t")
	fmt.Println("4. Delete Todo item")
	fmt.Println("0. Exit the program")
	_, err := fmt.Scan(&menuNumber)

	if err != nil {
		fmt.Println("Error:, please select the correct menu item")

	}

	switch menuNumber {
	case 1:
		createTodo()
	case 2:
		listTodo()
	case 3:
		updateTodo()
	case 4:
		deleteTodo()
	case 0:
		exitProgram()
	default:
		fmt.Println("Error: Invalid input")
	}
}

func newLine(numberOfLines int) {
	for i := 0; i < numberOfLines; {
		fmt.Println("\n")
		i++
	}
}

func welcome() {
	fmt.Println("*******{Welcome to the TODO CLI app}********")
	newLine(1)
}

func listTodo() {
	fmt.Println("list of item in your TODO list ")
	for index, item := range TODOLIST {
		fmt.Printf("TODO #%v: %v", index+1, item)
		newLine(1)
	}

	mainMenu()
}

func createTodo() {
	fmt.Println("Please enter your Todo item")
	var todoInput string
	_, err := fmt.Scan(&todoInput)
	if err != nil {
		fmt.Println("Error: Please enter a valid todo item")
	}

	var todo TODO
	todo.Id = strconv.Itoa(rand.Intn(100000))
	todo.Todo = todoInput

	TODOLIST = append(TODOLIST, todo)
	fmt.Println("Todo item added")
	mainMenu()
}

func updateTodo() {
	var updateItem string
	var updateInput string
	fmt.Println("Enter the id of the item you want to update: ")
	_, err := fmt.Scan(&updateItem)
	if err != nil {
		fmt.Println("Error:, please select the correct menu item")
	}
	for index, item := range TODOLIST {
		if item.Id == updateItem {
			TODOLIST = append(TODOLIST[:index], TODOLIST[index+1:]...)
			var newTodo TODO
			//Todo: display current todo item to user

			fmt.Println("Enter your todo: ")
			_, err := fmt.Scan(&updateInput)
			if err != nil {
				fmt.Println("Error:, please select the correct menu item")
			}

			newTodo.Id = updateItem
			newTodo.Todo = updateInput

			TODOLIST = append(TODOLIST, newTodo)

			fmt.Println("Todo item updated")
			mainMenu()
		}
	}

	fmt.Println("Todo item not found!")
	mainMenu()
}

func deleteTodo() {
	var deleteItem string

	fmt.Println("Enter the id of the todo item to delete:")
	_, err := fmt.Scan(&deleteItem)
	if err != nil {
		fmt.Println("Error: please select the correct menu item")
	}

	for index, item := range TODOLIST {
		if item.Id == deleteItem {
			TODOLIST = append(TODOLIST[:index], TODOLIST[index+1:]...)
			fmt.Println("Todo Item deleted")
			mainMenu()
		}
	}
	mainMenu()
}

func exitProgram() {
	fmt.Println("Exiting Todo app, Goodbye 👋🏾")
	os.Exit(0)
}
