package groupie

// import (
// 	"html/template"
// 	// "log"
// 	"net/http"
// 	"strconv"
// 	"strings"
// )

// func filter(res http.ResponseWriter, req *http.Request) {
// 	temp, er := template.ParseFiles("../Webserver/groupie.html")
// 	if er != nil {
// 		handleError(res, PageData{Err: 500, ErrStr: "Error 500 found"})
// 		return
// 	}
// 	result := []Artist{}
// 	// query := strings.ToLower(req.FormValue("search"))
// 	creationFromDate := req.FormValue("rangeFilter")
// 	creationToDate := req.FormValue("creationDateFilter")
// 	firstAlbumFromDate := req.FormValue("firstAlbumDateFilter")
// 	firstAlbumToDate := req.FormValue("otherDateFilter")
// 	// Assuming checkboxFilter represents the checkboxes for number of members
// 	members := []string{
// 	    req.FormValue("checkboxFilter1"),
// 	    req.FormValue("checkboxFilter2"),
// 	    req.FormValue("checkboxFilter3"),
// 	    req.FormValue("checkboxFilter4"),
// 	    req.FormValue("checkboxFilter5"),
// 	    req.FormValue("checkboxFilter6"),
// 	    req.FormValue("checkboxFilter7"),
// 	}
// 	fromDate, err := strconv.Atoi(creationFromDate)
// 	toDate, err := strconv.Atoi(creationToDate)
// 	fromAlbum, err := strconv.Atoi(firstAlbumFromDate)
// 	toAlbum, err := strconv.Atoi(firstAlbumToDate)
// 	if err != nil {
// 		return
// 	}
// 	//if statement for creation date
// 	if creationFromDate > creationToDate {
// 		creationFromDate, creationToDate = creationToDate, creationFromDate
// 	} else if creationFromDate == creationToDate {
// 		fromDate -= 1
// 		toDate += 1
// 		creationFromDate = strconv.Itoa(fromDate)
// 		creationToDate = strconv.Itoa(toDate)
// 	}
// 	//if statement for first album date
// 	if firstAlbumFromDate > firstAlbumToDate {
// 	    firstAlbumFromDate, firstAlbumToDate = firstAlbumToDate, firstAlbumFromDate
// 	} else if firstAlbumFromDate == firstAlbumToDate {
// 	    fromAlbum -= 1
// 	    toAlbum += 1
// 	    firstAlbumFromDate = strconv.Itoa(fromAlbum)
// 	    firstAlbumToDate = strconv.Itoa(toAlbum)
// 	}
// 	var artists []Artist
// 	for _, v := range artists {
// 		if handleCreation(creationFromDate, creationToDate, v.CreationDate) && handleAlbum(firstAlbumFromDate, firstAlbumToDate, v.FirstAlbum) && handleMembers(len(v.Members),members) {
// 			result = append(result, v)
// 		}
// 	}
// 	temp.Execute(res, result)
// }

// func handleError(w http.ResponseWriter, data PageData) {
// 	tmpl, err := template.ParseFiles("../Webserver/error.html")
// 	if err != nil {
// 		// Render a generic error page if template parsing fails
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	if data.Err == 400 {
// 		w.WriteHeader(http.StatusBadRequest)
// 	} else if data.Err == 404 {
// 		w.WriteHeader(http.StatusNotFound)
// 	} else if data.Err == 500 {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// 	err = tmpl.Execute(w, data)
// }

// func handleCreation(creationFromDate string, creationToDate string, artistCreation int) bool {
// 	creationFrom, err := strconv.Atoi(creationFromDate)
// 	if err != nil {
// 		return false
// 	}
// 	creationTo, err := strconv.Atoi(creationToDate)
// 	if artistCreation >= creationFrom && artistCreation <= creationTo {
// 		return true
// 	}
// 	return false
// }

// func handleAlbum(firstAlbumFromDate string,firstAlbumToDate string, artistAlbum string) bool {
// 	//take the year from artistalbum which contains aa date format i want the year only use the time package
// 	parts := strings.Split(artistAlbum, "-")
// 	if (firstAlbumFromDate <= parts[2]) && (parts[2] <= firstAlbumToDate) {
// 		return true
// 	}
// 	return false
// }


// func handleMembers(artistMembers int, memberFilter []string) bool {
// 	allfalse := true
//     for i, checkbox := range memberFilter {
//         if checkbox == "on" && artistMembers == (i+1) {
//             return true
//         }
// 		if checkbox == "on" {
// allfalse = false
// 		}
//     }
// 	if allfalse {
// 		return true
// 	}
//     return false
// }

// // func containsLocation(locations []string, query string) bool {
// // 	for _, location := range locations {
// // 		if strings.HasPrefix(strings.ToLower(location), query) {
// // 			return true
// // 		}
// // 	}
// // 	// if query == "" {
// // 	// 	return true
// // 	// }
// // 	return false
// // }