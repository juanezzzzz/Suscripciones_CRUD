package routes

import (
	"api/controllers"
	"github.com/gorilla/mux"
)

func SetupUsuariosSuscripcionesRoutes(r *mux.Router) {
	r.HandleFunc("/usuarios-suscripciones", controllers.GetAllUsuariosSuscripciones).Methods("GET")
	r.HandleFunc("/usuarios-suscripciones/{id}", controllers.GetUsuariosSuscripcionesByID).Methods("GET")
	r.HandleFunc("/usuarios-suscripciones", controllers.CreateUsuariosSuscripciones).Methods("POST")
	r.HandleFunc("/usuarios-suscripciones/{id}", controllers.UpdateUsuariosSuscripciones).Methods("PUT")
	r.HandleFunc("/usuarios-suscripciones/{id}", controllers.DeleteUsuariosSuscripciones).Methods("DELETE")
}
