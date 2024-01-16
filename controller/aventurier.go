package controller
import (
	"ProjetFinalYmmersion/temps"
	"net/http"
)

func Aventureur(w http.ResponseWriter, r *http.Request) {
	temps.Temp.ExecuteTemplate(w, "Aventurier", Aventurier)
}