package models

import (
	"github/JCesarBat/web_Postgres/db"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model

	Nombre   string `gorm:"unique;not null"`
	Password string
	Edad     uint8
	Activo   bool
}

func (this *Usuario) Usuario_save() error {
	result := db.DB.Create(&this)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
