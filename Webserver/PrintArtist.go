// Webserver/PrintArtists.go

package groupie

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

type Artist struct {
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Id           int      `json:"id"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`
	ConcertDates string   `json:"concertDates"`
	Locations    string   `json:"locations"`
	Relations    string   `json:"relations"`
}

type ConcertDate struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Location struct {
	ID  int    `json:"id"`
	Loc string `json:"locations"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func PrintArtist() []Artist {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal("Error fetching artists:", err)
	}
	defer response.Body.Close()
	log.Println("Status code for artists:", response.StatusCode)

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code for artists: %d", response.StatusCode)
	}

	var artists []Artist
	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		log.Fatal("Error decoding artists:", err)
	}

	var wg sync.WaitGroup
	for i := range artists {
		wg.Add(1)
		go fetchArtistData(&artists[i], &wg)
	}
	wg.Wait()

	return artists
}

func fetchArtistData(artist *Artist, wg *sync.WaitGroup) {
	defer wg.Done()
	concertDates := fetchConcertDates(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", artist.Id))
	artist.ConcertDates = strings.Join(concertDates, ", ")

	locations := fetchLocations(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", artist.Id))
	artist.Locations = strings.Join(locations, ", ")

	relations := fetchRelations(fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", artist.Id))
	artist.Relations = formatRelations(relations)
}
