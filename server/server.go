package server

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ivido/go-oauth2/auth"
	"github.com/ivido/go-oauth2/oauth2"
)

type Server struct {
	Config                        *Config
	AuthorizationRequestValidator auth.AuthorizationRequestValidator
	PreRedirectErrorHandler       PreRedirectErrorHandler
	InternalErrorHandler          InternalErrorHandler
	OAuthErrorHandler             OAuthErrorHandler
}

func NewDefaultServer() *Server {
	return NewServer(&Config{
		TokenType: oauth2.TokenTypeJwt,
	})
}

func NewServer(cfg *Config) *Server {
	return &Server{
		Config:                        cfg,
		AuthorizationRequestValidator: auth.DefaultAuthorizationRequestValidator,
	}
}

func (s *Server) HandleAuthorizationRequest(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	req, err := s.AuthorizationRequestValidator(ctx)
	if err != nil {
		return s.handleError(w, req, err)
	}

	return nil
}

func (s *Server) redirect(w http.ResponseWriter, req *auth.AuthorizationRequest, data map[string]interface{}) error {
	uri, err := s.getRedirectURI(req, data)
	if err != nil {
		return err
	}

	w.Header().Set("Location", uri)
	w.WriteHeader(302)
	return nil
}

func (s *Server) getRedirectURI(req *auth.AuthorizationRequest, data map[string]interface{}) (string, error) {
	u, err := url.Parse(req.RedirectURI)
	if err != nil {
		return "", err
	}

	q := u.Query()
	if req.State != "" {
		q.Set("state", req.State)
	}

	for k, v := range data {
		q.Set(k, fmt.Sprint(v))
	}

	switch req.ResponseType {
	case oauth2.Code:
		u.RawQuery = q.Encode()
	case oauth2.Token:
		u.RawQuery = ""
		fragment, err := url.QueryUnescape(q.Encode())
		if err != nil {
			return "", err
		}
		u.Fragment = fragment
	}

	return u.String(), nil
}
