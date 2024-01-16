package controller

import (
	"ProjetFinalYmmersion/data"
	"ProjetFinalYmmersion/temps"
	"net/http"
)

var Aventurier []data.Aventurier

func Liste(w http.ResponseWriter, r *http.Request) {
	GetDataFromJson()
	temps.Temp.ExecuteTemplate(w, "Liste", Aventurier)
}
