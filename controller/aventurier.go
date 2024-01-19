package controller

import (
	"ProjetFinalYmmersion/temps"
	"net/http"
	"strconv"
)

func FicheAventurier(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id")) // récup le type dans l'url depuis la query string
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	data := SearchId(id)                              	//on stocke l'id correspondant dans data pour l'envoyer dans le template
	temps.Temp.ExecuteTemplate(w, "Aventurier", data) 	// renvoi le bon id dans le template et execution du template
}

