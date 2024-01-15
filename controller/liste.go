package controller

import (
	"ProjetFinalYmmersion/data"
	"ProjetFinalYmmersion/temps"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)


var Aventurier []data.Aventurier

func Liste(w http.ResponseWriter, r *http.Request) {
	GetDataFromJson()
	temps.Temp.ExecuteTemplate(w, "Liste", Aventurier)
}

func GetDataFromJson() {
	data, err := os.ReadFile("data/data.json") //ouverture et lecture du json
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	json.Unmarshal(data, &Aventurier) //passage en json vers la struct
}
	