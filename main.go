package main

import (
	"log"
	"net/http"

	"api/config"
	"api/routes"

	"github.com/gorilla/mux"
)

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	r := mux.NewRouter() // Crear un nuevo router

	r.Use(EnableCORS) // Habilitar CORS

	routes.SetupPlanesRoutes(r)
	routes.SetupPlanFeaturesRoutes(r)
	routes.SetupUsuariosSuscripcionesRoutes(r)

	log.Println("Suscripciones API started on port 8084")
	http.ListenAndServe(":8084", r)
}
