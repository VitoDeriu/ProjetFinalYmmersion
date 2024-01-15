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
)

// ajout des données de Aventurier dans notre fichier JSON
func SetDataToJson() {
	data, err := json.Marshal(Aventurier)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	os.WriteFile("data/data.json", data, 0644)
}

// fonction pour ajouter un article à notre tableau et potentiellement au json
func AddAventurier(aventurier data.Aventurier, save bool) {
	GetDataFromJson()
	Aventurier = append(Aventurier, aventurier)
	if save {
		SetDataToJson()
	}
}

// fonction pour supprimer un article de notre tableau et potentiellement du json
func RemoveAventurier(index int, save bool) {
	GetDataFromJson()
	var NewAvent []data.Aventurier
	for _, avent := range Aventurier {
		if avent.Id != index {
			NewAvent = append(NewAvent, avent)
		}
	}
	Aventurier = NewAvent
	if save {
		SetDataToJson()
	}
}

func UploadFile(w http.ResponseWriter, r *http.Request) {

	// Récup du fichier img
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	//Créa de l'img dans le dossier /assets/img
	path, _ := os.Getwd()
	filePath := filepath.Join(path, "/assets/img/", handler.Filename)
	outFile, err := os.OpenFile(filePath, os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error creating the file")
		fmt.Println(err)
		return
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		fmt.Println("Error writing the file")
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)

	if r.Method == "POST" {
		r.ParseForm()
		//récupération des données du formulaire
		var Aventurier data.Aventurier
		Aventurier.LastName = r.FormValue("LastName")
		Aventurier.FirstName = r.FormValue("FirstName")
		Aventurier.Age = r.FormValue("Age")
		Aventurier.Team = r.FormValue("Team")
		Aventurier.Sexe = r.FormValue("Sexe")
		Aventurier.Img = handler.Filename
		Aventurier.Id = GetArticleId()
		AddArticle(Article, true)
	} else if r.Method == "GET" {
		temps.Temp.ExecuteTemplate(w, "ajout", nil)
	}
}


func GetAventurierAge() int{
	r.ParseForm()
	age := int(r.FormValue("Age"))
}