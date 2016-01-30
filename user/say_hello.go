package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<b>Hello!</b>")
	})
	fmt.Println("HTTP_PORT:", os.Getenv("HTTP_PORT"))
	http.ListenAndServe(":" + os.Getenv("HTTP_PORT"), nil)
}
