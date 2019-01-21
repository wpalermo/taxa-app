package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rf/taxa-app/models"
	"rf/taxa-app/services"
	"rf/taxa-app/util"
)

//Get busca as taxas
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get VERB ", r.URL.Path[1:])

	r.ParseForm()
	for key, value := range r.Form {
		log.Print("%s = %s \n", key, value)
	}
}

//Post calcula a taxa de juros de acordo com a data
func Post(w http.ResponseWriter, r *http.Request) {

	var req models.Request

	json.NewDecoder(r.Body).Decode(&req)

	dataAgendamentoTime := util.StringToDate(req.DataAgendamento, "02/01/2006")
	dataTransferenciaTime := util.StringToDate(req.DataTransferencia, "02/01/2006")

	services.CalcularTaxa(dataTransferenciaTime, dataAgendamentoTime, req.Valor)

}

//Put atualiza um registro
func Put(w http.ResponseWriter, r *http.Request) {

}

//Delete remove um registro
func Delete(w http.ResponseWriter, r *http.Request) {

}
