package main

import (
	"fmt"
	"net/http"
	"os"

	"url-shortener/internal/db"
	"url-shortener/internal/handlers"
)

func main() {
	db.ConnectDB()    //connects MongoDB
	db.ConnectRedis() //connects redis
	//http.HandleFunc("/", homeHandler)            // we will register the server here when user visits the / it will trigger the func homeFunction
	http.HandleFunc("/", handlers.RedirectHandler)     //redirect function
	http.HandleFunc("/shorten", handlers.URLShortener) // we will register the route in main

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println("Server Listening on port", port) // to print in the terminal

	err := http.ListenAndServe(":"+port, nil) //we start web server on port 8080
	if err != nil {
		fmt.Println("Server failed", err) //if server failed to start
	}
}
