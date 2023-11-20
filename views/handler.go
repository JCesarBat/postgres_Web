package views

import (
	"encoding/json"
	"github/JCesarBat/web_Postgres/models"
	"html/template"
	"net/http"
	"regexp"
)

var plantilla = template.Must(template.ParseFiles("./template/Home.html", "./template/login.html", "./template/registro.html"))
var regex_ruta = regexp.MustCompile("^/(Home)/([a-zA-Z0-9]+)$")

func Views_Default(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/Home/", http.StatusFound)

}

func Views_Home(w http.ResponseWriter, r *http.Request) {

	plantilla.ExecuteTemplate(w, "Home.html", nil)

}
func Views_Login(w http.ResponseWriter, r *http.Request) {

	plantilla.ExecuteTemplate(w, "login.html", nil)
}

func Views_Registro(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	message := r.URL.Query()["error"]

	if message != nil {
		data["mensaje"] = message[0]
	}

	plantilla.ExecuteTemplate(w, "registro.html", data)
}

func Handler_validar_usuario(w http.ResponseWriter, r *http.Request) {

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
	error := user.Usuario_save()
	if error != nil {
		http.Redirect(w, r, "/Registro/?error=Ese nombre de usuario ya existe", http.StatusFound)
	}
	http.Redirect(w, r, "/Home/", http.StatusFound)
}
