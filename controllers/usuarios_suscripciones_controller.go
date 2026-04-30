package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/config"
	"api/models"

	"github.com/gorilla/mux"
)

func GetAllUsuariosSuscripciones(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`
		SELECT id_usuarios_suscripciones, id_usuarios, id_planes,
		       id_modalidad_suscripcion, id_estado, inicio_en, fin_en, auto_renovar, activo
		FROM usuarios_suscripciones`)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.UsuariosSuscripciones
	for rows.Next() {
		var us models.UsuariosSuscripciones
		rows.Scan(&us.ID_Usuarios_Suscripciones, &us.ID_Usuarios, &us.ID_Planes,
			&us.ID_Modalidad_Suscripcion, &us.ID_Estado, &us.Inicio_En, &us.Fin_En, &us.Auto_Renovar, &us.Activo)
		list = append(list, us)
	}
	ResponseJSON(w, 200, list)
}

func GetUsuariosSuscripcionesByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var us models.UsuariosSuscripciones
	err = config.DB.QueryRow(`
		SELECT id_usuarios_suscripciones, id_usuarios, id_planes,
		       id_modalidad_suscripcion, id_estado, inicio_en, fin_en, auto_renovar, activo
		FROM usuarios_suscripciones WHERE id_usuarios_suscripciones = $1`, id).
		Scan(&us.ID_Usuarios_Suscripciones, &us.ID_Usuarios, &us.ID_Planes,
			&us.ID_Modalidad_Suscripcion, &us.ID_Estado, &us.Inicio_En, &us.Fin_En, &us.Auto_Renovar, &us.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, us)
}

func CreateUsuariosSuscripciones(w http.ResponseWriter, r *http.Request) {
	var us models.UsuariosSuscripciones
	json.NewDecoder(r.Body).Decode(&us)
	err := config.DB.QueryRow(`
		INSERT INTO usuarios_suscripciones (id_usuarios, id_planes, id_modalidad_suscripcion, id_estado, inicio_en, fin_en, auto_renovar, activo)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id_usuarios_suscripciones`,
		us.ID_Usuarios, us.ID_Planes, us.ID_Modalidad_Suscripcion, us.ID_Estado,
		us.Inicio_En, us.Fin_En, us.Auto_Renovar, us.Activo).Scan(&us.ID_Usuarios_Suscripciones)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, us)
}

func UpdateUsuariosSuscripciones(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var us models.UsuariosSuscripciones
	json.NewDecoder(r.Body).Decode(&us)
	_, err := config.DB.Exec(`
		UPDATE usuarios_suscripciones SET id_usuarios=$1, id_planes=$2, id_modalidad_suscripcion=$3,
		    id_estado=$4, inicio_en=$5, fin_en=$6, auto_renovar=$7, activo=$8
		WHERE id_usuarios_suscripciones=$9`,
		us.ID_Usuarios, us.ID_Planes, us.ID_Modalidad_Suscripcion, us.ID_Estado,
		us.Inicio_En, us.Fin_En, us.Auto_Renovar, us.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "UsuariosSuscripciones updated successfully"})
}

func DeleteUsuariosSuscripciones(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM usuarios_suscripciones WHERE id_usuarios_suscripciones = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "UsuariosSuscripciones deleted successfully"})
}
