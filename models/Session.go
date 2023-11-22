package models

import (
	"errors"
	"github/JCesarBat/web_Postgres/db"
	"gorm.io/gorm"
	"log"
)

type Session struct {
	gorm.Model
	Code      string
	UsuarioID uint
}

func Session_Save(session Session) (*Session, error) {
	result := db.DB.Create(&session)
	if result.Error != nil {
		log.Fatal("ocurrio un error al ingresar datos")
	}

	return &session, nil
}

func Reed_Code(Code string) (*Session, error) {
	var session Session
	result := db.DB.Where("code = ?", Code).First(&session)

	if result.Error != nil {
		return nil, errors.New("no existe la session con esa cookie")
	}

	return &session, nil
}

func (this *Session) Delete(code string) {
	var session Session
	result := db.DB.Where("code=?", code).First(&session)
	db.DB.Delete(&session)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Println("se elimino correctamente")

}
