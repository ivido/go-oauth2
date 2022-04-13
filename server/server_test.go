package server_test

import (
	"testing"

	"github.com/ivido/go-oauth2/oauth2"
	"github.com/ivido/go-oauth2/server"
)

func TestNewDefaultServer(t *testing.T) {
	cfg := server.NewDefaultConfig()
	srv := server.NewDefaultServer()

	if srv.Config != cfg {
		t.Errorf("Server config should be %+v, got %+v", cfg, srv.Config)
	}
}

func TestNewServer(t *testing.T) {
	cfg := server.Config{
		TokenType: oauth2.TokenTypeOpaque,
	}

	srv := server.NewServer(cfg)

	if srv.Config != cfg {
		t.Errorf("Server config should be %+v, got %+v", cfg, srv.Config)
	}
}
