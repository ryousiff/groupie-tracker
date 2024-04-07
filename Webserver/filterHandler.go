package groupie
import (
	// "log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)
func Filter(w http.ResponseWriter, r *http.Request) {
    // Parse template
    temp, err := template.ParseFiles("../Webserver/groupie.html")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        http.ServeFile(w, r, "../Webserver/error500.html")
        return
    }

    // Extract filter parameters from request
    creationFromDate := r.FormValue("rangeFilter")
    creationToDate := r.FormValue("creationDateFilter")
    firstAlbumFromDate := r.FormValue("firstAlbumDateFilter")
    firstAlbumToDate := r.FormValue("otherDateFilter")

    // Adjust filter parameters if necessary
    fromDate, _ := strconv.Atoi(creationFromDate)
    toDate, _ := strconv.Atoi(creationToDate)
    fromAlbum, _ := strconv.Atoi(firstAlbumFromDate)
    toAlbum, _ := strconv.Atoi(firstAlbumToDate)

    // Swap dates if toDate is greater than fromDate
    if toDate < fromDate {
        creationFromDate, creationToDate = creationToDate, creationFromDate
        fromDate, toDate = toDate, fromDate
    }

    // Swap album dates if toAlbum is greater than fromAlbum
    if toAlbum < fromAlbum {
        firstAlbumFromDate, firstAlbumToDate = firstAlbumToDate, firstAlbumFromDate
        fromAlbum, toAlbum = toAlbum, fromAlbum
    }

    // Update filter parameters with adjusted values
    creationFromDate = strconv.Itoa(fromDate)
    creationToDate = strconv.Itoa(toDate)
    firstAlbumFromDate = strconv.Itoa(fromAlbum)
    firstAlbumToDate = strconv.Itoa(toAlbum)

    // Handle missing or invalid filter parameters
    if creationFromDate == "" || creationToDate == "" || firstAlbumFromDate == "" || firstAlbumToDate == "" {
        w.WriteHeader(http.StatusNotFound)
        http.ServeFile(w, r, "../Webserver/error404.html")
        return
    }

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

    // Perform filtering based on filter parameters
    result := []Artist{}
    query := strings.ToLower(r.FormValue("search"))
    newdata := allData
    for _, v := range newdata {
        if handleCreation(creationFromDate, creationToDate, v.CreationDate) && handleAlbum(firstAlbumFromDate, firstAlbumToDate, v.FirstAlbum) && handleMembers(len(v.Members), members) && ContainsLocation(v.Locations, query) {
            result = append(result, v)
        }
    }

    // Serve error page if no results found
    if len(result) == 0 {
        w.WriteHeader(http.StatusNotFound)
        http.ServeFile(w, r, "../Webserver/error404.html")
        return
    }

    // Execute template with filtered results
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
	// take the year from artistalbum which contains aa date format i want the year only use the time package
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
