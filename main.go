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

	if err := json.NewDecoder(res.Body).Decode(&Groupie.Data.Art); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	fmt.Println(Groupie.Data.Art)
}

func main() {
	http.HandleFunc("/", Groupie.HandelHome)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
