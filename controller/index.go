package controller

import (
	"ProjetFinalYmmersion/data"
	"ProjetFinalYmmersion/temps"
	"net/http"
)

var Aventuriers []data.Aventurier

func Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		temps.Temp.ExecuteTemplate(w, "error404", nil)
		return
	}

	GetDataFromJson()
	temps.Temp.ExecuteTemplate(w, "Index", nil)
}
