package controller

import (
	"ProjetFinalYmmersion/data"
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

	fmt.Println("la route marche")

	if r.Method == "POST" {

		r.ParseForm()
		var Aventurier data.Aventurier

		//récupération des données du formulaire avec FormValue uniquement pour les strings
		Aventurier.LastName = r.FormValue("LastName")   //récup du Nom
		Aventurier.FirstName = r.FormValue("FirstName") //récup du Prénom
		Aventurier.Team = r.FormValue("Team")           //récup de l'équipe
		Aventurier.Sexe = r.FormValue("Sexe")           //récup du sexe
		// Aventurier.Img = handler.Filename               //récup le nom de l'image qu'on a créer plus haut.
		Aventurier.Age = GetAgeInt(r)                   //on récup l'age via une fonction qui transforme en int
		// Aventurier.Id = GetAventurierIdSmart()          //on récup l'id via une fonction qui attribut un id en fonction du nombre d'aventurier déjà présent
		// AddAventurier(Aventurier, true)                 //et on lance la fonction pour ajouter l'aventurier avec toutes les infos qu'on a récupérer

		//il faut la fonction is exist



		// idStringed := strconv.Itoa(Aventurier.Id) //transform l'id en string pour la query de la redirection

		//demander aux mentor si c'est pas plus propre de faire comme ca
		//var adress = fmt.Sprintf("aventurier?id=" + strconv.Itoa(Adventurer.Id))

		// http.Redirect(w, r, "/aventurier?id="+idStringed, http.StatusSeeOther) //redirect vers la page de l'aventurier créé grace a la query + StatusSeeOther (code 303) qui évite le nouvel envoi de formulaire si on F5 la page

	}

}
