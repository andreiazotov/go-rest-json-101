package main

import "fmt"

var currentID int
var todos Todos

func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

// RepoFindTodo Exported
func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}

	return Todo{}
}

// RepoCreateTodo Exported
func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

// RepoDestroyTodo Exported
func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[:i+1]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
