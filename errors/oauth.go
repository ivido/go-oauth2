package errors

type OAuthError struct {
	Code        string
	Description string
	StatusCode  int
}

func NewOAuthError(code string, description string, statusCode int) *OAuthError {
	return &OAuthError{
		Code:        code,
		Description: description,
		StatusCode:  statusCode,
	}
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

// OAuth errors
var (
	ErrInvalidRequest                 = NewOAuthError(ErrCodeInvalidRequest, "The request is missing a required parameter includes an invalid parameter value includes a parameter more than once or is otherwise malformed", 400)
	ErrUnauthorizedClient             = NewOAuthError(ErrCodeUnauthorizedClient, "The client is not authorized to request an authorization code using this method", 401)
	ErrAccessDenied                   = NewOAuthError(ErrCodeAccessDenied, "The resource owner or authorization server denied the request", 403)
	ErrUnsupportedResponseType        = NewOAuthError(ErrCodeUnsupportedResponseType, "The authorization server does not support obtaining an authorization code using this method", 401)
	ErrInvalidScope                   = NewOAuthError(ErrCodeInvalidScope, "The requested scope is invalid unknown or malformed", 400)
	ErrServerError                    = NewOAuthError(ErrCodeServerError, "The authorization server encountered an unexpected condition that prevented it from fulfilling the request", 500)
	ErrTemporarilyUnavailable         = NewOAuthError(ErrCodeTemporarilyUnavailable, "The authorization server is currently unable to handle the request due to a temporary overloading or maintenance of the server", 503)
	ErrInvalidClient                  = NewOAuthError(ErrCodeInvalidClient, "Client authentication failed", 401)
	ErrInvalidGrant                   = NewOAuthError(ErrCodeInvalidGrant, "The provided authorization grant (e.g. authorization code resource owner credentials) or refresh token is invalid expired revoked does not match the redirection URI used in the authorization request or was issued to another client", 401)
	ErrUnsupportedGrantType           = NewOAuthError(ErrCodeUnsupportedGrantType, "The authorization grant type is not supported by the authorization server", 401)
	ErrCodeChallengeRquired           = NewOAuthError(ErrCodeInvalidRequest, "PKCE is required. code_challenge is missing", 400)
	ErrUnsupportedCodeChallengeMethod = NewOAuthError(ErrCodeInvalidRequest, "Selected code_challenge_method not supported", 400)
	ErrInvalidCodeChallengeLen        = NewOAuthError(ErrCodeInvalidRequest, "Code challenge length must be between 43 and 128 charachters long", 400)
)
