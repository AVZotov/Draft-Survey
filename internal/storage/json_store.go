package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/AVZotov/draft-survey/internal/types"
)

var _ SurveyRepository = (*JSONStore)(nil)

type JSONStore struct {
	Path     string
	TempPath string
}

func (j JSONStore) Save(id string, survey *types.Survey) error {
	filename := id + ".json"
	path := filepath.Join(j.Path, filename)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}(file)

	encoder := json.NewEncoder(file)
	if err = encoder.Encode(survey); err != nil {
		return err
	}
	return nil
}

func (j JSONStore) Get(id string) (*types.Survey, error) {
	filename := id + ".json"
	path := filepath.Join(j.Path, filename)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}(file)

	decoder := json.NewDecoder(file)
	survey := &types.Survey{}
	if err = decoder.Decode(survey); err != nil {
		return nil, err
	}

	return survey, nil
}

func (j JSONStore) GetBy(field SurveyField, value string) (*types.Survey, error) {
	//TODO implement me
	panic("implement me")
}

func (j JSONStore) GetAll() ([]*types.Survey, error) {
	//TODO implement me
	panic("implement me")
}

func (j JSONStore) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
