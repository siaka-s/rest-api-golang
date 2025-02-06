package main

import (
	"log"

	"github.com/siaka-s/rest-api-golang/cmd/api"
)

func main() {

	//initialiser la base de donnés

	// lancer le server d'api

	apiServer := api.NewApiServer(":8080")

	if err := apiServer.Run(); err != nil {
		log.Fatal("erreur de lancement du serveur")
	}

}
