package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rf/taxa-app/models"
	"rf/taxa-app/services"
	"time"
)

//Get busca as taxas
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get VERB ", r.URL.Path[1:])
}

//Post calcula a taxa de juros de acordo com a data
func Post(w http.ResponseWriter, r *http.Request) {

	var req models.Request

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	services.CalcularTaxa(time.Now(), time.Now(), 2000)
}

//Put atualiza um registro
func Put(w http.ResponseWriter, r *http.Request) {

}

//Delete remove um registro
func Delete(w http.ResponseWriter, r *http.Request) {

}
