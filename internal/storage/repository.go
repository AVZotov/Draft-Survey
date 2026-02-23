package storage

import (
	"github.com/AVZotov/draft-survey/internal/types"
)

type SurveyRepository interface {
	Save(id string, survey *types.Survey) error
	Get(id string) (*types.Survey, error)
	GetAll() ([]*types.Survey, error)
	Delete(id string) error
}

type UserRepository interface {
	Save(user *types.User) error
	Get() (*types.User, error)
	Delete() error
}
