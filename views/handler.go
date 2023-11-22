package views

import (
	"html/template"
	"net/http"
	"regexp"
)

var plantilla = template.Must(template.ParseFiles("./template/Home.html", "./template/login.html", "./template/registro.html"))
var regex_ruta = regexp.MustCompile("^/(Home)/([a-zA-Z0-9]+)$")

func Views_Default(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/Home/", http.StatusFound)

}
