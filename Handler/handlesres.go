package GroupieTracker

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type Data struct {
	ArtistrLink   string `json:"artists"`
	LocationsLink string `json:"locations"`
	DatesLink     string `json:"dates"`
	RelationLink  string `json:"relation"`
}

type Person struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
}

var (
	FirstData Data
	Artist    []Person
)

type relation struct {
	Rel map[string][]string `json:"datesLocations"`
}

var Index struct {
	Index1 []relation `json:"index"`
}

func HandlerRes(w http.ResponseWriter, r *http.Request) {
	ThirdRes, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
	}

	defer ThirdRes.Body.Close()

	if err := json.NewDecoder(ThirdRes.Body).Decode(&Index); err != nil {
		http.Error(w, "Error decoding base API JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	Id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/infos/"))
	if err != nil {
		http.Error(w, "Error extract Id", http.StatusInternalServerError)
	}

	tmpl, err := template.ParseFiles("Html/infos.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]any{
		"Artist":   Artist[Id-1],
		"Relation": Index.Index1[Id-1].Rel,
	})
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
