package main

import (
	"fmt"
	"net/http"

	Groupie "groupie/func"
)

func main() {
	http.HandleFunc("/", Groupie.HandelHome)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
