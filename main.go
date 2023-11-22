package main

import (
	"github/JCesarBat/web_Postgres/Global"
	"github/JCesarBat/web_Postgres/db"
	"github/JCesarBat/web_Postgres/models"
	"log"
	"net/http"

	"github/JCesarBat/web_Postgres/views"
)

func main() {

	//base datos
	dsn := "host=localhost" +
		" user=postgres" +
		" password=01090679369" +
		" dbname=Web_Postgres " +
		"port=5432 sslmode=disable" +
		" TimeZone=Asia/Shanghai"

	db.DBconnection(dsn)
	db.DB.AutoMigrate(models.Session{})
	db.DB.AutoMigrate(models.Usuario{})

	//servidor
	http.HandleFunc("/Cerrar/", func(w http.ResponseWriter, r *http.Request) {
		cookie := Global.GlobalSessions.BuscarCookie(r)

		session, _ := models.Reed_Code(cookie)
		session.Delete(cookie)
		Global.GlobalSessions.EliminarCookie(w, r)
		http.Redirect(w, r, "/Home/", http.StatusFound)

	})
	http.HandleFunc("/", views.Views_Default)
	http.HandleFunc("/Home/", views.Views_Home)
	http.HandleFunc("/login/", views.Views_Login)
	http.HandleFunc("/Registro/", views.Views_Registro)
	http.HandleFunc("/manejo/", views.Handler_validar_usuario)
	http.HandleFunc("/lOGEARSE/", views.Handler_validar_login)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources"))))

	log.Fatal(http.ListenAndServe(":3000", nil))

}
