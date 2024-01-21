package controller

import (
	"ProjetFinalYmmersion/data"
	"ProjetFinalYmmersion/temps"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

	// Récup du fichier img
	file, handler, err := r.FormFile("Image") //on récup le fichier depuis le form et on le stock dans file
	if err != nil {                           //gestion d'erreur
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		return
	}
	defer file.Close() //on ferme le fichier

	// Créa de l'img dans le dossier /assets/img
	path, _ := os.Getwd()                                             //on récup le chemin d'acces
	filePath := filepath.Join(path, "/assets/img/", handler.Filename) //set le chemin de l'image vers /asset/img
	outFile, err := os.OpenFile(filePath, os.O_CREATE, 0644)          //je crois que ca attribut le chemin a l'image
	if err != nil {                                                   //gestion d'erreur comme d'hab
		fmt.Println("Error creating the file")
		fmt.Println(err)
		return
	}
	defer outFile.Close() //on ferme le ficher quand on a fini

	// Ecriture du fichier
	_, err = io.Copy(outFile, file) //l'a on écrit le fichier dans le dossier après l'avoir créer
	if err != nil {
		fmt.Println("Error writing the file")
		fmt.Println(err)
		return
	}
	
	//gestion de recup des modif
	var AventurierModif data.Aventurier //variable d'attente pour stocker les info avant des les renvoyer dans le json
	
	if r.Method == "POST" {
		
		
		id, err := strconv.Atoi(r.URL.Query().Get("id")) // récup le type dans l'url depuis la query string
		if err != nil {
			fmt.Println()
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}

		fmt.Println("methode post reconue")
		fmt.Println("print de r :", r)
		
		r.ParseForm()
		
		fmt.Println("print de r.form :", r.Form)
		
		AventurierModif.LastName = r.FormValue("LastName")
		AventurierModif.FirstName = r.FormValue("FirstName")
		AventurierModif.Team = r.FormValue("Team")
		AventurierModif.Sexe = r.FormValue("Sexe")
		AventurierModif.Age = GetAgeInt(r)
		AventurierModif.Img = handler.Filename
		AventurierModif.Id = id

		fmt.Println(AventurierModif.LastName)
		fmt.Println(AventurierModif.FirstName)
		fmt.Println(AventurierModif.Team)
		fmt.Println(AventurierModif.Sexe)
		fmt.Println(AventurierModif.Age)
		fmt.Println(AventurierModif.Img)
		fmt.Println(AventurierModif.Id)
	}

		//il faut la fonction is exist

		for _, A := range Aventuriers {                   	//on regarde dans la liste d'aventurier
			if A.Id == AventurierModif.Id{					//si y'a un id qui correspond a celui de l'aventurier modifié


				fmt.Println("index de l'aventurier dans le json :", A.Id)
				fmt.Println("index du nouvel aventurier :", AventurierModif.Id)

				if AventurierExist(A.Id-1){					//alors on check s'il existe un aveturier dans le json a cet indice
					RemoveAventurier(A.Id, true)			//si oui on le remove
					AddAventurier(AventurierModif, true)	//puis on ajoute l'aventurier modifié
				}

				fmt.Println("Last name de l'aventurier supprimé (devrait etre égal a rien ou au nouveau nom ?) : ", A.LastName)
				
				// fmt.Println(A.LastName)
				// fmt.Println(A.FirstName)
				// fmt.Println(A.Team)
			}
		}
		
		for i, P := range Aventuriers {

			fmt.Println("Aventurier",i)
			fmt.Println(P.LastName)
			fmt.Println(P.Id)
			i++

		}







	// fmt.Println("la route marche")
	// var AventurierMod data.Aventurier

	// fmt.Println(r)

	// if r.Method == "POST" {

	// 	body, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		fmt.Println("Erreur lors de la lecture du corps de la requête:", err)
	// 		 // Gère l'erreur ici
	// 		return
	// 	}

	// 	fmt.Println("Corps de la requête:", string(body))

	// 	erre := r.ParseForm()
	// 	if erre != nil {
	// 		fmt.Println("Erreur lors de l'analyse du formulaire:", err)
	// 		// Gère l'erreur ici
	// 		return
	// 	}

	// 	fmt.Println("post pris en compte")
	// 	r.ParseForm()
	// 	fmt.Println("parseform pris en compte")
	// 	fmt.Println("New Name :", r.FormValue("LastName"))
	// 	fmt.Println("New Prénom :", r.FormValue("FirstName"))
	// 	fmt.Println("New Team :", r.FormValue("Team"))
	// 	fmt.Println("New Sexe :", r.FormValue("Sexe"))

	// 	//récupération des données du formulaire avec FormValue uniquement pour les strings
	// 	AventurierMod.LastName = r.FormValue("LastName")   //récup du Nom
	// 	fmt.Println(AventurierMod.LastName)
	// 	AventurierMod.FirstName = r.FormValue("FirstName") //récup du Prénom
	// 	fmt.Println(AventurierMod.FirstName)
	// 	AventurierMod.Team = r.FormValue("Team")           //récup de l'équipe
	// 	fmt.Println(AventurierMod.Team)
	// 	AventurierMod.Sexe = r.FormValue("Sexe")           //récup du sexe
	// 	fmt.Println(AventurierMod.Sexe)
	// 	// Aventurier.Img = handler.Filename               //récup le nom de l'image qu'on a créer plus haut.
	// 	AventurierMod.Age = GetAgeInt(r)                   //on récup l'age via une fonction qui transforme en int
	// 	fmt.Println(AventurierMod.Age)
	// 	// Aventurier.Id = GetAventurierIdSmart()          //on récup l'id via une fonction qui attribut un id en fonction du nombre d'aventurier déjà présent
	// 	// AddAventurier(Aventurier, true)                 //et on lance la fonction pour ajouter l'aventurier avec toutes les infos qu'on a récupérer

	

	// 	}

	// 	// idStringed := strconv.Itoa(Aventurier.Id) //transform l'id en string pour la query de la redirection

	// 	//demander aux mentor si c'est pas plus propre de faire comme ca
	// 	//var adress = fmt.Sprintf("aventurier?id=" + strconv.Itoa(Adventurer.Id))

	// 	// http.Redirect(w, r, "/aventurier?id="+idStringed, http.StatusSeeOther) //redirect vers la page de l'aventurier créé grace a la query + StatusSeeOther (code 303) qui évite le nouvel envoi de formulaire si on F5 la page

	// }

}
