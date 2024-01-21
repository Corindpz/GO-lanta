package controller

import (
	"EpreuveGo/backend"
	"EpreuveGo/templates"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func AccueilPage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "accueil", nil)
}

func CreatePage(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "create", nil)
}

func SubmitCreate(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	ageInt, _ := strconv.Atoi(r.FormValue("age"))
	id := rand.Intn(90000) + 10000
	nouveauAventurier := backend.Aventurier{
		ID:          id,
		Nom:         r.FormValue("nom"),
		Prenom:      r.FormValue("prenom"),
		Age:         ageInt,
		Sexe:        r.FormValue("sexe"),
		Description: r.FormValue("description"),
	}

	var aventuriersData backend.AventuriersData

	file, _ := ioutil.ReadFile("persos.json")

	json.Unmarshal(file, &aventuriersData)

	aventuriersData.Aventuriers = append(aventuriersData.Aventuriers, nouveauAventurier)

	data, _ := json.MarshalIndent(aventuriersData, "", "  ")

	ioutil.WriteFile("persos.json", data, 0644)

	fmt.Println("Aventurier ajouté avec succès")
	http.Redirect(w, r, "/create", http.StatusSeeOther)
	return
}

func PersosPage(w http.ResponseWriter, r *http.Request) {
	var aventuriersData backend.AventuriersData

	jsonData, _ := ioutil.ReadFile("persos.json")

	err := json.Unmarshal(jsonData, &aventuriersData)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}
	templates.Temp.ExecuteTemplate(w, "persos", aventuriersData)
}

func DeletePerso(w http.ResponseWriter, r *http.Request) {
	jsonData, _ := ioutil.ReadFile("persos.json")
	queryParams := r.URL.Query()

	idParam := queryParams.Get("id")

	idToDelete, _ := strconv.Atoi(idParam)

	type AventuriersData struct {
		Aventuriers []backend.Aventurier `json:"aventuriers"`
	}

	var aventurierData AventuriersData
	json.Unmarshal(jsonData, &aventurierData)

	if backend.SupprimerAventurierParID(idToDelete, &aventurierData.Aventuriers) {

		jsonUpdated, _ := json.MarshalIndent(aventurierData, "", "  ")

		ioutil.WriteFile("persos.json", jsonUpdated, os.ModePerm)

		http.Redirect(w, r, "/persos", http.StatusSeeOther)
	} else {
		fmt.Printf("Aventurier avec ID %d non trouvé.\n", idToDelete)
	}

}

func ModifPage(w http.ResponseWriter, r *http.Request) {
	var aventurierData backend.AventuriersData
	queryParams := r.URL.Query()

	idParam := queryParams.Get("id")

	id, _ := strconv.Atoi(idParam)

	jsonData, _ := ioutil.ReadFile("persos.json")

	json.Unmarshal(jsonData, &aventurierData)

	var aventurierRecherche backend.Aventurier
	for _, aventurier := range aventurierData.Aventuriers {
		if aventurier.ID == id {
			aventurierRecherche = aventurier
			break
		}
	}

	if aventurierRecherche.ID == 0 {
		fmt.Println("Aventurier non trouvé avec l'ID:", id)
		http.Error(w, "Aventurier non trouvé", http.StatusNotFound)
		return
	}

	templates.Temp.ExecuteTemplate(w, "modif", aventurierRecherche)
}

func SubmitModif(w http.ResponseWriter, r *http.Request) {
	idStr := r.PostFormValue("id")
	id, _ := strconv.Atoi(idStr)

	data, _ := ioutil.ReadFile("persos.json")

	var aventuriersData backend.AventuriersData
	json.Unmarshal(data, &aventuriersData)

	index := -1
	for i, aventurier := range aventuriersData.Aventuriers {
		if aventurier.ID == id {
			index = i
			break
		}
	}

	aventuriersData.Aventuriers[index].Nom = r.PostFormValue("nom")
	aventuriersData.Aventuriers[index].Prenom = r.PostFormValue("prenom")
	aventuriersData.Aventuriers[index].Age, _ = strconv.Atoi(r.PostFormValue("age"))
	aventuriersData.Aventuriers[index].Sexe = r.PostFormValue("sexe")
	aventuriersData.Aventuriers[index].Description = r.PostFormValue("description")

	nouvellesDonneesJSON, _ := json.MarshalIndent(aventuriersData, "", "  ")

	ioutil.WriteFile("persos.json", nouvellesDonneesJSON, 0644)

	http.Redirect(w, r, "/persos", http.StatusSeeOther)
}
