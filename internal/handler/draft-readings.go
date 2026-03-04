package handler

import (
	"net/http"
	"strconv"
	"time"

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
		{
			Type:   types.DraftTypeInitial,
			Status: types.DraftStatusPending,
		},
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

func (h *Handler) startDraft(c *fiber.Ctx) error {
	id := c.Params("id")
	index, err := strconv.Atoi(c.Params("index"))
	if err != nil {
		return err
	}

	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return err
	}

	survey.Drafts[index].Status = types.DraftStatusActive
	survey.Drafts[index].StartedAt = time.Now()

	if err = h.surveyRepository.Save(survey); err != nil {
		return err
	}

	c.Set("HX-Redirect", "/survey/"+id+"/draft")
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) finishDraft(c *fiber.Ctx) error {
	id := c.Params("id")
	index, err := strconv.Atoi(c.Params("index"))
	if err != nil {
		return err
	}

	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return err
	}

	survey.Drafts[index].Status = types.DraftStatusComplete
	survey.Drafts[index].FinishedAt = time.Now()

	if err = h.surveyRepository.Save(survey); err != nil {
		return err
	}

	c.Set("HX-Redirect", "/survey/"+id+"/draft")
	return c.SendStatus(http.StatusOK)
}
