package groupietracker

import (
	"html/template"
	"net/http"
	"strconv"
)

//Handle errors comming from the other handlers
func ErrorHandler(w http.ResponseWriter, s string, n int) {
	temp, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := map[string]string{
		"Status": strconv.Itoa(n),
		"Error": s,
	}
	w.WriteHeader(n)
	temp.Execute(w, data)
}
