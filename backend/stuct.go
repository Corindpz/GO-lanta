package backend

type Aventurier struct {
	ID          int    `json:"id"`
	Nom         string `json:"nom"`
	Prenom      string `json:"prenom"`
	Age         int    `json:"age"`
	Sexe        string `json:"sexe"`
	Description string `json:"description"`
}

type AventuriersData struct {
	Aventuriers []Aventurier `json:"aventuriers"`
}
