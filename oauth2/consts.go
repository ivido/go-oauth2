package oauth2

// TokenType is the type of access token to be issued by the server
type TokenType string

const (
	TokenTypeJwt    TokenType = "jwt"
	TokenTypeOpaque TokenType = "opaque"
)

// ResponseType is the OAuth2 response type
type ResponseType string

const (
	Code  ResponseType = "code"
	Token ResponseType = "token"
)

// GrantType is the OAuth2 authorization grant type
type GrantType string

const (
	AuthorizationCode   GrantType = "authorization_code"
	PasswordCredentials GrantType = "password"
	ClientCredentials   GrantType = "client_credentials"
	Refreshing          GrantType = "refresh_token"
	Implicit            GrantType = "__implicit"
)

func (gt GrantType) String() string {
	if gt == AuthorizationCode ||
		gt == PasswordCredentials ||
		gt == ClientCredentials ||
		gt == Refreshing {
		return string(gt)
	}
	return ""
}
