package main

import (
    "net/http"
)

// Define a struc to store & manage all of the request routes and handling function.
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

// Define the path, hander functions, etc. for the application.
var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "TodoIndex",
        "GET",
        "/todos",
        TodoIndex,
    },
    Route{
        "TodoShow",
        "GET",
        "/todos/{todoId}",
        TodoShow,
    },
    Route{
        "TodoCreate",
        "POST",
        "/todos",
        TodoCreate,
    },
}