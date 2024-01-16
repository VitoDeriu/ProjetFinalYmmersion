package controller
import (
	"ProjetFinalYmmersion/temps"
	"net/http"
)

func Aventureur(w http.ResponseWriter, r *http.Request) {
	temps.Temp.ExecuteTemplate(w, "Aventurier", Aventurier)
}

//a modifier pour faire correspondre a ce projet

// func ArticleTemp(w http.ResponseWriter, r *http.Request) {

// 	id, err := strconv.Atoi(r.URL.Query().Get("type")) // récup le type dans l'url
// 	if err != nil {
// 		http.Error(w, "ID invalide", http.StatusBadRequest)
// 		return
// 	}
// 	data := SearchId(id) //on stocke l'id correspondant dans data pour l'envoyer dans le template
// 	temps.Temp.ExecuteTemplate(w, "article", data) // renvoi le bon id dans le template
// }

// recherche par ID
// func SearchId(id int) []data.ArtStruct {
// 	var pertinentArticles []data.ArtStruct //déclare une variable qui correspond a la struct
// 	for _, article := range Articles { //on va checker chaque articles
// 		if article.Id == id { //on va voir si l'id qu'on a rentrer dans la query del'url correspond a l'id de l'article
// 			pertinentArticles = append(pertinentArticles, article) //si oui on stocke l'article dans pertinentArticle
// 		}
// 	}
// 	return pertinentArticles //comme ca on renvoit la struct dans le template en haut
// }
