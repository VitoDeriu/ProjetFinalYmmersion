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
	
	//on exec le template
	temps.Temp.ExecuteTemplate(w, "Index", nil)
}
