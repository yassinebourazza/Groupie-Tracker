package GroupieTracker

import (
	"encoding/json"
	"net/http"
	"text/template"
)

type Locations struct {
	Location []string `json:"locations"`
}

type Index struct {
	Index Locations `json:"index"`
}

var Index1 Locations

func HandlerRes(w http.ResponseWriter, r *http.Request) {
	LocationRes, err := http.Get(FirstData.LocationsLink)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	defer LocationRes.Body.Close()

	if err := json.NewDecoder(LocationRes.Body).Decode(&Index1); err != nil {
		http.Error(w, "Error decoding base API JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Artist[].Locations = Index1.Location

	tmpl, err := template.ParseFiles("Html/infos.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Artist[Id-1])
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
