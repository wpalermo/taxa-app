package controllers

import (
	"fmt"
	"net/http"
)

//Get busca as taxas
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get VERB ", r.URL.Path[1:])
}

//Post calcula a taxa de juros de acordo com a data
func Post(w http.ResponseWriter, r *http.Request) {

}

//Put atualiza um registro
func Put(w http.ResponseWriter, r *http.Request) {

}

//Delete remove um registro
func Delete(w http.ResponseWriter, r *http.Request) {

}
