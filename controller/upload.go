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

// Upload une image et la set dans la struct correspondante
func UploadFile(w http.ResponseWriter, r *http.Request) {

	// Récup du fichier img
	file, handler, err := r.FormFile("Image") 	//on récup le fichier depuis le form et on le stock dans file
	if err != nil {                          	//gestion d'erreur
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

	// Récupération des info du formulaire et stockage dans les attributs de la structure
	if r.Method == "POST" {

		r.ParseForm()
		var Aventurier data.Aventurier

		//récupération des données du formulaire avec FormValue uniquement pour les strings
		Aventurier.LastName = r.FormValue("LastName")   //récup du Nom
		Aventurier.FirstName = r.FormValue("FirstName") //récup du Prénom
		Aventurier.Team = r.FormValue("Team")           //récup de l'équipe
		Aventurier.Sexe = r.FormValue("Sexe")           //récup du sexe
		Aventurier.Img = handler.Filename               //récup le nom de l'image qu'on a créer plus haut.
		Aventurier.Age = GetAgeInt(r)                   //on récup l'age via une fonction qui transforme en int
		Aventurier.Id = GetAventurierId()              	//on récup l'id via une fonction qui attribut un id en fonction du nombre d'aventurier déjà présent
		AddAventurier(Aventurier, true)                 //et on lance la fonction pour ajouter l'aventurier avec toutes les infos qu'on a récupérer
	}

	// Renvoie sur la page de l'Aventurier créé
	temps.Temp.ExecuteTemplate(w, "Aventurier", Aventuriers)
}

// récupère l'age de l'aventurier et le transforme en int
func GetAgeInt(r *http.Request) int {
	a, err := strconv.ParseInt(r.FormValue("Age"), 10, 0) 	//on converti l'age qu'on a récup en int
	if err != nil {
		fmt.Println("Error parsing age", err)
	}
	return int(a) 											//et on return un int pour pouvoir l'envoyer dans la struct
}

// récupère l'id de l'aventurier depuis le form et le transforme en int. remplacé par GetAventurierId qui le fait en auto
/*func GetIdInt(r *http.Request) int {
	a, err := strconv.ParseInt(r.FormValue("Id"), 10, 0)
	if err != nil {
		fmt.Println("Error parsing age", err)
	}
	return int(a)
}*/

// créé un id auto qui s'incrémente de 1 a chaque nouvelle création d'aventurier
func GetAventurierId() int {
	id := len(Aventuriers) + 1
	return id
}