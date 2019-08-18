package auth

import (
	"log"

	"golang.org/x/oauth2"
)

// PersistingTokenSource is a token source that persist token on refresh
type PersistingTokenSource struct {
	original  oauth2.TokenSource
	persister TokenPersister
}

// NewPersistingTokenSource creates a new persisting token source
func NewPersistingTokenSource(original oauth2.TokenSource, persister TokenPersister) *PersistingTokenSource {
	return &PersistingTokenSource{original: original, persister: persister}
}

// Token returns a token from the original source and persists it
func (source PersistingTokenSource) Token() (*oauth2.Token, error) {
	token, err := source.original.Token()
	if err != nil {
		return nil, err
	}

	err = source.persister.Persist(token)
	if err != nil {
		log.Printf("Error while persisting token: %s\n", err.Error())
	}

	return token, nil
}
