package main

import (
    "log"
    "net/http"
)

func main() {

    // Instantiate a new router (this is defined in Routes.go)
    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}
