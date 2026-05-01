package models

import "database/sql"

type Planes struct {
	ID_Planes      int             `json:"id_planes"`
	Nombre         string          `json:"nombre"`
	Precio_Mensual float64         `json:"precio_mensual"`
	Precio_Anual   sql.NullFloat64 `json:"precio_anual"`
	Descripcion    sql.NullString  `json:"descripcion"`
	Activo         bool            `json:"activo"`
}
