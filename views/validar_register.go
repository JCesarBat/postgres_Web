package views

import (
	"encoding/json"
	"github/JCesarBat/web_Postgres/Global"
	"github/JCesarBat/web_Postgres/db"
	"github/JCesarBat/web_Postgres/models"
	"log"
	"net/http"
)

func Handler_validar_login(w http.ResponseWriter, r *http.Request) {

	idCookie := Global.GlobalSessions.BuscarCookie(r)
	var users []models.Usuario
	nombre := r.PostFormValue("nombre")
	password := r.PostFormValue("password")

	_ = db.DB.Find(&users)
	for _, v := range users {

		if nombre == v.Nombre && password == v.Password {
			session := models.Session{
				Code:      idCookie,
				UsuarioID: v.ID,
			}
			models.Session_Save(session)

			http.Redirect(w, r, "/Home/", http.StatusFound)
		}

	}
	http.Redirect(w, r, "/login/?error=CREDENCIALES iNVALIDOS", http.StatusFound)

}

func Handler_validar_usuario(w http.ResponseWriter, r *http.Request) {

	idCookie := Global.GlobalSessions.BuscarCookie(r)
	if idCookie == "" {
		http.Redirect(w, r, "/Registro/", http.StatusFound)
	}

	if idCookie == "no encontro la cookie" {
		log.Fatal("tenemos un problema")

	}

	nombre := r.PostFormValue("nombre")
	password := r.PostFormValue("password")
	password_confir := r.PostFormValue("password_confirmacion")
	edad := r.PostFormValue("edad")

	if nombre == "" || password == "" || edad == "" {
		http.Redirect(w, r, "/Registro/?error=A value is empty", http.StatusFound)
	}
	if len(password) < 8 {
		http.Redirect(w, r, "/Registro/?error=The password must be more large", http.StatusFound)
	}
	if password != password_confir {

		http.Redirect(w, r, "/Registro/?error=password dont match", http.StatusFound)
	}

	bytes := []byte(edad)
	var numero uint8
	json.Unmarshal(bytes, &numero)
	user := models.Usuario{
		Nombre:   nombre,
		Password: password,
		Edad:     numero,
		Activo:   true,
	}

	userR, error := models.Usuario_save(user)
	if error != nil {
		http.Redirect(w, r, "/Registro/?error=Ese nombre de usuario ya existe", http.StatusFound)
	}
	session := models.Session{
		Code:      idCookie,
		UsuarioID: userR.ID,
	}
	_, error2 := models.Session_Save(session)
	if error2 != nil {
		log.Fatal("error al ingresar la session")
	}

	http.Redirect(w, r, "/Home/", http.StatusFound)
}
