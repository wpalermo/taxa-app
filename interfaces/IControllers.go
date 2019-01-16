package controllers

import "net/http"

//ControllerInterface contem os metodos obrigatorios para os controllers
type ControllerInterface interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
