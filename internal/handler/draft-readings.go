package handler

import (
	"fmt"
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

	h.parseDraftMarks(c, survey)

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

	h.parseDraftMarks(c, survey)

	survey.Drafts[index].Status = types.DraftStatusComplete
	survey.Drafts[index].FinishedAt = time.Now()

	if err = h.surveyRepository.Save(survey); err != nil {
		return err
	}

	c.Set("HX-Redirect", "/survey/"+id+"/draft")
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) addIntermediateDraft(c *fiber.Ctx) error {
	id := c.Params("id")
	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return err
	}

	survey.Drafts = append(survey.Drafts, types.Draft{
		Type:   types.DraftTypeIntermediate,
		Status: types.DraftStatusPending,
	})

	if err = h.surveyRepository.Save(survey); err != nil {
		return err
	}

	c.Set("HX-Redirect", "/survey/"+id+"/draft")
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) addFinalDraft(c *fiber.Ctx) error {
	id := c.Params("id")
	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return err
	}

	survey.Drafts = append(survey.Drafts, types.Draft{
		Type:   types.DraftTypeFinal,
		Status: types.DraftStatusPending,
	})

	if err = h.surveyRepository.Save(survey); err != nil {
		return err
	}

	c.Set("HX-Redirect", "/survey/"+id+"/draft")
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) saveDraft(c *fiber.Ctx) error {
	id := c.Params("id")
	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return err
	}

	h.parseDraftMarks(c, survey)

	if err = h.surveyRepository.Save(survey); err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

func (h *Handler) parseDraftMarks(c *fiber.Ctx, survey *types.Survey) {
	for i := range survey.Drafts {
		survey.Drafts[i].Marks.FwdPort = types.Mark{Value: parseFloat(c, fmt.Sprintf("fwd_port-d%d", i))}
		survey.Drafts[i].Marks.MidPort = types.Mark{Value: parseFloat(c, fmt.Sprintf("mid_port-d%d", i))}
		survey.Drafts[i].Marks.AftPort = types.Mark{Value: parseFloat(c, fmt.Sprintf("aft_port-d%d", i))}
		survey.Drafts[i].Marks.FwdStarboard = types.Mark{Value: parseFloat(c, fmt.Sprintf("fwd_stbd-d%d", i))}
		survey.Drafts[i].Marks.MidStarboard = types.Mark{Value: parseFloat(c, fmt.Sprintf("mid_stbd-d%d", i))}
		survey.Drafts[i].Marks.AftStarboard = types.Mark{Value: parseFloat(c, fmt.Sprintf("aft_stbd-d%d", i))}
	}
}
