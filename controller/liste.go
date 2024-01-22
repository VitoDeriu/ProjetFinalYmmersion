package controller

import (
	"ProjetFinalYmmersion/temps"
	"net/http"
	"strconv"
)



func Liste(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		data := Aventuriers
		temps.Temp.ExecuteTemplate(w, "Liste", data)
	} else if r.Method == "POST" {
		_ = r.ParseForm()
		if r.PostForm == nil {
			temps.Temp.ExecuteTemplate(w, "Error", nil)
			return
		}
		AventId, _ := strconv.Atoi(r.PostForm.Get("Suppr"))
		if AventurierExist(int(AventId)) {
			RemoveAventurier(int(AventId), true)
			data := Aventuriers
			temps.Temp.ExecuteTemplate(w, "Liste", data)
		} else {
			temps.Temp.ExecuteTemplate(w, "Error", nil)
			return
		}
	}
}
