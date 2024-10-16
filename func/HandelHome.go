package groupie

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Concert struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type DataConcertDates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DataLocations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:dates`
}

type Artist struct {
	Id               int      `json:"id"`
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	Members          []string `json:"members"`
	CreationDate     int      `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	Locations        string   `json:"locations"`
	ConcertDates     string   `json:"concertDates"`
	Relations        string   `json:"relations"`
	DataLocations    DataLocations
	DataConcertDates DataConcertDates
	Concerts         Concert
}

type Page struct {
	Code      int
	MsgError  string
	Arts      []Artist
	Art       Artist
	ArtGroups [][]Artist
}

var Data = &Page{}

func RenderPage(page string, res http.ResponseWriter) {
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
		if strings.ContainsRune(req.URL.Path[1:], '/') {
			http.Redirect(res, req, "/notFound", http.StatusFound)
			return
		}
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
		http.Redirect(res, req, "/notFound", http.StatusFound)
		return
	}
	if idTemp < 1 || idTemp > len(Data.Arts) {
		http.Redirect(res, req, "/notFound", http.StatusFound)
		return
	}
	Data.Art = Data.Arts[idTemp-1]
	var wg sync.WaitGroup
	wg.Add(1)
	go Fetch(&wg, Data.Art.Relations, &Data.Art.Concerts)
	wg.Add(1)
	go Fetch(&wg, Data.Art.Locations, &Data.Art.DataLocations)
	wg.Add(1)
	go Fetch(&wg, Data.Art.ConcertDates, &Data.Art.DataConcertDates)
	wg.Wait()
	RenderPage("artist", res)
}

func Fetch(wg *sync.WaitGroup, url string, data any) {
	defer wg.Done()
	resGet, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err.Error())
	}
	defer resGet.Body.Close()
	if err := json.NewDecoder(resGet.Body).Decode(&data); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
}

func CssHandler(res http.ResponseWriter, req *http.Request) {
	filePath := "res/css/" + req.URL.Path[len("/css/"):]
	_, err := os.Stat(filePath)
	if err != nil {
		http.Redirect(res, req, "/notFound", http.StatusFound)
		return
	}
	http.StripPrefix("/css/", http.FileServer(http.Dir("res/css/"))).ServeHTTP(res, req)
}
