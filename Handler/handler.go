package GroupieTracker

import (
	"encoding/json"
	"net/http"
	"text/template"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	FirstResponse, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	defer FirstResponse.Body.Close()

	if err := json.NewDecoder(FirstResponse.Body).Decode(&FirstData); err != nil {
		http.Error(w, "Error decoding base API JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	SecondResponse, err := http.Get(FirstData.ArtistrLink)
	if err != nil {
		http.Error(w, "Error fetching artists data", http.StatusInternalServerError)
		return
	}
	defer SecondResponse.Body.Close()

	if err := json.NewDecoder(SecondResponse.Body).Decode(&Artist); err != nil {
		http.Error(w, "Error decoding base API JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("Html/index.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Artist)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
