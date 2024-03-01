// //Webserver/fetch.go
// package groupie

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strings"
// )

// func fetchConcertDates(apiURL string) []string {
// 	response, err := http.Get(apiURL)
// 	if err != nil {
// 		log.Println("Error fetching concert dates:", err)
// 		return nil
// 	}
// 	defer response.Body.Close()
// 	log.Println("Status code for concert dates:", response.StatusCode)

// 	if response.StatusCode != http.StatusOK {
// 		// Error 500: Internal Server Error
// 		log.Printf("Unexpected status code for concert dates: %d", response.StatusCode)
// 		return nil
// 	}

// 	var concertDates struct {
// 		Dates []string `json:"dates"`
// 	}

// 	err = json.NewDecoder(response.Body).Decode(&concertDates)
// 	if err != nil {
// 		log.Println("Error decoding concert dates:", err)
// 		return nil
// 	}

// 	return concertDates.Dates
// }

// func fetchLocations(apiURL string) []string {
// 	response, err := http.Get(apiURL)
// 	if err != nil {
// 		log.Println("Error fetching locations:", err)
// 		return nil
// 	}
// 	defer response.Body.Close()
// 	log.Println("Status code for locations:", response.StatusCode)

// 	if response.StatusCode != http.StatusOK {
// 		// Error 500: Internal Server Error
// 		log.Printf("Unexpected status code for locations: %d", response.StatusCode)
// 		return nil
// 	}

// 	var locations1 struct {
// 		Loc []string `json:"locations"`
// 	}

// 	err = json.NewDecoder(response.Body).Decode(&locations1)
// 	if err != nil {
// 		log.Println("Error decoding locations:", err)
// 		return nil
// 	}

// 	return locations1.Loc
// }

// func fetchRelations(apiURL string) Relation {
//     response, err := http.Get(apiURL)
//     if err != nil {
//         log.Println("Error fetching Relations:", err)
//         return Relation{}
//     }
//     defer response.Body.Close()
//     log.Println("Status code for Relations:", response.StatusCode)

//     if response.StatusCode != http.StatusOK {
//         // Error 500: Internal Server Error
//         log.Printf("Unexpected status code for Relations: %d", response.StatusCode)
//         return Relation{}
//     }

//     var relation Relation
//     err = json.NewDecoder(response.Body).Decode(&relation)
//     if err != nil {
//         log.Println("Error decoding Relations:", err)
//         return Relation{}
//     }

//     return relation
// }

// func formatRelations(relation Relation) string {
// 	var formattedRelations []string
// 	for location, dates := range relation.DatesLocations {
// 		for _, date := range dates {
// 			formattedRelations = append(formattedRelations, fmt.Sprintf("%s: %s", location, date))
// 		}
// 	}
// 	return strings.Join(formattedRelations, ", ")
// }

package groupie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)
func GatherDataUp(link string) []Artist {
	data1 := GetData(link)
	var wg sync.WaitGroup
	var mu sync.Mutex
	Artists := []Artist{}
	var numArtists int
	json.Unmarshal(data1, &Artists)
	numArtists = len(Artists)
	wg.Add(numArtists)
	for i := range Artists {
		go func(i int) {
			defer wg.Done()
			r := relation{}
			l := location{}
			d := dates{}
			relData := GetData(Artists[i].Relation)
			locData := GetData(Artists[i].LocationsLink)
			dateData := GetData(Artists[i].ConcertsDatesLink)
			json.Unmarshal(relData, &r)
			json.Unmarshal(locData, &l)
			json.Unmarshal(dateData, &d)
			mu.Lock()
			Artists[i].Concerts = r.Concerts
			Artists[i].Locations = l.Locations
			Artists[i].ConcertDates = d.Dates
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	return Artists
}

func GetData(link string) []byte {
	data1, e1 := http.Get(link)
	if e1 != nil {
		log.Fatal(e1)
		return nil
	}
	data2, e2 := ioutil.ReadAll(data1.Body)
	if e2 != nil {
		log.Fatal(e2)
		return nil
	}
	return data2
}

func returnArtists() []Artist {
	return GatherDataUp("https://groupietrackers.herokuapp.com/api/artists")
}