package main

import (
	"github/JCesarBat/web_Postgres/db"
	"github/JCesarBat/web_Postgres/models"
	"log"
	"net/http"

	"github/JCesarBat/web_Postgres/views"
)

//var globalSessions *session.Manager

func main() {
	//globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	//base datos
	dsn := "host=localhost" +
		" user=postgres" +
		" password=01090679369" +
		" dbname=Web_Postgres " +
		"port=5432 sslmode=disable" +
		" TimeZone=Asia/Shanghai"

	db.DBconnection(dsn)
	db.DB.AutoMigrate(models.Usuario{})

	//servidor
	http.HandleFunc("/", views.Views_Default)
	http.HandleFunc("/Home/", views.Views_Home)
	http.HandleFunc("/login/", views.Views_Login)
	http.HandleFunc("/Registro/", views.Views_Registro)
	http.HandleFunc("/manejo/", views.Handler_validar_usuario)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources"))))

	log.Fatal(http.ListenAndServe(":3000", nil))

}
