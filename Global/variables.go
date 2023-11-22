package Global

import "github/JCesarBat/web_Postgres/session"

var GlobalSessions = session.NewManager("memory", 3600)
