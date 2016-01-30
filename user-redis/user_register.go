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

		if r.Method == "POST" {
			err := saveUser(w, r)
			if err != nil {
				fmt.Fprintf(w, err.Error())
				return
			}

			fmt.Fprintf(w, "<b>User registered!</b>")
		} else if r.Method == "GET" {

		}
	})

	fmt.Println("Listening to :", os.Getenv("HTTP_PORT"))
	err := http.ListenAndServe(":" + os.Getenv("HTTP_PORT"), handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	if err != nil {
		fmt.Println("Error starting server.", err)
	}
}

func saveUser(w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var user User

	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("error parsing request..", err)
		fmt.Fprintf(w, "Failed to parse request")
		return err
	}

	redisURL := os.Getenv("REDIS_URL")
	client, err := redis.Dial("tcp", redisURL)
	if err != nil {
		fmt.Println("error connecting to redis." + redisURL + ". " + err.Error())
		return err
	}

	user_json, _ := json.Marshal(user)
	err = client.Cmd("SET", user.Name, user_json).Err
	if err != nil {
		fmt.Println("unable to save to redis" + err.Error())
		return err
	}

	return nil
}
