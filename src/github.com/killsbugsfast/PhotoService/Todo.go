package main

import "time"

type Todo struct {
	Id 	      int         `json:"id"`        // Stores the unique identifier of the Todo.
	Name      string      `json:"name"`      // Name of the Todo. Map this to the json 'name' as it is not idiomatic JSON to have uppercased keys
	Completed bool        `json:"completed"` // Flag to indicate whether the Todo is actually done. Map this to the json 'completed' as it is not idiomatic JSON to have uppercased keys
	Due       time.Time   `json:"due"`       // variable to store when we need the Todo to be completed. Map this to the json 'due' as it is not idiomatic JSON to have uppercased keys
}

// Create a slice (ordered collection) of Todo.
type Todos [] Todo