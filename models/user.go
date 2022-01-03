package models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Complete_name string `json:"complete_name"`
	Phone         string `json:"phone"`
	Password      string `json:"password"`
}
