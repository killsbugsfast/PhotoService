// Have the ability to log out web requests like most modern web serves. As there is no web logging package or functionality
// in the standard Go library, we have to create one ourselves
package main

import (
	"log"
	"net/http"
	"time"
)

// We are going to pass our handler to this function, which will then wrap the passed
// handler with loggin and timing functionality
func Logger (inner http.Handler, name string) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Here")
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
				"%s\t%s\t%s\t%s",
				r.Method, 
				r.RequestURI,
				name,
				time.Since(start),
		)
	})
}