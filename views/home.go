package views

import (
	"github/JCesarBat/web_Postgres/Global"
	"github/JCesarBat/web_Postgres/models"
	"net/http"
)

func Views_Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]*models.Usuario{}
	idcookie := Global.GlobalSessions.BuscarCookie(r)

	if idcookie != "" {
		session, err := models.Reed_Code(idcookie)
		if err != nil {
			goto final
		}
		usuario := models.Reed_id(session.UsuarioID)
		data["usuario"] = usuario
	}
final:
	plantilla.ExecuteTemplate(w, "Home.html", data)

}
