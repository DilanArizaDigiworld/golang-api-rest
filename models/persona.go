package models

import "github.com/jinzhu/gorm"

type Persona struct {
	gorm.Model
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
}
