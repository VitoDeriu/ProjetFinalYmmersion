package controller

import (
	"ProjetFinalYmmersion/temps"
	"net/http"
)



func Index(w http.ResponseWriter, r *http.Request) {

	//if pour la gestion de 404 not found pour un mauvais url
	if r.URL.Path != "/"{ 
		temps.Temp.ExecuteTemplate(w, "Error", nil)
		return
	}

	// on crée la liste de struct Aventurier a l'arrivé sur la page index,
	// a voir s'il faudrait pas le mettre dans le routeur a l'initialisation du serv ?
	

	//on exec le template
	temps.Temp.ExecuteTemplate(w, "Index", nil)
}
