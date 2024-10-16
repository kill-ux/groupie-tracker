package groupie

import (
	"fmt"
	"html/template"
	"net/http"
)

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
