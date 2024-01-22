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
	file, handler, err := r.FormFile("Image")		//on récup le fichier depuis le form et on le stock dans file
	if err != nil {                           		//gestion d'erreur
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		return
	}
	defer file.Close() //on ferme le fichier

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
	_, err = io.Copy(outFile, file) 	//on écrit le fichier dans le dossier après l'avoir créer
	if err != nil {
		fmt.Println("Error writing the file")
		fmt.Println(err)
		return
	}
	
	// Gestion de recup des modifications
	var AventurierModif data.Aventurier 	//variable d'attente pour stocker les infos avant des les renvoyer dans le json
	if r.Method == "POST" {
		
		id, err := strconv.Atoi(r.URL.Query().Get("id")) // récup l'id dans l'url depuis la query string
		if err != nil {
			fmt.Println()
			http.Error(w, "ID invalide", http.StatusBadRequest)
			return
		}
		
		r.ParseForm() //on parse les infos récupérer du formulaire grâce a la method POST.

		// ! Attention ! Quand l'attribut enctype dans le form est "multipart/form-data" il faut avoir un input de fichier
		// avec les fonctions au dessus si non le parsing ne renvera rien dans les FormValues plus bas.
		
		AventurierModif.LastName = r.FormValue("LastName")			//recup du nouveau nom
		AventurierModif.FirstName = r.FormValue("FirstName")		//recup du nouveau prénom
		AventurierModif.Team = r.FormValue("Team")					//recup de la nouvelle équipe
		AventurierModif.Sexe = r.FormValue("Sexe")					//recup du nouveau sexe
		AventurierModif.Age = GetAgeInt(r)							//recup du nouvel age
		AventurierModif.Img = handler.Filename						//recup de l'image, obligé d'en remettre une meme si c'est la meme, j'ai pas trouver comment conserver l'anciene dans le formulaire
		AventurierModif.Id = id										//recup du meme id (très important de récup le meme pour pouvoir remplacer)
	}

	//Remplacement de l'aventurier modifié
	for _, A := range Aventuriers {                   	//on regarde dans la liste d'aventurier
		if A.Id == AventurierModif.Id{					//s'il y'a un id qui correspond a celui de l'aventurier modifié
			if AventurierExist(A.Id-1){					//alors on check s'il existe un aveturier dans le json a cet indice
				RemoveAventurier(A.Id, true)			//si oui on le remove
				AddAventurier(AventurierModif, true)	//puis on ajoute l'aventurier modifié
			}
		}
	}

	idStringed := strconv.Itoa(AventurierModif.Id) 
	http.Redirect(w, r, "/aventurier?id="+idStringed, http.StatusSeeOther)
}