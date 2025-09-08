package groupietracker

import (
	"html/template"
	"net/http"
)

//handerl about page

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if !(r.Method == http.MethodGet ) {
		ErrorHandler(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	temp, err := template.ParseFiles("templates/about.html")
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w,nil)
}
