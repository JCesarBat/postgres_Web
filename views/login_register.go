package views

import (
	"github/JCesarBat/web_Postgres/Global"
	"net/http"
)

func Views_Login(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	Global.GlobalSessions.CrearCookie(w, r)
	message := r.URL.Query()["error"]

	if message != nil {
		data["mensaje"] = message[0]
	}

	plantilla.ExecuteTemplate(w, "login.html", data)
}

func Views_Registro(w http.ResponseWriter, r *http.Request) {
	Global.GlobalSessions.CrearCookie(w, r)
	data := make(map[string]string)
	message := r.URL.Query()["error"]

	if message != nil {
		data["mensaje"] = message[0]
	}

	plantilla.ExecuteTemplate(w, "registro.html", data)
}
