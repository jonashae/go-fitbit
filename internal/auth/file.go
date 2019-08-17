package auth

import (
	"encoding/json"
	"errors"
	"os"

	"golang.org/x/oauth2"
)

// FileTokenStorage stores the access and refresh tokens in a file
type FileTokenStorage struct {
	filepath string
}

// NewFileTokenStorage creates a new file token store
func NewFileTokenStorage(filepath string) FileTokenStorage {
	return FileTokenStorage{filepath: filepath}
}

// Load reads a previously saved token from file
func (store FileTokenStorage) Load() (*oauth2.Token, error) {
	if !store.fileExists() {
		return nil, nil
	}

	file, err := store.openFile(false)
	if err != nil {
		return nil, err
	}

	var token oauth2.Token

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&token)

	return &token, err
}

// Persist writes a token to a file
func (store FileTokenStorage) Persist(token *oauth2.Token) error {
	if token == nil {
		return errors.New("tried to persist empty token")
	}

	data, err := json.Marshal(token)
	if err != nil {
		return err
	}

	file, err := store.openFile(true)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}

// Wipe removes existing stored token
func (store FileTokenStorage) Wipe() error {
	return os.Remove(store.filepath)
}

func (store FileTokenStorage) openFile(create bool) (*os.File, error) {
	if create {
		return os.Create(store.filepath)
	}

	return os.Open(store.filepath)
}

func (store FileTokenStorage) fileExists() bool {
	info, err := os.Stat(store.filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()

}
