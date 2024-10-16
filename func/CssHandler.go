package groupie

import (
	"net/http"
	"os"
)

func CssHandler(res http.ResponseWriter, req *http.Request) {
	filePath := "res/css/" + req.URL.Path[len("/css/"):]
	_, err := os.Stat(filePath)
	if err != nil {
		http.Redirect(res, req, "/notFound", http.StatusFound)
		return
	}
	http.StripPrefix("/css/", http.FileServer(http.Dir("res/css/"))).ServeHTTP(res, req)
}
