package main

import (
	"log"
	"net/http"
	"rf/taxa-app/routers"
)

func main() {

	log.Print("Iniciando taxa-app")

	//Inicia as rotas
	router := routers.InitRoutes()

	log.Print("Test debug")

	//Inicia o servidor
	log.Fatal(http.ListenAndServe(":8086", router))

}
