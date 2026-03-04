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

	if len(survey.Drafts) == 0 {
		survey.Drafts = []types.Draft{
			{Type: types.DraftTypeInitial},
			{Type: types.DraftTypeFinal},
		}
	}

	props := web.DraftReadingsProps(user, survey)
	return tadaptor.Render(c, web.DraftReadings(props))

}
