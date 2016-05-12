package server

import (
	"net/http"

	"golang.org/x/net/websocket"
	"time"
	"github.com/ventu-io/go-shortid"
)

const (
	sessionCookie string = "SessionId"
	sessionExpiration time.Duration = time.Duration(24*365) * time.Hour;
)

type Session struct {
	Id     string
	Nick   string
	Expire time.Time
}

type HandlerSessionFunc func(http.ResponseWriter, *http.Request, *Session)
type HandlerWsSessionFunc func(*websocket.Conn, *Session)

func NewBasicAuthentication() *basicAuthentication {
	return &basicAuthentication{sessions:make(map[string]*Session)}
}

//basic authentication server
type basicAuthentication struct {
	sessions map[string]*Session
}

//private
func (this *basicAuthentication) getSession(httpRequest *http.Request) (session *Session, created bool) {
	cookie, _ := httpRequest.Cookie(sessionCookie)
	var ok bool

	if cookie != nil {
		if session, ok = this.sessions[cookie.Value]; ok {
			return
		}
	}

	session = this.createAndSetSession()
	created = true
	return
}

func (this *basicAuthentication) createAndSetSession() *Session {
	id, _ := shortid.Generate()
	session := &Session{Id:id,Expire:time.Now().Add(sessionExpiration)}
	this.sessions[session.Id] = session

	return session
}

//public
func (this *basicAuthentication) WrapHandler(handler HandlerSessionFunc) http.HandlerFunc {
	return func(httpResponse http.ResponseWriter, httpRequest *http.Request) {
		session, created := this.getSession(httpRequest)
		if created {
			http.SetCookie(httpResponse, &http.Cookie{Name:sessionCookie, Value:session.Id, Expires:session.Expire, Path:"/"})
		}
		handler(httpResponse, httpRequest, session)
	}
}

func (this *basicAuthentication) WrapWsHandler(handler HandlerWsSessionFunc) websocket.Handler {
	return func(socket *websocket.Conn) {
		session, _ := this.getSession(socket.Request())
		handler(socket, session)
	}
}