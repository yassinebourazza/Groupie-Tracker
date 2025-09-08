package groupietracker

import (
	"encoding/json"
	"net/http"
)

type artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationsUrl    string   `json:"locations"`
	ConcertDatesUrl string   `json:"concertDates"`
	RelationsUrl    string   `json:"relations"`
}

type dates struct {
	Dates []string `json:"dates"`
}

type location struct {
	Locations []string `json:"locations"`
}

type relation struct {
	Relations map[string][]string `json:"datesLocations"`
}

var LocationsList location

var DatesList dates

var RelationsList relation

var artists []artist

//fetch data comming from the api
func Fetch(w http.ResponseWriter) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

//fetch data by id 
func FetchById(w http.ResponseWriter, id int) {
	RelationsResponse, err := http.Get(artists[id-1].RelationsUrl)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(RelationsResponse.Body).Decode(&RelationsList)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	LocationsResponse, err := http.Get(artists[id-1].LocationsUrl)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(LocationsResponse.Body).Decode(&LocationsList)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	DatesResponse, err := http.Get(artists[id-1].ConcertDatesUrl)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.NewDecoder(DatesResponse.Body).Decode(&DatesList)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
