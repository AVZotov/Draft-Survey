package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/AVZotov/draft-survey/internal/types"
)

var _ SurveyRepository = (*SurveyStore)(nil)

const extension = ".json"

type SurveyStore struct {
	Path     string
	TempPath string
}

func NewSurveyStore(path, tempPath string) (*SurveyStore, error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(tempPath, 0755); err != nil {
		return nil, err
	}

	return &SurveyStore{
		Path:     path,
		TempPath: tempPath,
	}, nil
}

func (j *SurveyStore) Save(survey *types.Survey) error {
	filename := survey.ID + extension
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

func (j *SurveyStore) Get(id string) (*types.Survey, error) {
	filename := id + extension
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

func (j *SurveyStore) GetAll() ([]*types.Survey, error) {
	var surveys []*types.Survey
	files, err := os.ReadDir(j.Path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != extension {
			continue
		}
		id := strings.TrimSuffix(file.Name(), extension)
		survey, err := j.Get(id)
		if err != nil {
			return nil, err
		}
		surveys = append(surveys, survey)
	}
	return surveys, nil
}

func (j *SurveyStore) Delete(id string) error {
	return os.Remove(filepath.Join(j.Path, id+extension))
}
