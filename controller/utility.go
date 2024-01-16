package controller

import (
	"ProjetFinalYmmersion/data"
	"ProjetFinalYmmersion/temps"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// ajout des données de Aventurier dans notre fichier JSON
func SetDataToJson() {
	data, err := json.Marshal(Aventurier) //envoi la struct vers un json
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	os.WriteFile("data/data.json", data, 0644) //réécriture du json
}

// récupere les data du json pour les envoyer dans une struct
func GetDataFromJson() {
	data, err := os.ReadFile("data/data.json") //ouverture et lecture du json
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	json.Unmarshal(data, &Aventurier) //passage en json vers la struct
}

// fonction pour ajouter un article à notre tableau et potentiellement au json
func AddAventurier(aventurier data.Aventurier, save bool) {
	GetDataFromJson()                           //on ouvre le json et on récup la struct
	Aventurier = append(Aventurier, aventurier) //on ajoute l'aventurier a une nouvelle struct pour la renvoyer dans le json
	if save {
		SetDataToJson() //on réécrit le json
	}
}

// fonction pour supprimer un article de notre tableau et potentiellement du json
func RemoveAventurier(index int, save bool) {
	GetDataFromJson() //on récup le json en struct
	var NewAventurier []data.Aventurier
	for _, avent := range Aventurier { //on range les aventurier
		if avent.Id != index { //si l'id est différent de celui qu'on rentre en param
			NewAventurier = append(NewAventurier, avent) //on l'append au NewAventurier
		} //donc si l'id de celui qu'on range == a celui qu'on a rentré en parametre alors il n'est pas append donc il est supprimé
	}
	Aventurier = NewAventurier
	if save {
		SetDataToJson() // et on réécrit le json sans celui qu'on a pas append du coup
	}
}

// Upload une image et la set dans la struct correspondante
func UploadFile(w http.ResponseWriter, r *http.Request) {

	// Récup du fichier img
	file, handler, err := r.FormFile("Image") //on récup le fichier depuis le form et on le stock dans file
	if err != nil {                          //gestion d'erreur
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		return
	}
	defer file.Close() //on ferme le fichier

	//Créa de l'img dans le dossier /assets/img
	path, _ := os.Getwd()                                             //on récup le chemin d'acces
	filePath := filepath.Join(path, "/assets/img/", handler.Filename) //set le chemin de l'image vers /asset/img
	outFile, err := os.OpenFile(filePath, os.O_CREATE, 0644)          //je crois que ca attribut le chemin a l'image
	if err != nil {                                                   //gestion d'erreur comme d'hab
		fmt.Println("Error creating the file")
		fmt.Println(err)
		return
	}

	defer outFile.Close() //on ferme le ficher quand on a fini

	_, err = io.Copy(outFile, file) //l'a on écrit le fichier dans le dossier après l'avoir créer
	if err != nil {
		fmt.Println("Error writing the file")
		fmt.Println(err)
		return
	}
	// fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)

	//on récup les infos du formulaire et on les stocks dans les attribut de la structure
	if r.Method == "POST" {
		r.ParseForm()
		//récupération des données du formulaire
		var Aventurier data.Aventurier
		Aventurier.LastName = r.FormValue("LastName")   //récup du Nom
		Aventurier.FirstName = r.FormValue("FirstName") //récup du Prénom
		Aventurier.Team = r.FormValue("Team")           //récup de l'équipe
		Aventurier.Sexe = r.FormValue("Sexe")           //récup du sexe
		Aventurier.Img = handler.Filename               //récup le nom de l'image qu'on a créer plus haut.
		Aventurier.Age = GetAgeInt(r)                   //on récup l'age via une fonction qui transforme en int
		Aventurier.Id = GetAventurierId()               //on récup l'id via une fonction qui attribut un id en fonction du nombre d'aventurier déjà présent
		AddAventurier(Aventurier, true)                 //et on lance la fonction pour ajouter l'aventurier avec toutes les infos qu'on a récupérer
	}

	temps.Temp.ExecuteTemplate(w, "Aventurier", Aventurier)
}

// récupère l'age de l'aventurier et le transforme en int
func GetAgeInt(r *http.Request) int {
	a, err := strconv.ParseInt(r.FormValue("Age"), 10, 0) //on converti l'age qu'on a récup en int
	if err != nil {
		fmt.Println("Error parsing age", err)
	}
	return int(a) //et on return un int pour pouvoir l'envoyer dans la struct
}

// récupère l'id de l'aventurier et le transforme en int. remplacé par GetAventurierId qui le fait en auto
func GetIdInt(r *http.Request) int {
	a, err := strconv.ParseInt(r.FormValue("Id"), 10, 0)
	if err != nil {
		fmt.Println("Error parsing age", err)
	}
	return int(a)
}

// créé un id auto qui s'incrémente de 1 a chaque nouvelle création d'aventurier
func GetAventurierId() int {
	id := len(Aventurier) + 1
	return id
}
