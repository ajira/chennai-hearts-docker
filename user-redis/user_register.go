package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"encoding/json"
	"github.com/mediocregopher/radix.v2/redis"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	http.HandleFunc("/sign_up", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
    	var user User

    	err := decoder.Decode(&user)
    	if err != nil {
    		fmt.Println("error parsing request..", err)
    		fmt.Fprintf(w, "Failed to parse request")
    		return
    	}

    	err = saveUser(user)
    	if err != nil {
    		fmt.Fprintf(w, "Faild to connect to redis")
    		return
    	}

    	fmt.Println("Received: ", user)
		fmt.Fprintf(w, "<b>User " + user.Name + " registered!</b>")
	})

	fmt.Println("Listening to :", os.Getenv("HTTP_PORT"))
	err := http.ListenAndServe(":" + os.Getenv("HTTP_PORT"), handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	if err != nil {
		fmt.Println("Error starting server.", err)
	}
}

func saveUser(user User) error {
	redisURL := os.Getenv("REDIS_URL")
	_, err := redis.Dial("tcp", redisURL)
	if err != nil {
		fmt.Println("error connecting to redis." + redisURL + ". " + err.Error())
		return err
	}

	return nil
}
