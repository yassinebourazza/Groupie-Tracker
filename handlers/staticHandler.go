package groupietracker

import (
	"net/http"
	"os"
)

// handler static files and protect static directory
func HandleStatic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, "Method not allowed", http.StatusMethodNotAllowed)
		return

	} else {
		infos, err := os.Stat(r.URL.Path[1:])
		if err != nil {
			ErrorHandler(w, "Page not found", http.StatusNotFound)
			return
		} else if infos.IsDir() {
			ErrorHandler(w, "Access Forbidden !", http.StatusForbidden)
			return
		} else {
			http.ServeFile(w, r, r.URL.Path[1:])
		}
	}
}
