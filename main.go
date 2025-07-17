package main

import (
	"fmt"
	"net/http"
	"GroupieTracker/Handler" // Adjust the import path as necessary
)

func main() {
	fs := http.FileServer(http.Dir("CSS"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", fs))
	http.HandleFunc("/", GroupieTracker.Handler)
	http.HandleFunc("/infos/", GroupieTracker.HandlerRes)
	fmt.Println("Server is running on port http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
