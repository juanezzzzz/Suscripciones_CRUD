package models

import "database/sql"

type UsuariosSuscripciones struct {
	ID_Usuarios_Suscripciones  int          `json:"id_usuarios_suscripciones"`
	ID_Usuarios                int          `json:"id_usuarios"`
	ID_Planes                  int          `json:"id_planes"`
	ID_Modalidad_Suscripcion   int          `json:"id_modalidad_suscripcion"`
	ID_Estado                  int          `json:"id_estado"`
	Inicio_En                  string       `json:"inicio_en"`
	Fin_En                     sql.NullTime `json:"fin_en"`
	Auto_Renovar               bool         `json:"auto_renovar"`
	Activo                     bool         `json:"activo"`
}
