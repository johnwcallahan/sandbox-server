package auth

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"

	oidc "github.com/coreos/go-oidc"
)

// Authenticator ...
type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

// NewAuthenticator ...
func NewAuthenticator() (*Authenticator, error) {
	fmt.Println("New Authentication")
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://"+os.Getenv("DOMAIN"))
	if err != nil {
		log.Printf("failed to get provider: %v", err)
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  "http://localhost:" + os.Getenv("PORT") + "/callback",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}
