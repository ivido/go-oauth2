package auth

import (
	"net/http"

	"github.com/ivido/go-oauth2/oauth2"
)

type AuthorizationRequest struct {
	ResponseType oauth2.ResponseType
	ClientID     string
	Scopes       []string
	RedirectURI  string
	State        string
	Request      *http.Request
}
