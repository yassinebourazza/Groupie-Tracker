package groupietracker

import (
	"html/template"
	"net/http"
)

//Handle and shows the home page 
func HandlHome(w http.ResponseWriter, r *http.Request) {
	Fetch(w)

	if r.URL.Path != "/" {
		ErrorHandler(w, "Page not found", http.StatusNotFound)
		return
	}
	if !(r.Method == http.MethodGet ) {
		ErrorHandler(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	temp.Execute(w,artists)
}
