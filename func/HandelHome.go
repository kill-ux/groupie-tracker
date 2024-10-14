package groupie

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Concert struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Artist struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
	Concerts     Concert
}

type Page struct {
	Code     int
	MsgError string
	Arts     []Artist
	Art      Artist
	ArtGroups [][]Artist
}

var Data = &Page{}



func RenderPage(page string, res http.ResponseWriter) {
//


	temp, err := template.ParseFiles("templates/" + page + ".html")
	if err != nil {
		fmt.Println(err)
		if page == "error" {

			http.Error(res, "Internal Server Error", 500)
			return
		}
		Error(res, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	err1 := temp.Execute(res, Data)
	if err1 != nil {
		fmt.Println(err1.Error())
		if page == "error" {
			http.Error(res, "Internal Server Error", 500)
			return
		}
		Error(res, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

func Error(res http.ResponseWriter, status int, msgerr string) {
	Data.MsgError = msgerr
	res.WriteHeader(status)
	Data.Code = status
	RenderPage("error", res)
}

func HandelHome(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		Error(res, 405, "Method Not Allowed")
		return
	}
	if req.URL.Path != "/" {
		Error(res, 404, "Oops!! Page Not Found")
		return
	}

	RenderPage("index", res)
}

func HandelArtist(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		Error(res, 405, "Method Not Allowed")
	}
	id := strings.TrimPrefix(req.URL.Path, "/artist/")
	idTemp, err := strconv.Atoi(id)
	if err != nil {
		Error(res, 404, "Oops!! Page Not Found")
		return
	}
	if idTemp < 1 || idTemp > len(Data.Arts) {
		Error(res, 404, "Oops!! Page Not Found")
		return
	}
	Data.Art = Data.Arts[idTemp-1]
	//
	resGet, err := http.Get(Data.Art.Relations)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err.Error())
	}
	defer resGet.Body.Close()
	if resGet.StatusCode != http.StatusOK {
		log.Fatalf("Error: Status code %d ", resGet.StatusCode)
	}
	if err := json.NewDecoder(resGet.Body).Decode(&Data.Art.Concerts); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	//
	RenderPage("artist", res)
}
