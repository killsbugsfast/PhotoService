package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"

    // Add the MUX request router & dispatcher
    "github.com/gorilla/mux"
)

// Function for the / dispatcher
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

// Function for the /todos dispatcher (if the call is a GET)
func TodoIndex(w http.ResponseWriter, r *http.Request) {

    // Set the content type in the http response, and indicate that the response that will be returned will be a json
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    // Set the HTTP status in the header to be 200 OK
    w.WriteHeader(http.StatusOK)

    // Encode the json response and return it to the response writer. Panic if error is not null.
    if err := json.NewEncoder(w).Encode(todos); err != nil { // todos is defined in the repo.go file
        panic(err)
    }
}

// Function for the /todos dispatcher (if the call is a POST)
func TodoCreate(w http.ResponseWriter, r *http.Request) {

    // Define a local variable called todo
    var todo Todo

    // Read the reqeust body and store the contents in the body local variable
    // The first thing we do is open up the body of the request. Notice that we use io.LimitReader. This 
    // is a good way to protect against malicious attacks on your server. Imagine if someone wanted to send 
    // you 500GBs of json!
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

    // If there is an error in reading the request body, then panic
    if err != nil {
        panic(err)
    }

    // Try to close the body. If we error out as part of this operation, then panic
    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    // Unmarshall the json that was read in from the request body, and store the contents within the todo
    // local variable
    if err := json.Unmarshal(body, &todo); err != nil {

        // If we were unsuccessful in un-marshalling the HTTP body into the todo local variable, then we
        // need to return a response back to the requester. SEt the HTTP header content type to be a json.
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")

        // Indicate wtihin the HTTP response header the 422 - unprocessable entity status
        // code
        w.WriteHeader(422)
        
        // Encode the error that resulted from attempting to un-marchall the original HTTP request and 
        // service that back to the caller. If the encoding process errors out for whatever reason, then
        // panic.
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    // If we get to this point, then the unmarshalling of the HTTP request body was successful. Go to the
    // mock database, and call the function to create a new Todo struct and store it within the database
    t := RepoCreateTodo(todo)

    // Set the HTTP content type in the response to be a json object
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    // Set the HTTP response status to be Created (201)
    w.WriteHeader(http.StatusCreated)

    // Encode the marshalled todo, and send it back to the caller as the response. Note that this marshalled
    // response differs from the requester, in that this json object will also contain the ID / unique identifier
    // that our service would have set for this new Todo. If this operation errors out at all, then panic
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}

// Function for the /todos/{todoId} dispatcher
func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}