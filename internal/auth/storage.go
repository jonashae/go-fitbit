package auth

import (
	"golang.org/x/oauth2"
)

// TokenPersister is an interface for persisting tokens
type TokenPersister interface {
	Persist(*oauth2.Token) error
	Wipe() error
}

// TokenLoader is an interface for loading tokens
type TokenLoader interface {
	Load() (*oauth2.Token, error)
}

// TokenStorage is a union of TokenPersister and TokenLoader
type TokenStorage interface {
	TokenPersister
	TokenLoader
}
