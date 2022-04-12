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
	ResponseTypeCode  ResponseType = "code"
	ResponseTypeToken ResponseType = "token"
)

// GrantType is the OAuth2 authorization grant type
type GrantType string

const (
	GrantTypeAuthorizationCode   GrantType = "authorization_code"
	GrantTypePasswordCredentials GrantType = "password"
	GrantTypeClientCredentials   GrantType = "client_credentials"
	GrantTypeRefreshing          GrantType = "refresh_token"
	GrantTypeImplicit            GrantType = "__implicit"
)

func (gt GrantType) String() string {
	if gt == GrantTypeAuthorizationCode ||
		gt == GrantTypePasswordCredentials ||
		gt == GrantTypeClientCredentials ||
		gt == GrantTypeRefreshing {
		return string(gt)
	}
	return ""
}
