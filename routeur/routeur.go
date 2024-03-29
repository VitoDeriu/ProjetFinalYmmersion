package routeur

import (
	"ProjetFinalYmmersion/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServer() {
	controller.GetDataFromJson() // On initialise le tableau "Aventurier" ici directement pour eviter de le faire dans plusieurs fonctions.

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/ajout", controller.Ajout)
	http.HandleFunc("/liste", controller.Liste)
	http.HandleFunc("/upload", controller.UploadFile)
	http.HandleFunc("/aventurier", controller.FicheAventurier)
	http.HandleFunc("/treatment_to_modif", controller.Treatment_to_modifs)
	http.HandleFunc("/treatment_for_modif", controller.Modification)


	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8080/) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)
}
