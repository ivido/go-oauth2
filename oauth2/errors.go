package oauth2

import "fmt"

type OAuthError struct {
	StatusCode  int
	Code        string
	Description string
}

func NewOAuthError(statusCode int, code string, description string) OAuthError {
	return OAuthError{
		StatusCode:  statusCode,
		Code:        code,
		Description: description,
	}
}

func (e OAuthError) Error() string {
	return fmt.Sprintf("Error (%s) %s", e.Code, e.Description)
}

// OAuth error codes
var (
	ErrCodeAccessDenied            = "access_denied"
	ErrCodeInvalidClient           = "invalid_client"
	ErrCodeInvalidGrant            = "invalid_grant"
	ErrCodeInvalidRequest          = "invalid_request"
	ErrCodeInvalidScope            = "invalid_scope"
	ErrCodeServerError             = "server_error"
	ErrCodeTemporarilyUnavailable  = "temporarily_unavailable"
	ErrCodeUnauthorizedClient      = "unauthorized_client"
	ErrCodeUnsupportedGrantType    = "unsupported_grant_type"
	ErrCodeUnsupportedResponseType = "unsupported_response_type"
)

// OAuth error descriptions
var (
	ErrMsgInvalidRequest          = "The request is missing a required parameter includes an invalid parameter value includes a parameter more than once or is otherwise malformed"
	ErrMsgUnauthorizedClient      = "The client is not authorized to request an authorization code using this method"
	ErrMsgAccessDenied            = "The resource owner or authorization server denied the request"
	ErrMsgUnsupportedResponseType = "The authorization server does not support obtaining an authorization code using this method"
	ErrMsgInvalidScope            = "The requested scope is invalid unknown or malformed"
	ErrMsgServerError             = "The authorization server encountered an unexpected condition that prevented it from fulfilling the request"
	ErrMsgTemporarilyUnavailable  = "The authorization server is currently unable to handle the request due to a temporary overloading or maintenance of the server"
	ErrMsgInvalidClient           = "Client authentication failed"
	ErrMsgInvalidGrant            = "The provided authorization grant (e.g. authorization code resource owner credentials) or refresh token is invalid expired revoked does not match the redirection URI used in the authorization request or was issued to another client"
	ErrMsgUnsupportedGrantType    = "The authorization grant type is not supported by the authorization server"
)

// OAuth errors
var (
	ErrInvalidRequest          = NewOAuthError(400, ErrCodeInvalidRequest, ErrMsgInvalidRequest)
	ErrUnauthorizedClient      = NewOAuthError(401, ErrCodeUnauthorizedClient, ErrMsgUnauthorizedClient)
	ErrAccessDenied            = NewOAuthError(403, ErrCodeAccessDenied, ErrMsgAccessDenied)
	ErrUnsupportedResponseType = NewOAuthError(401, ErrCodeUnsupportedResponseType, ErrMsgUnsupportedResponseType)
	ErrInvalidScope            = NewOAuthError(400, ErrCodeInvalidScope, ErrMsgInvalidScope)
	ErrServerError             = NewOAuthError(500, ErrCodeServerError, ErrMsgServerError)
	ErrTemporarilyUnavailable  = NewOAuthError(503, ErrCodeTemporarilyUnavailable, ErrMsgTemporarilyUnavailable)
	ErrInvalidClient           = NewOAuthError(401, ErrCodeInvalidClient, ErrMsgInvalidClient)
	ErrInvalidGrant            = NewOAuthError(401, ErrCodeInvalidGrant, ErrMsgInvalidGrant)
	ErrUnsupportedGrantType    = NewOAuthError(401, ErrCodeUnsupportedGrantType, ErrMsgUnsupportedGrantType)
)
