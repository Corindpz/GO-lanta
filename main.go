package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"os"
)

type Aventurier struct {
	Nom string `json:"nom"`
	Age string `json:"age"`
	Equipe string `json:"equipe"`
	Profession    string `json:"profession"`
	Motivation   string `json:"motivation"`
	Photo    string `json:"photo"`
}

type Form struct {
	Nom string `json:"nom"`
	Age string `json:"age"`
	Equipe string `json:"equipe"`
	Profession    string `json:"profession"`
	Motivation   string `json:"motivation"`
	Photo    string `json:"photo"`
	ID       int    `json:"id"`
}


func	main(){

	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Printf(fmt.Sprintf("ERREUR => %s", err.Error()))
		return
	}

	css := http.FileServer(http.Dir("./asset/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))


	http.HandleFunc("/creator", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "creator", nil)
	})

	http.HandleFunc("/aventuriers", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "aventuriers", nil)
	})

	http.HandleFunc("/form/treatment", FormSubmission) 
	


	fmt.Println("Serveur démarré sur le port 8080...")
	http.ListenAndServe("localhost:8080", nil)
}

func LoadAventuriers() ([]Form, error) {
	fileData, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}

	var forms []Form

	err = json.Unmarshal(fileData, &forms)
	if err != nil {
		return nil, err
	}

	return forms, nil
}

func FormSubmission(w http.ResponseWriter, r *http.Request) {

	nomFichier := "data.json"

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erreur lors de l'analyse du formulaire", http.StatusInternalServerError)
		return
	}

	dataFile, headerFile, errFile := r.FormFile("file")
	if errFile != nil {

		fmt.Println("erreur avec le fichier....")
	}
	defer dataFile.Close()

	File, errOpen := os.OpenFile(("./asset/uploads/" + headerFile.Filename), os.O_CREATE, 0644)
	if errOpen != nil {
		fmt.Println("Erreur lors de l'ouverture :", err)
		return

	}

	defer File.Close()

	_, errCopy := io.Copy(File, dataFile)
	if errCopy != nil {
		fmt.Println("Erreur lors de la copie :", err)
		return
	
	}

	// Créer une nouvelle instance de Form à partir des données du formulaire
	form := Form{
		Nom:    r.Form.Get("nom"),
		Age:       r.Form.Get("age"),
		Equipe:		r.Form.Get("equipe"),
		Profession: r.Form.Get("profession"),
		Motivation:        r.Form.Get("motivation"),
		Photo:       headerFile.Filename,
	}

	// Ajouter la date actuelle si elle n'est pas fournie dans le formulair

	dataForms, errForms := LoadAventuriers()
	if errForms != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier : %v", errForms), http.StatusInternalServerError)
		return
	}

	// Ajouter la nouvelle forme à la liste
	dataForms = append(dataForms, form)

	dataWrite, errWrite := json.Marshal(dataForms)
	if errWrite != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier : %v", errWrite), http.StatusInternalServerError)
		return
	}

	errWriteFile := os.WriteFile(nomFichier, dataWrite, fs.FileMode(0644))
	if errWriteFile != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'ouverture du fichier : %v", errWriteFile), http.StatusInternalServerError)
		return
	}

	fmt.Println("Ajouté avec succès")
	http.Redirect(w, r, "http://localhost:8080/aventuriers", http.StatusSeeOther)
}
