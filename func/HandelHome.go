package groupie

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Code     int
	MsgError string
}

var Data *Page

func RenderPage(page string, res http.ResponseWriter) {
	temp, err := template.ParseFiles("templates/" + page + ".html")
	if err != nil {
		if page == "error" {
			http.Error(res, "Internal Server Error", 500)
			return
		}
		Error(res, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	fmt.Println()
	err = temp.Execute(res, Data)
	if err != nil {
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
	res.WriteHeader(status)
	RenderPage("error", res)
}

func HandelHome(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		Error(res, 405, "Method Not Allowed")
	}
	RenderPage("index", res)
}
