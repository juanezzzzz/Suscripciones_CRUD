package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/config"
	"api/models"

	"github.com/gorilla/mux"
)

func GetAllPlanes(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_planes, nombre, precio_mensual, precio_anual, descripcion, activo FROM planes")
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.Planes
	for rows.Next() {
		var p models.Planes
		rows.Scan(&p.ID_Planes, &p.Nombre, &p.Precio_Mensual, &p.Precio_Anual, &p.Descripcion, &p.Activo)
		list = append(list, p)
	}
	ResponseJSON(w, 200, list)
}

func GetPlanesByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var p models.Planes
	err = config.DB.QueryRow(`
		SELECT id_planes, nombre, precio_mensual, precio_anual, descripcion, activo
		FROM planes WHERE id_planes = $1`, id).
		Scan(&p.ID_Planes, &p.Nombre, &p.Precio_Mensual, &p.Precio_Anual, &p.Descripcion, &p.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, p)
}

func CreatePlanes(w http.ResponseWriter, r *http.Request) {
	var p models.Planes
	json.NewDecoder(r.Body).Decode(&p)
	err := config.DB.QueryRow(
		"INSERT INTO planes (nombre, precio_mensual, precio_anual, descripcion, activo) VALUES ($1,$2,$3,$4,$5) RETURNING id_planes",
		p.Nombre, p.Precio_Mensual, p.Precio_Anual, p.Descripcion, p.Activo).Scan(&p.ID_Planes)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, p)
}

func UpdatePlanes(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var p models.Planes
	json.NewDecoder(r.Body).Decode(&p)
	_, err := config.DB.Exec(
		"UPDATE planes SET nombre=$1, precio_mensual=$2, precio_anual=$3, descripcion=$4, activo=$5 WHERE id_planes=$6",
		p.Nombre, p.Precio_Mensual, p.Precio_Anual, p.Descripcion, p.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Planes updated successfully"})
}

func DeletePlanes(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM planes WHERE id_planes = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "Planes deleted successfully"})
}
