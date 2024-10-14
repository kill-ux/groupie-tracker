package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	Groupie "groupie/func"
)

func init() {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Error: Status code %d ", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&Groupie.Data.Arts); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	for i := 0; i < len(Groupie.Data.Arts); i += 10 {
		end := i + 10
		if end > len(Groupie.Data.Arts) {
			end = len(Groupie.Data.Arts)
		}
		Groupie.Data.ArtGroups = append(Groupie.Data.ArtGroups, Groupie.Data.Arts[i:end])
	}
}

func main() {
	if len(Groupie.Data.Arts) == 0 {
		log.Fatalf("Error fetching data")
	}
	http.HandleFunc("/", Groupie.HandelHome)
	http.HandleFunc("/artist/", Groupie.HandelArtist)
	http.HandleFunc("/css/", Groupie.CssHandler)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
