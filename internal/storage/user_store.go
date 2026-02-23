package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/AVZotov/draft-survey/internal/types"
)

var _ UserRepository = (*UserStore)(nil)

const fileName = "user.json"

type UserStore struct {
	Path string
}

func (u UserStore) Save(user *types.User) error {
	path := filepath.Join(u.Path, fileName)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}(file)

	err = json.NewEncoder(file).Encode(user)
	if err != nil {
		return err
	}

	return nil
}

func (u UserStore) Get() (*types.User, error) {
	path := filepath.Join(u.Path, fileName)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}(file)

	user := &types.User{}
	err = json.NewDecoder(file).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserStore) Delete() error {
	err := os.Remove(filepath.Join(u.Path, fileName))
	if err != nil {
		return err
	}

	return nil
}
