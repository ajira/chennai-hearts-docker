package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
)

const (
	HTTP_BIND = ":8080"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, os.Getenv("SAY_HELLO"))
	})

	fmt.Println("Binding to ", HTTP_BIND)
	err := http.ListenAndServe(HTTP_BIND, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	if err != nil {
		fmt.Println("Error starting server.", err)
	}
}
