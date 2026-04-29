package routes

import (
	"api/controllers"
	"github.com/gorilla/mux"
)

func SetupPlanFeaturesRoutes(r *mux.Router) {
	r.HandleFunc("/plan-features", controllers.GetAllPlanFeatures).Methods("GET")
	r.HandleFunc("/plan-features/{id}", controllers.GetPlanFeaturesByID).Methods("GET")
	r.HandleFunc("/plan-features", controllers.CreatePlanFeatures).Methods("POST")
	r.HandleFunc("/plan-features/{id}", controllers.UpdatePlanFeatures).Methods("PUT")
	r.HandleFunc("/plan-features/{id}", controllers.DeletePlanFeatures).Methods("DELETE")
}
