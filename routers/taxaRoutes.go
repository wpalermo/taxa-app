package routers

import (
	"rf/taxa-app/controllers"

	"github.com/gorilla/mux"
)

//SetTaxaRoutes configura as rotas da taxa
func SetTaxaRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/taxa", controllers.Get).Methods("GET")

	return router

}
