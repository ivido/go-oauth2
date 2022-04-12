package server

import (
	"net/http"

	"github.com/ivido/go-oauth2/auth"
	"github.com/ivido/go-oauth2/oauth2"
)

// PreRedirectErrorHandler is used to override the default redirect-on-error bahviour
type PreRedirectErrorHandler func(w http.ResponseWriter, req *auth.AuthorizationRequest, err error) error

// InternalErrorHandler is used to handle internal (unexpected) errors
type InternalErrorHandler func(err error) *oauth2.OAuthError

// ResponseErrorHandler is used to handle response er
type OAuthErrorHandler func(oerr *oauth2.OAuthError)

// Use this function to handle all errors
func (s *Server) handleError(w http.ResponseWriter, req *auth.AuthorizationRequest, err error) error {
	if fn := s.PreRedirectErrorHandler; fn != nil {
		return fn(w, req, err)
	}

	return s.redirectError(w, req, err)
}

// redirectError Handle an error by redirecting
func (s *Server) redirectError(w http.ResponseWriter, req *auth.AuthorizationRequest, err error) error {
	if req == nil {
		return err
	}

	data, _ := s.getErrorData(err)
	return s.redirect(w, req, data)
}

// getErrorData Get error response data
func (s *Server) getErrorData(err error) (map[string]interface{}, int) {
	oerr, ok := err.(oauth2.OAuthError)
	if !ok {
		oerr = oauth2.ErrServerError

		if fn := s.InternalErrorHandler; fn != nil {
			if v := fn(err); v != nil {
				oerr = *v
			}
		}
	}

	if fn := s.OAuthErrorHandler; fn != nil {
		fn(&oerr)
	}

	data := make(map[string]interface{})

	data["error"] = oerr.Error()

	if v := oerr.Code; v != "" {
		data["error_code"] = v
	}

	if v := oerr.Description; v != "" {
		data["error_description"] = v
	}

	var statusCode int
	if v := oerr.StatusCode; v > 0 {
		statusCode = v
	} else {
		statusCode = http.StatusInternalServerError
	}

	return data, statusCode
}
