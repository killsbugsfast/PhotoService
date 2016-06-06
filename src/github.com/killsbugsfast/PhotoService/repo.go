// This is a non-production / crude database. This will need to be replaced with something more robust later on.
package main

import "fmt"

// Keeps track of the last-used ID / index of Todos in our mock database
var currentId int

// Variable to store the list of Todos in the slice (defined in Todo.go)
var todos Todos

// Give us some seed data
func init() {
    RepoCreateTodo(Todo{Name: "Write presentation"})
    RepoCreateTodo(Todo{Name: "Host meetup"})
}

// Function - iterate through the Todos slice (defined in Todo.go), and find the one with the matching ID. Return
// that one to the caller.
func RepoFindTodo(id int) Todo {
    for _, t := range todos {
        if t.Id == id {
            return t
        }
    }
    // return empty Todo if not found
    return Todo{}
}

// Function - Add the new Todo to the database (i.e. at the end of the slice), with the next index #/value
func RepoCreateTodo(t Todo) Todo {
    currentId += 1
    t.Id = currentId
    todos = append(todos, t)
    return t
}

// Function - iterate through all of the Todos in the slice and remove the todo with the ID that was identified
// by the caller. Return nothing if it was successful; otherwise return an error back to the caller.
func RepoDestroyTodo(id int) error {
    for i, t := range todos {
        if t.Id == id {
            todos = append(todos[:i], todos[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}