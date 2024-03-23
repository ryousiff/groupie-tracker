//Webserver/searchHandler.go
package groupie

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("search")
	if query == "" {
		// Error 400: Bad Request
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	artists := GatherDataUp("https://groupietrackers.herokuapp.com/api/artists")
	if artists == nil {
		// Error 500: Internal Server Error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var selectedArtists []Artist
	for _, artist := range artists {
		// Convert CreationDate to string for comparison
		creationDate := strconv.Itoa(artist.CreationDate)

		// Check if the query matches any of the artist's details
		if strings.HasPrefix(strings.ToLower(artist.Name), strings.ToLower(query)) ||
			strings.HasPrefix(strings.ToLower(artist.FirstAlbum), strings.ToLower(query)) ||
			strings.HasPrefix(strings.ToLower(creationDate), strings.ToLower(query)) {
			selectedArtists = append(selectedArtists, artist)
			continue // Skip checking members and locations if other criteria match
		}

		// Check if the query matches any member's name
		for _, member := range artist.Members {
			if strings.HasPrefix(strings.ToLower(member), strings.ToLower(query)) {
				selectedArtists = append(selectedArtists, artist)
				break // Break the loop once a match is found
			}
		}

		// Check if the query matches any location
		for _, location := range artist.Locations {
			if strings.HasPrefix(strings.ToLower(location), strings.ToLower(query)) { // Cast location to string
				selectedArtists = append(selectedArtists, artist)
				break // Break the loop once a match is found
			}
		}
	}

	if len(selectedArtists) == 0 {
		// Error 404: Not Found
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := LoadTemplates() // Assuming LoadTemplates is defined in the same package.
	if err != nil {
		log.Fatal("Error loading templates:", err)
	}
	renderTemplate(w, tmpl, selectedArtists)
}
