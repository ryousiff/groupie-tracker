//Webserver/fetch.go
package groupie

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func fetchConcertDates(apiURL string) []string {
	response, err := http.Get(apiURL)
	if err != nil {
		log.Println("Error fetching concert dates:", err)
		return nil
	}
	defer response.Body.Close()
	log.Println("Status code for concert dates:", response.StatusCode)

	if response.StatusCode != http.StatusOK {
		// Error 500: Internal Server Error
		log.Printf("Unexpected status code for concert dates: %d", response.StatusCode)
		return nil
	}

	var concertDates struct {
		Dates []string `json:"dates"`
	}

	err = json.NewDecoder(response.Body).Decode(&concertDates)
	if err != nil {
		log.Println("Error decoding concert dates:", err)
		return nil
	}

	return concertDates.Dates
}

func fetchLocations(apiURL string) []string {
	response, err := http.Get(apiURL)
	if err != nil {
		log.Println("Error fetching locations:", err)
		return nil
	}
	defer response.Body.Close()
	log.Println("Status code for locations:", response.StatusCode)

	if response.StatusCode != http.StatusOK {
		// Error 500: Internal Server Error
		log.Printf("Unexpected status code for locations: %d", response.StatusCode)
		return nil
	}

	var locations1 struct {
		Loc []string `json:"locations"`
	}

	err = json.NewDecoder(response.Body).Decode(&locations1)
	if err != nil {
		log.Println("Error decoding locations:", err)
		return nil
	}

	return locations1.Loc
}

func fetchRelations(apiURL string) Relation {
	response, err := http.Get(apiURL)
	if err != nil {
		log.Println("Error fetching Relations:", err)
		return Relation{}
	}
	defer response.Body.Close()
	log.Println("Status code for Relations:", response.StatusCode)

	if response.StatusCode != http.StatusOK {
		// Error 500: Internal Server Error
		log.Printf("Unexpected status code for Relations: %d", response.StatusCode)
		return Relation{}
	}

	var relation Relation
	err = json.NewDecoder(response.Body).Decode(&relation)
	if err != nil {
		log.Println("Error decoding Relations:", err)
		return Relation{}
	}

	return relation
}

func formatRelations(relation Relation) string {
	var formattedRelations []string
	for location, dates := range relation.DatesLocations {
		for _, date := range dates {
			formattedRelations = append(formattedRelations, fmt.Sprintf("%s: %s", location, date))
		}
	}
	return strings.Join(formattedRelations, ", ")
}
