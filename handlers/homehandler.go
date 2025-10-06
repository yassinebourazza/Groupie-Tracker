package groupietracker

import (
	"html/template"
	"net/http"
)

var classicRock = []string{
	"Queen",
	"Pink Floyd",
	"Led Zeppelin",
	"The Jimi Hendrix Experience",
	"The Rolling Stones",
	"ACDC",
	"Scorpions",
	"Aerosmith",
	"Deep Purple",
	"Genesis",
	"Dire Straits",
	"Bee Gees",
}

var hardRock = []string{
	"Metallica",
	"Guns N' Roses",
	"Nickelback",
}

var ModernRock = []string{
	"Pearl Jam",
	"U2",
	"Linkin Park",
	"Red Hot Chili Peppers",
	"Green Day",
	"Coldplay",
	"Maroon 5",
	"Muse",
	"Foo Fighters",
	"Gorillaz",
	"Arctic Monkeys",
	"Fall Out Boy",
	"Imagine Dragons",
	"Thirty Seconds to Mars",
	"Twenty One Pilots",
}

var hiphop = []string{
	"Kendrick Lamar",
	"Eminem",
	"XXXTentacion",
	"Juice Wrld",
	"Mac Miller",
	"Logic",
	"J. Cole",
	"Joyner Lucas",
	"Post Malone",
	"Travis Scott",
	"Mobb Deep",
	"NWA",
}

var pop = []string{
	"Katy Perry",
	"Rihanna",
	"The Chainsmokers",
}

var diverseGenres = []string{
	"SOJA",
	"Mamonas Assassinas",
	"Phil Collins",
	"Bobby McFerrins",
	"Alec Benjamin",
	"R3HAB",
}

// Handle and shows the home page
func HandlHome(w http.ResponseWriter, r *http.Request) {
	Fetch(w)

	if r.URL.Path != "/" {
		ErrorHandler(w, "Page not found", http.StatusNotFound)
		return
	}

	option := r.FormValue("MusicType")
	var compare []string

	switch option {
	case "Hip-Hop":
		compare = hiphop
	case "Pop":
		compare = pop
	case "Classic Rock":
		compare = classicRock
	case "Hard Rock":
		compare = hardRock
	case "Modern Rock":
		compare = ModernRock
	case "Diverse Genres":
		compare = diverseGenres
	default:
		compare = nil
	}

	var someArtist []artist
	if compare != nil {
		for _, artist := range artists {
			if TheArtistExist(artist.Name, compare) {
				someArtist = append(someArtist, artist)
			}
		}
	} else {
		someArtist = append(someArtist, artists...)
	}

	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	temp.Execute(w, someArtist)
}

func TheArtistExist(artistName string, compare []string) bool {
	for _, artist := range compare {
		if artistName == artist {
			return true
		}
	}
	return false
}
