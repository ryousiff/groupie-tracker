// Webserver/webserver.go
package groupie

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

func WebServer() {
	// Load templates
	templates, err := LoadTemplates()
	if err != nil {
		log.Fatal("Error loading templates:", err)
	}
	// Serve static files
	http.Handle("/groupie.css", http.FileServer(http.Dir("../Webserver")))
	http.Handle("/info.css", http.FileServer(http.Dir("../Webserver")))
	http.Handle("/error.css", http.FileServer(http.Dir("../Webserver")))

	// http.HandleFunc("/filter", filter)
	// Define routes and handlers
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homeHandler(w, r, templates)
	})
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/filter", Filter)
	allData = returnArtists()
	SetData(allData)
	// Start the HTTP server
	fmt.Println("Server started on http://localhost:8800")
	err = http.ListenAndServe(":8800", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
	if allData == nil {
		fmt.Println("Failed to gather Data from API")
		os.Exit(1)
	}
}
func homeHandler(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	if r.URL.Path == "/" {
		artists := allData
		if artists == nil {
			// Error 500: Internal Server Error
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		renderTemplate(w, tmpl, artists)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}
}
func renderTemplate(w http.ResponseWriter, tmpl *template.Template, artists []Artist) {
	err := tmpl.Execute(w, artists)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func renderInfoTemplate(w http.ResponseWriter, artist Artist) {
	tmpl, err := template.ParseFiles("../Webserver/info.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		// Error 500: Internal Server Error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artist)
	if err != nil {
		log.Println("Error executing template:", err)
		// Error 500: Internal Server Error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func infoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Bad Request", http.StatusNotFound)
		return
	}
	artistID, err := strconv.Atoi(id)
	if err != nil {
		// Error 400: Bad Request
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	artists := allData
	if artists == nil {
		// Error 500: Internal Server Error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var selectedArtist Artist
	for _, artist := range artists {
		if artist.Id == artistID {
			selectedArtist = artist
			break
		}
	}
	if selectedArtist.Id == 0 {
		// Error 404: Not Found
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	renderInfoTemplate(w, selectedArtist)
}
