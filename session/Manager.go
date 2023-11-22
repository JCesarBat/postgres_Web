package session

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"sync"
)

type Manager struct {
	CookieName  string
	lock        sync.Mutex
	Maxlifetime int64
}

func NewManager(ckname string, maxlt int64) *Manager {
	return &Manager{
		CookieName:  ckname,
		Maxlifetime: maxlt,
	}

}

func (m *Manager) SessionID() string {
	code := make([]byte, 32)
	io.ReadFull(rand.Reader, code)

	return base64.URLEncoding.EncodeToString(code)

}
func (m *Manager) BuscarCookie(r *http.Request) string {

	cookie, err := r.Cookie(m.CookieName)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (m *Manager) CrearCookie(w http.ResponseWriter, r *http.Request) {
	m.lock.Lock()
	defer m.lock.Unlock()
	_, err := r.Cookie(m.CookieName)
	if err != nil {
		Cookie := &http.Cookie{
			Name:   m.CookieName,
			Value:  m.SessionID(),
			Path:   "/",
			MaxAge: int(m.Maxlifetime),
		}
		http.SetCookie(w, Cookie)
	}
}
func (m *Manager) EliminarCookie(w http.ResponseWriter, r *http.Request) {
	m.lock.Lock()
	defer m.lock.Unlock()
	_, err := r.Cookie(m.CookieName)
	if err == nil {
		Cookie := &http.Cookie{
			Name:   m.CookieName,
			Path:   "/",
			Value:  "cambie",
			MaxAge: -1,
		}
		http.SetCookie(w, Cookie)
	}
}
