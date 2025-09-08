package main

import (
	"fmt"
	groupietracker "groupietracker/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/",groupietracker.HandlHome)
	http.HandleFunc("/infos/",groupietracker.InfoHandler)
	http.HandleFunc("/static/",groupietracker.HandleStatic)
	http.HandleFunc("/about",groupietracker.AboutHandler)
	http.HandleFunc("/contact",groupietracker.ContactHandler)
	fmt.Print("the server work on: http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Error", err)
	}
}
