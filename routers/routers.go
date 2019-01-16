package routers

import (
	"github.com/gorilla/mux"
)

//InitRoutes inicia as rotas
func InitRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)
	router = SetTaxaRoutes(router)

	return router
}
