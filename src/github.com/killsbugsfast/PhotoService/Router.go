package main

import (
    "net/http"

    // Add the MUX request router & dispatcher
    "github.com/gorilla/mux"
)

// Define a function, iterate through all of the routes that have been defined in the route variable,
// and set-up appropriate routes for each for the routes within the router
func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {

        // Wrap our handler with the logger functionality.
        var handler http.Handler
        handler = route.HandlerFunc // The handling function that will need to service the request (defined in var routes below)
        handler = Logger (handler, route.Name) // Wrap it around our logger.

        router.
            Methods(route.Method). // HTTP Method (i.e. GET or POST)
            Path(route.Pattern). // URL PATH or pattern that the handler will need to service ("/whatever")
            Name(route.Name). // The name of the route
            Handler(handler) // The wrapped Handler function that will be servicing the request
    }

    return router
}