package auth_test

import (
	"errors"
	"testing"

	"github.com/byonchev/go-fitbit/internal/auth"
	"github.com/stretchr/testify/assert"

	"golang.org/x/oauth2"
)

type MockTokenPersister struct {
	err    error
	called bool
}

func (p *MockTokenPersister) Persist(*oauth2.Token) error {
	p.called = true
	return p.err
}

func TestPersistTokenSource(t *testing.T) {
	token := &oauth2.Token{}
	originalSource := oauth2.StaticTokenSource(token)

	var tests = []struct {
		description    string
		persisterError error
	}{
		{"no error", nil},
		{"persister error", errors.New("")},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			persister := &MockTokenPersister{err: test.persisterError, called: false}
			persistingSource := auth.NewPersistingTokenSource(originalSource, persister)

			result, err := persistingSource.Token()

			assert.Nil(t, err)
			assert.Equal(t, token, result)
			assert.True(t, persister.called)
		})
	}
}
