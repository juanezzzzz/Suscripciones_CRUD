package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/config"
	"api/models"

	"github.com/gorilla/mux"
)

func GetAllPlanFeatures(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id_plan_features, id_planes, feature, incluida, activo FROM plan_features")
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	var list []models.PlanFeatures
	for rows.Next() {
		var pf models.PlanFeatures
		rows.Scan(&pf.ID_Plan_Features, &pf.ID_Planes, &pf.Feature, &pf.Incluida, &pf.Activo)
		list = append(list, pf)
	}
	ResponseJSON(w, 200, list)
}

func GetPlanFeaturesByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseJSON(w, 400, map[string]string{"error": "Invalid ID"})
		return
	}
	var pf models.PlanFeatures
	err = config.DB.QueryRow(`
		SELECT id_plan_features, id_planes, feature, incluida, activo
		FROM plan_features WHERE id_plan_features = $1`, id).
		Scan(&pf.ID_Plan_Features, &pf.ID_Planes, &pf.Feature, &pf.Incluida, &pf.Activo)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, pf)
}

func CreatePlanFeatures(w http.ResponseWriter, r *http.Request) {
	var pf models.PlanFeatures
	json.NewDecoder(r.Body).Decode(&pf)
	err := config.DB.QueryRow(
		"INSERT INTO plan_features (id_planes, feature, incluida, activo) VALUES ($1,$2,$3,$4) RETURNING id_plan_features",
		pf.ID_Planes, pf.Feature, pf.Incluida, pf.Activo).Scan(&pf.ID_Plan_Features)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 201, pf)
}

func UpdatePlanFeatures(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var pf models.PlanFeatures
	json.NewDecoder(r.Body).Decode(&pf)
	_, err := config.DB.Exec(
		"UPDATE plan_features SET id_planes=$1, feature=$2, incluida=$3, activo=$4 WHERE id_plan_features=$5",
		pf.ID_Planes, pf.Feature, pf.Incluida, pf.Activo, id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "PlanFeatures updated successfully"})
}

func DeletePlanFeatures(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := config.DB.Exec("DELETE FROM plan_features WHERE id_plan_features = $1", id)
	if err != nil {
		ResponseJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	ResponseJSON(w, 200, map[string]string{"message": "PlanFeatures deleted successfully"})
}
