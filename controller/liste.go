package controller

import (
	"ProjetFinalYmmersion/temps"
	"net/http"
)



func Liste(w http.ResponseWriter, r *http.Request) {
	GetDataFromJson()
	temps.Temp.ExecuteTemplate(w, "Liste", Aventuriers)
}
