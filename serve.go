package main

import (
	"net/http"
	"os"
	"time"
	"log"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
   	"fmt"
)


func main() {	
	fmt.Println("Init Frond-end Server")
    	pwd, dir_error := os.Getwd()
    	if dir_error != nil {
        	fmt.Println(dir_error)
        	os.Exit(1)
    	}
	fmt.Println("PWD VAR : " + pwd)
	r := mux.NewRouter()

	// Serve static assets directly.
	r.PathPrefix("/dist").Handler(http.FileServer(http.Dir("/go/src/github.com/Igor-Rast/bit/dist")))

	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	r.PathPrefix("/").HandlerFunc(IndexHandler(pwd +"/go/src/github.com/Igor-Rast/bit/dist"))

	srv := &http.Server{
		Handler: handlers.LoggingHandler(os.Stdout, r),
		Addr:    "0.0.0.0:" + "8040",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

        err := srv.ListenAndServe()
        if err != nil {
                log.Fatal("Error listening: ", err)
        }
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}


