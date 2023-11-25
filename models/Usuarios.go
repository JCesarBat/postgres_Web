package models

import (
	"github/JCesarBat/web_Postgres/db"
	"gorm.io/gorm"
	"log"
)

type Usuario struct {
	gorm.Model

	Nombre   string `gorm:"unique;not null"`
	Password string
	Edad     uint8
	Activo   bool
	session  Session
}

func Usuario_save(user Usuario) (*Usuario, error) {

	result := db.DB.Create(&user)
	if result.Error != nil {
		return &user, result.Error
	}

	return &user, nil
}

func Reed_id(id uint) *Usuario {
	var user Usuario
	result := db.DB.Where("id = ?", id).First(&user)

	if result.Error != nil {
		println(" NO SE ENCONTRO ALGUN USUARIO  CON ESE ID ....")

	}

	return &user
}
func (this *Usuario) Delete() {
	var user Usuario
	result := db.DB.Where(Reed_id(this.ID)).First(&user)
	db.DB.Delete(&user)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Println("se elimino correctamente")

}
