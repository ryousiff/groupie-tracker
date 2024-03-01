package groupie

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"log"
)
func Filter(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("../Webserver/groupie.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		// Error 500: Internal Server Error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	result := []Artist{}
	query := strings.ToLower(r.FormValue("search"))
	creationFromDate := r.FormValue("rangeFilter")
	creationToDate := r.FormValue("creationDateFilter")
	firstAlbumFromDate := r.FormValue("firstAlbumDateFilter")
	firstAlbumToDate := r.FormValue("otherDateFilter")
	// Assuming checkboxFilter represents the checkboxes for number of members
	members := []string{
		r.FormValue("checkboxFilter1"),
		r.FormValue("checkboxFilter2"),
		r.FormValue("checkboxFilter3"),
		r.FormValue("checkboxFilter4"),
		r.FormValue("checkboxFilter5"),
		r.FormValue("checkboxFilter6"),
		r.FormValue("checkboxFilter7"),
	}
	fromDate, err := strconv.Atoi(creationFromDate)
	toDate, err := strconv.Atoi(creationToDate)
	fromAlbum, err := strconv.Atoi(firstAlbumFromDate)
	toAlbum, err := strconv.Atoi(firstAlbumToDate)
	if err != nil {
		return
	}
	//if statement for creation date
	if creationFromDate > creationToDate {
		creationFromDate, creationToDate = creationToDate, creationFromDate
	} else if creationFromDate == creationToDate {
		fromDate -= 1
		toDate += 1
		creationFromDate = strconv.Itoa(fromDate)
		creationToDate = strconv.Itoa(toDate)
	}
	//if statement for first album date
	if firstAlbumFromDate > firstAlbumToDate {
		firstAlbumFromDate, firstAlbumToDate = firstAlbumToDate, firstAlbumFromDate
	} else if firstAlbumFromDate == firstAlbumToDate {
		fromAlbum -= 1
		toAlbum += 1
		firstAlbumFromDate = strconv.Itoa(fromAlbum)
		firstAlbumToDate = strconv.Itoa(toAlbum)
	}
	for _, v := range allData {
		if handleCreation(creationFromDate, creationToDate, v.CreationDate) && handleAlbum(firstAlbumFromDate, firstAlbumToDate, v.FirstAlbum) && handleMembers(len(v.Members), members) && ContainsLocation(v.Locations, query) {
			result = append(result, v)
		}
	}
	temp.Execute(w, result)
}
func handleCreation(creationFromDate string, creationToDate string, artistCreation int) bool {
	creationFrom, err := strconv.Atoi(creationFromDate)
	if err != nil {
		return false
	}
	creationTo, err := strconv.Atoi(creationToDate)
	if artistCreation >= creationFrom && artistCreation <= creationTo {
		return true
	}
	return false
}
func handleAlbum(firstAlbumFromDate string, firstAlbumToDate string, artistAlbum string) bool {
	//take the year from artistalbum which contains aa date format i want the year only use the time package
	parts := strings.Split(artistAlbum, "-")
	if (firstAlbumFromDate <= parts[2]) && (parts[2] <= firstAlbumToDate) {
		return true
	}
	return false
}
func handleMembers(artistMembers int, memberFilter []string) bool {
	allfalse := true
	for i, checkbox := range memberFilter {
		if checkbox == "on" && artistMembers == (i+1) {
			return true
		}
		if checkbox == "on" {
			allfalse = false
		}
	}
	if allfalse {
		return true
	}
	return false
}

func ContainsLocation(locations []string, query string) bool {
	for _, location := range locations {
		if strings.HasPrefix(strings.ToLower(location), query) {
			return true
		}
	}
	return false
}