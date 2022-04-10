package server

import "github.com/ivido/go-oauth2/auth"

// RedirectURIBuilder creates the redirect URI for the authorization session
type RedirectURIBuilder func(req auth.AuthorizationRequest, data map[string]interface{}) (string, error)

func DefaultRedirectURIBuilder(req auth.AuthorizationRequest, data map[string]interface{}) (string, error) {
	return "", nil
}
