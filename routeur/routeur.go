package routeur

import (
	"EpreuveGo/controller"
	"fmt"
	"log"
	"net/http"
)

func Initserv() {

	css := http.FileServer(http.Dir("./assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	http.HandleFunc("/accueil", controller.AccueilPage)
	http.HandleFunc("/create", controller.CreatePage)
	http.HandleFunc("/submit_create", controller.SubmitCreate)
	http.HandleFunc("/persos", controller.PersosPage)
	http.HandleFunc("/delete", controller.DeletePerso)
	http.HandleFunc("/modif", controller.ModifPage)
	http.HandleFunc("/submit_modif", controller.SubmitModif)

	log.Println("[✅ ] Serveur lancé !")
	fmt.Println("[🌐] http://localhost:8082/accueil")
	http.ListenAndServe(":8082", nil)
	log.Fatal()
	
}
