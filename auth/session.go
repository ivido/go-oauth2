package auth

type AuthSession struct {
	AuthorizationRequest AuthorizationRequest
	SessionState         SessionState
}

type SessionState string

const (
	SessionStateStarted       SessionState = "STARTED"
	SessionStateAuthenticated SessionState = "AUTHENTICATED"
	SessionStateAuthorized    SessionState = "AUTHORIZED"
	SessionStateCodeIssued    SessionState = "CODE_ISSUED"
)
