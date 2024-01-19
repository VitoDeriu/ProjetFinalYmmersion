package controller

import (
	"ProjetFinalYmmersion/temps"
	"fmt"
	"net/http"
	"strconv"
)

func Treatment_to_modifs(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id")) // récup le type dans l'url depuis la query string
	if err != nil {
		fmt.Println()
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	data := SearchId(id)
	// http.Redirect(w, r, "/modif", data, http.StatusSeeOther)
	temps.Temp.ExecuteTemplate(w, "Modif", data)

}

func Modification(w http.ResponseWriter, r *http.Request) {

	fmt.Println("ca marche")







}