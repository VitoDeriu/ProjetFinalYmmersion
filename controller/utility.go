package controller

import (
	"ProjetFinalYmmersion/data"
	"encoding/json"
	"fmt"
	"os"
)


var Aventuriers []data.Aventurier  //déclaration de variable Aventuriers qui correspond a la liste des struct d'aventuriers

// ajout des données de Aventurier dans notre fichier JSON
func SetDataToJson() {
	data, err := json.Marshal(Aventuriers) 					//envoi la struct vers un json
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	os.WriteFile("data/data.json", data, 0644) 				//réécriture du json
}

// récupere les data du json pour les envoyer dans une struct
func GetDataFromJson() {
	data, err := os.ReadFile("data/data.json") 				//ouverture et lecture du json
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	json.Unmarshal(data, &Aventuriers) 						//passage en json vers la struct
}

// fonction pour ajouter un article à notre tableau et potentiellement au json
func AddAventurier(aventurier data.Aventurier, save bool) {
	GetDataFromJson()   									//on ouvre le json et on récup la struct
	Aventuriers = append(Aventuriers, aventurier) 			//on ajoute l'aventurier a une nouvelle 
	if save {												//struct pour la renvoyer dans le json
		SetDataToJson() 									//on réécrit le json
	}
}

// fonction pour supprimer un article de notre tableau et potentiellement du json
func RemoveAventurier(index int, save bool) {
	GetDataFromJson() 										//on récup le json en struct
	var NewAventurier []data.Aventurier
	for _, avent := range Aventuriers { 					//on range les aventurier
		if avent.Id != index { 								//si l'id est différent de celui qu'on rentre en param
			NewAventurier = append(NewAventurier, avent) 	//on l'append au NewAventurier
		} 													//donc si l'id de celui qu'on range == a celui qu'on a rentré en parametre 
	}														//alors il n'est pas append donc il est supprimé
	
	Aventuriers = NewAventurier
	if save {
		SetDataToJson() 									//et on réécrit le json sans celui qu'on a pas append du coup
	}
}

//chekc si l'aventurier existe dans la liste de struct globale
func AventurierExist(id int) bool {
	GetDataFromJson()
	for _, avent := range Aventuriers { 	//on range dans la liste de struct aventurier pour voir si l'id qu'on a rentrer en parametre existe dedans
		if avent.Id == id { 				//si l'id existe ca nous renvoie true 
			return true
		}
	}
	return false 							//si l'id n'existe pas ca renvoie false
}

// recherche par ID
func SearchId(id int) []data.Aventurier {
	var pertinentAventurier []data.Aventurier 								//déclare une variable qui correspond a la struct
	for _, aventurier := range Aventuriers {  								//on va checker chaque articles
		if aventurier.Id == id { 											//on va voir si l'id qu'on a rentrer dans la query de l'url correspond a l'id de l'article
			pertinentAventurier = append(pertinentAventurier, aventurier) 	//si oui on stocke l'article dans pertinentArticle
		}
	}
	return pertinentAventurier //comme ca on renvoit la struct dans le template en haut
}
