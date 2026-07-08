package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}
type LocationsResponse struct {
	Index []Location `json:"index"`
}
type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
type RelationsResponse struct {
	Index []Relation `json:"index"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type DatesResponse struct {
	Index []Date `json:"index"`
}
type FullArtist struct {
	Artist   Artist
	Location Location
	Date     Date
	Relation Relation
}

func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
	
}

func getAllArtists() ([]FullArtist, error) {

	artistsBody, err := fetchData("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	var artists []Artist
	if err := json.Unmarshal(artistsBody, &artists); err != nil {
		fmt.Println("error:", err)
	}

	locationBody, err := fetchData("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	var location LocationsResponse
	if err := json.Unmarshal(locationBody, &location); err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	dateBody, err := fetchData("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	var date DatesResponse
	if err := json.Unmarshal(dateBody, &date); err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	relationBody, err := fetchData("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	var relation RelationsResponse
	if err := json.Unmarshal(relationBody, &relation); err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	var fullArtists []FullArtist
	for _, a := range artists {
		for _, r := range relation.Index {
			if r.Id == a.ID {

				var MatchLocation Location
				for _, l := range location.Index {
					if l.Id == a.ID {
						MatchLocation = l
					}
				}

				var MatchDate Date
				for _, d := range date.Index {
					if d.Id == a.ID {
						MatchDate = d
					}
				}

				fa := FullArtist{
					Artist:   a,
					Location: MatchLocation,
					Date:     MatchDate,
					Relation: r,
				}
				fullArtists = append(fullArtists, fa)
			}
		}
	}
	return fullArtists, nil
}