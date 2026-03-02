package handler

import "github.com/AVZotov/draft-survey/internal/storage"

type Handler struct {
	userRepository   storage.UserRepository
	surveyRepository storage.SurveyRepository
}

func New(userRepository storage.UserRepository, surveyRepository storage.SurveyRepository) *Handler {
	return &Handler{
		userRepository:   userRepository,
		surveyRepository: surveyRepository,
	}
}
