package controller

import (
	"ProjetFinalYmmersion/temps"
	"net/http"
)

func Ajout(w http.ResponseWriter, r *http.Request) {
	temps.Temp.ExecuteTemplate(w, "Ajout", nil)
}