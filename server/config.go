package server

import "github.com/ivido/go-oauth2/oauth2"

type Config struct {
	TokenType oauth2.TokenType
}

func NewDefaultConfig() Config {
	return Config{
		TokenType: oauth2.TokenTypeJwt,
	}
}
