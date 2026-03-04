package handler

import (
	"github.com/AVZotov/draft-survey/internal/handler/tadaptor"
	"github.com/AVZotov/draft-survey/internal/types"
	"github.com/AVZotov/draft-survey/web"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) draftReadings(c *fiber.Ctx) error {
	id := c.Params("id")
	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return err
	}
	user, err := h.userRepository.Get()
	if err != nil {
		return err
	}

	survey.Status = types.SurveyStatusInProgress

	drafts := []types.Draft{
		{Type: types.DraftTypeInitial},
		{Type: types.DraftTypeFinal},
	}

	if survey.Drafts == nil {
		survey.Drafts = drafts
		if err = h.surveyRepository.Save(survey); err != nil {
			return err
		}
	}

	props := web.DraftReadingsProps(user, survey)
	return tadaptor.Render(c, web.DraftReadings(props))
}
