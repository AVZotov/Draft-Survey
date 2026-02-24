package handler

import "github.com/AVZotov/draft-survey/internal/storage"

type Handler struct {
	userRepository storage.UserRepository
}

func New(userRepository storage.UserRepository) *Handler {
	return &Handler{
		userRepository: userRepository,
	}
}
