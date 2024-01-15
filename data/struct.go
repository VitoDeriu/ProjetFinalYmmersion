package data

type Aventurier struct {
	LastName  string `json:"lastname"`
	FirstName string `json:"firstname"`
	Id        int    `json:"id"`
	Age       int    `json:"age"`
	Team      string `json:"team"`
	Sexe      string `json:"sexe"`
	Img       string `json:"img"`
}
