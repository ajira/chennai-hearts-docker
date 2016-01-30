package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<b>Hello!</b>")
	})

	fmt.Println("Listening to :", os.Getenv("HTTP_PORT"))
	err := http.ListenAndServe(":" + os.Getenv("HTTP_PORT"), handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	if err != nil {
		fmt.Println("Error starting server.", err)
	}
}
