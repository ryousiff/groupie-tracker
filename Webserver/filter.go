package groupie

import (
	"strings"
	"strconv"
	"net/http"
)

type Filter struct {
    CreationDateRange struct {
        Start int
        End   int
    }
    FirstAlbumDateRange struct {
        Start int
        End   int
    }
    NumMembers int
    Locations  []string
}

func FilterArtists(artists []Artist, filter Filter) []Artist {
    var filteredArtists []Artist

    for _, artist := range artists {
        // Convert artist.FirstAlbum to an integer for comparison
        firstAlbumDate, err := strconv.Atoi(artist.FirstAlbum)
        if err != nil {
            // Handle parsing error if necessary
            continue
        }

        // Apply filters one by one
        if artist.CreationDate >= filter.CreationDateRange.Start &&
            artist.CreationDate <= filter.CreationDateRange.End &&
            firstAlbumDate >= filter.FirstAlbumDateRange.Start &&
            firstAlbumDate <= filter.FirstAlbumDateRange.End &&
            len(artist.Members) == filter.NumMembers &&
            containsAll(strings.Split(artist.Locations, ", "), filter.Locations) {
            filteredArtists = append(filteredArtists, artist)
        }
    }

    return filteredArtists
}




func containsAll(superset, subset []string) bool {
    set := make(map[string]struct{}, len(superset))
    for _, s := range superset {
        set[s] = struct{}{}
    }
    for _, s := range subset {
        if _, ok := set[s]; !ok {
            return false
        }
    }
    return true
}

func parseFilterParameters(r *http.Request) Filter {
    filter := Filter{}

    // Parse creation date range, first album date range, number of members, locations, etc.
    // Example:
    filter.CreationDateRange.Start, _ = strconv.Atoi(r.FormValue("creation_date_start"))
    filter.CreationDateRange.End, _ = strconv.Atoi(r.FormValue("creation_date_end"))

    // Parse other filter parameters accordingly...

    return filter
}