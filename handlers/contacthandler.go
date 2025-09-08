package groupietracker

import (
	"html/template"
	"net/http"
)

//handle a page that shows contact informations about authors

func ContactHandler(w http.ResponseWriter , r *http.Request){
	if !(r.Method == http.MethodGet ) {
		ErrorHandler(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	temp,err:= template.ParseFiles("templates/contact.html")
	if err != nil{
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w,nil)
}