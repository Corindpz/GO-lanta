package main

import (
	"EpreuveGo/routeur"
	"EpreuveGo/templates"
)

func main() {
	templates.InitTemplate()
	routeur.Initserv()
}
