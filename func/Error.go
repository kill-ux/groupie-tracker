package groupie

import "net/http"

func Error(res http.ResponseWriter, status int, msgerr string) {
	Data.MsgError = msgerr
	res.WriteHeader(status)
	Data.Code = status
	RenderPage("error", res)
}
