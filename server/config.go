package server

import (
	"github.com/valyala/fasthttp"
	"github.com/yaameen/oauth2"
	"time"
)

// Config configuration parameters
type Config struct {
	TokenType             string                // token type
	AllowGetAccessRequest bool                  // to allow GET requests for the token
	AllowedResponseTypes  []oauth2.ResponseType // allow the authorization type
	AllowedGrantTypes     []oauth2.GrantType    // allow the grant type
}

// NewConfig create to configuration instance
func NewConfig() *Config {
	return &Config{
		TokenType:            "Bearer",
		AllowedResponseTypes: []oauth2.ResponseType{oauth2.Code, oauth2.Token},
		AllowedGrantTypes: []oauth2.GrantType{
			oauth2.AuthorizationCode,
			oauth2.PasswordCredentials,
			oauth2.ClientCredentials,
			oauth2.Refreshing,
		},
	}
}

// AuthorizeRequest authorization request
type AuthorizeRequest struct {
	ResponseType   oauth2.ResponseType
	ClientID       string
	Scope          string
	RedirectURI    string
	State          string
	UserID         string
	AccessTokenExp time.Duration
	Context        *fasthttp.RequestCtx
}
