package main

import (
	"log"
	"net/http"
	"rf/taxa-app/routers"
)

func main() {

	//Inicia as rotas
	router := routers.InitRoutes()

	//Inicia o servidor
	log.Fatal(http.ListenAndServe(":8085", router))

}
