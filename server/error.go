package server

import (
	"net/http"

	"github.com/ivido/go-oauth2/auth"
)

// PreRedirectErrorHandler is used to override the default redirect-on-error bahviour
type PreRedirectErrorHandler func(w http.ResponseWriter, req auth.AuthorizationRequest, err error) error

// InternalErrorHandler is used to handle internal (unexpected) errors
type InternalErrorHandler func()

// ResponseErrorHandler is used to handle response er
type OAuthErrorHandler func()
