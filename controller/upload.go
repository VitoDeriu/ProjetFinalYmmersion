package controller

import (
	"ProjetFinalYmmersion/data"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// Upload une image et la set dans la struct correspondante
func UploadFile(w http.ResponseWriter, r *http.Request) {

	// Récup du fichier img
	file, handler, err := r.FormFile("Image")		//on récup le fichier depuis le form et on le stock dans file
	if err != nil {                           		//gestion d'erreur
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		return
	}
	defer file.Close()		//on ferme le fichier

	// Création de l'img dans le dossier /assets/img
	path, _ := os.Getwd()                                             //on récup le chemin d'acces
	filePath := filepath.Join(path, "/assets/profilePic/", handler.Filename) //set le chemin de l'image vers /asset/img
	outFile, err := os.OpenFile(filePath, os.O_CREATE, 0644)          //je crois que ca attribut le chemin a l'image
	if err != nil {                                                   //gestion d'erreur comme d'hab
		fmt.Println("Error creating the file")
		fmt.Println(err)
		return
	}
	defer outFile.Close()	//on ferme le ficher quand on a fini

	// Ecriture du fichier
	_, err = io.Copy(outFile, file)		//on écrit le fichier dans le dossier après l'avoir créer
	if err != nil {
		fmt.Println("Error writing the file")
		fmt.Println(err)
		return
	}

	// Récupération des infos du formulaire et stockage dans les attributs de la structure
	if r.Method == "POST" {

		r.ParseForm()
		var Aventurier data.Aventurier

		//récupération des données du formulaire avec FormValue uniquement pour les strings
		Aventurier.LastName = r.FormValue("LastName")   	//récup du Nom
		Aventurier.FirstName = r.FormValue("FirstName") 	//récup du Prénom
		Aventurier.Team = r.FormValue("Team")           	//récup de l'équipe
		Aventurier.Sexe = r.FormValue("Sexe")           	//récup du sexe
		Aventurier.Img = handler.Filename               	//récup le nom de l'image qu'on a créer plus haut.
		Aventurier.Age = GetAgeInt(r)                   	//on récup l'age via une fonction qui transforme en int
		Aventurier.Id = GetAventurierIdSmart()          	//on récup l'id via une fonction qui attribut un id en fonction du nombre d'aventurier déjà présent
		AddAventurier(Aventurier, true)                 	//et on lance la fonction pour ajouter l'aventurier avec toutes les infos qu'on a récupérer

		idStringed := strconv.Itoa(Aventurier.Id) //transform l'id en string pour la query de la redirection

		//demander aux mentor si c'est pas plus propre de faire comme ca
		//var adress = fmt.Sprintf("aventurier?id=" + strconv.Itoa(Adventurer.Id))

		http.Redirect(w, r, "/aventurier?id="+idStringed, http.StatusSeeOther) //redirect vers la page de l'aventurier créé grace a la query + StatusSeeOther (code 303) qui évite le nouvel envoi de formulaire si on F5 la page
	}
}

// Récupère l'age de l'aventurier et le transforme en int
func GetAgeInt(r *http.Request) int {
	ageStr := r.FormValue("Age")
	if ageStr == "" {
		fmt.Println("l'age est une chaine vide")
		return 0
	}

	a, err := strconv.ParseInt(r.PostFormValue("Age"), 10, 0) 	//on converti l'age qu'on a récup en int
	if err != nil {
		fmt.Println("Error parsing age", err)
		return 0
	}
	return int(a) 	//et on return un int pour pouvoir l'envoyer dans la struct
}

func GetAventurierIdSmart() int {  //a commenter en détail quand j'ai le temps
	var id int
	var exist bool
	for id = 1; !exist; id++ { 					//on va check chaque aventurier pour voir si leur id existe déjà 
		exist = true							//et dès qu'il existe pas on le crée pour l'envoyer dans le nouvel aventurier
		for _, j := range Aventuriers {
			if id == j.Id {
				exist = false
			}
		}
	}
	id--
	return id
}