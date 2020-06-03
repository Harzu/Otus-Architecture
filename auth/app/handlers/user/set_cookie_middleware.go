package user

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

type SessionCookieResponder struct {
	cookie    *http.Cookie
	responder middleware.Responder
}

func NewSessionCookieResponder(responder middleware.Responder, session *http.Cookie) *SessionCookieResponder {
	return &SessionCookieResponder{
		cookie:    session,
		responder: responder,
	}
}

func (this *SessionCookieResponder) WriteResponse(rw http.ResponseWriter, p runtime.Producer) {
	http.SetCookie(rw, this.cookie)
	this.responder.WriteResponse(rw, p)
}
