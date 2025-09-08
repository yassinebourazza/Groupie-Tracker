package groupietracker

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

//Handle the informations comming from the api by id
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	Fetch(w)
	if !(r.Method == http.MethodGet ) {
		ErrorHandler(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/infos/"))
	if err != nil {
		ErrorHandler(w, "Page not found", http.StatusNotFound)
		return
	}
	if id < 1 || id > len(artists) {
		ErrorHandler(w, "Page not found", http.StatusNotFound)
		return
	}
	FetchById(w, id)
	temp, err := template.ParseFiles("templates/info.html")
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, map[string]any{
		"Artist":    artists[id-1],
		"Relations": RelationsList.Relations,
		"Locations": LocationsList.Locations,
		"Dates":     DatesList.Dates,
	})
}
