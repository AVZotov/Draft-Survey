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

func (J JSONStore) Save(id string, survey *types.Survey) error {
	filename := id + ".json"
	path := filepath.Join(J.Path, filename)
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
