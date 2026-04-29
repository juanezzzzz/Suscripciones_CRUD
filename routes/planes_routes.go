package routes

import (
	"api/controllers"
	"github.com/gorilla/mux"
)

func SetupPlanesRoutes(r *mux.Router) {
	r.HandleFunc("/planes", controllers.GetAllPlanes).Methods("GET")
	r.HandleFunc("/planes/{id}", controllers.GetPlanesByID).Methods("GET")
	r.HandleFunc("/planes", controllers.CreatePlanes).Methods("POST")
	r.HandleFunc("/planes/{id}", controllers.UpdatePlanes).Methods("PUT")
	r.HandleFunc("/planes/{id}", controllers.DeletePlanes).Methods("DELETE")
}
