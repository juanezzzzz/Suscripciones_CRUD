package models

type PlanFeatures struct {
	ID_Plan_Features int    `json:"id_plan_features"`
	ID_Planes        int    `json:"id_planes"`
	Feature          string `json:"feature"`
	Incluida         bool   `json:"incluida"`
	Activo           bool   `json:"activo"`
}
