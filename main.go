package main

import (
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
}

func main() {
	http.HandleFunc("/", Groupie.HandelHome)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
