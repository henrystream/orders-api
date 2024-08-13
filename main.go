package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/hello", basicHandler)
	router.Post("/hello", basicHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to listen to server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if r.URL.Path == "/hello" {
			//handle to get foo
			fmt.Fprintf(w, "Welcome to hello page")
			//w.Write([]byte("Welcome to the foo page!"))
		}
		if r.Method == http.MethodPost {
			if r.URL.Path == "/hello" {
				fmt.Fprintf(w, "hello post")
			}
		}
	}
	//w.Write([]byte("Hello, world!"))
}
