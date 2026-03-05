package handler

import (
	"errors"
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

	h.parseDraft(c, survey)

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

	h.parseDraft(c, survey)

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

	h.parseDraft(c, survey)

	if err = h.surveyRepository.Save(survey); err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

func (h *Handler) parseDraft(c *fiber.Ctx, survey *types.Survey) {
	total := len(survey.Drafts)
	for i := range survey.Drafts {
		prefix := draftPrefix(i, total)

		fwdPort, err := parseMark(c, fmt.Sprintf("fwd_port-d%d", i))
		if err == nil {
			survey.Drafts[i].Marks.FwdPort = types.Mark{Value: fwdPort}
		}
		midPort, err := parseMark(c, fmt.Sprintf("mid_port-d%d", i))
		if err == nil {
			survey.Drafts[i].Marks.MidPort = types.Mark{Value: midPort}
		}
		aftPort, err := parseMark(c, fmt.Sprintf("aft_port-d%d", i))
		if err == nil {
			survey.Drafts[i].Marks.AftPort = types.Mark{Value: aftPort}
		}
		fwdStbd, err := parseMark(c, fmt.Sprintf("fwd_stbd-d%d", i))
		if err == nil {
			survey.Drafts[i].Marks.FwdStarboard = types.Mark{Value: fwdStbd}
		}
		midStbd, err := parseMark(c, fmt.Sprintf("mid_stbd-d%d", i))
		if err == nil {
			survey.Drafts[i].Marks.MidStarboard = types.Mark{Value: midStbd}
		}
		aftStbd, err := parseMark(c, fmt.Sprintf("aft_stbd-d%d", i))
		if err == nil {
			survey.Drafts[i].Marks.AftStarboard = types.Mark{Value: aftStbd}
		}
		survey.Drafts[i].SeaCondition.Type = types.SeaConditionType(c.FormValue(fmt.Sprintf("%s_sea_type", prefix)))

		survey.Drafts[i].Deductibles.HFO = parseFloat(c, fmt.Sprintf("hfo-d%d", i))
		survey.Drafts[i].Deductibles.MDO = parseFloat(c, fmt.Sprintf("mdo-d%d", i))
		survey.Drafts[i].Deductibles.LubOil = parseFloat(c, fmt.Sprintf("lub-d%d", i))
		survey.Drafts[i].Deductibles.BilgeWater = parseFloat(c, fmt.Sprintf("bilW-d%d", i))
		survey.Drafts[i].Deductibles.Others = parseFloat(c, fmt.Sprintf("others-d%d", i))

		survey.Drafts[i].Density = parseFloat(c, fmt.Sprintf("dwDens-d%d", i))

		if survey.Drafts[i].Type != types.DraftTypeInitial {
			survey.Drafts[i].CargoDeclared = parseFloat(c, fmt.Sprintf("cargoDeclared-d%d", i))
		}
	}
	survey.Drafts[0].ConstantDeclared = parseFloat(c, "constDeclared")
}

func draftPrefix(i, total int) string {
	if i == 0 {
		return "i"
	}
	if i == total-1 {
		return "f"
	}
	return fmt.Sprintf("m%d", i)
}

func parseMark(c *fiber.Ctx, name string) (*float64, error) {
	var ErrEmptyField = errors.New("empty field")
	v := c.FormValue(name)
	if v == "" {
		return nil, ErrEmptyField
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return nil, ErrEmptyField
	}
	return &f, nil
}
