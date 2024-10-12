package groupie

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

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
}

type Page struct {
	Code     int
	MsgError string
	Arts     []Artist
	Art      Artist
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

	Data.Code = status

	RenderPage("error", res)
}

func HandelHome(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		Error(res, 405, "Method Not Allowed")
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
		Error(res, 404, "OOpes Page not Found!!...")
	}
	if idTemp < 1 || idTemp > len(Data.Arts) {
		fmt.Println("error")
	}
	Data.Art = Data.Arts[idTemp-1]
	RenderPage("artist", res)
}
