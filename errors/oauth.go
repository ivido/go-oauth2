package errors

type OAuthError struct {
	StatusCode  int
	Code        string
	Description string
}

func NewOAuthError(statusCode int, code string, description string) *OAuthError {
	return &OAuthError{
		StatusCode:  statusCode,
		Code:        code,
		Description: description,
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
	ErrInvalidRequest                 = NewOAuthError(400, ErrCodeInvalidRequest, "The request is missing a required parameter includes an invalid parameter value includes a parameter more than once or is otherwise malformed")
	ErrUnauthorizedClient             = NewOAuthError(401, ErrCodeUnauthorizedClient, "The client is not authorized to request an authorization code using this method")
	ErrAccessDenied                   = NewOAuthError(403, ErrCodeAccessDenied, "The resource owner or authorization server denied the request")
	ErrUnsupportedResponseType        = NewOAuthError(401, ErrCodeUnsupportedResponseType, "The authorization server does not support obtaining an authorization code using this method")
	ErrInvalidScope                   = NewOAuthError(400, ErrCodeInvalidScope, "The requested scope is invalid unknown or malformed")
	ErrServerError                    = NewOAuthError(500, ErrCodeServerError, "The authorization server encountered an unexpected condition that prevented it from fulfilling the request")
	ErrTemporarilyUnavailable         = NewOAuthError(503, ErrCodeTemporarilyUnavailable, "The authorization server is currently unable to handle the request due to a temporary overloading or maintenance of the server")
	ErrInvalidClient                  = NewOAuthError(401, ErrCodeInvalidClient, "Client authentication failed")
	ErrInvalidGrant                   = NewOAuthError(401, ErrCodeInvalidGrant, "The provided authorization grant (e.g. authorization code resource owner credentials) or refresh token is invalid expired revoked does not match the redirection URI used in the authorization request or was issued to another client")
	ErrUnsupportedGrantType           = NewOAuthError(401, ErrCodeUnsupportedGrantType, "The authorization grant type is not supported by the authorization server")
	ErrCodeChallengeRquired           = NewOAuthError(400, ErrCodeInvalidRequest, "PKCE is required. code_challenge is missing")
	ErrUnsupportedCodeChallengeMethod = NewOAuthError(400, ErrCodeInvalidRequest, "Selected code_challenge_method not supported")
	ErrInvalidCodeChallengeLen        = NewOAuthError(400, ErrCodeInvalidRequest, "Code challenge length must be between 43 and 128 charachters long")
)
