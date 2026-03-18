package handler

import (
	"errors"
	"net/http"
	"slices"
	"strconv"

	"github.com/AVZotov/draft-survey/internal/calculation"
	"github.com/AVZotov/draft-survey/internal/format"
	"github.com/AVZotov/draft-survey/internal/handler/tadaptor"
	"github.com/AVZotov/draft-survey/internal/types"
	"github.com/AVZotov/draft-survey/web"
	"github.com/AVZotov/draft-survey/web/components"
	"github.com/AVZotov/draft-survey/web/widgets/tanks"
	"github.com/AVZotov/draft-survey/web/widgets/tanks/corrections"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) tanks(c *fiber.Ctx) error {
	id := c.Params("id")
	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return err
	}

	draftIndexStr := c.Params("draftIndex")
	draftIndex, err := strconv.Atoi(draftIndexStr)
	if err != nil {
		return err
	}

	user, err := h.userRepository.Get()
	if err != nil {
		return err
	}

	var totalBwWeight, totalFwWeight string
	bwTanks := survey.Drafts[draftIndex].BallastWaterTanks
	if bwTanks != nil {
		totalBwWeight = format.WeightFormatted(calculation.TotalBallastWater(bwTanks))
	}
	fwTanks := survey.Drafts[draftIndex].FreshWaterTanks
	if fwTanks != nil {
		totalFwWeight = format.WeightFormatted(calculation.TotalFreshWater(fwTanks))
	}

	props := web.TanksPageProps(user, survey, draftIndexStr, string(survey.Drafts[draftIndex].Type), totalBwWeight, totalFwWeight)
	return tadaptor.Render(c, web.Tanks(props, bwTanks, fwTanks))
}

func (h *Handler) newBwTank(c *fiber.Ctx) error {
	tankID := uuid.New().String()
	id := c.Params("id")
	draftIndexStr := c.Params("draftIndex")
	draftIndex, err := strconv.Atoi(draftIndexStr)
	if err != nil {
		return err
	}
	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return err
	}

	bwt := new(types.BallastWaterTank)
	bwt.ID = tankID
	h.parseBwTank(c, bwt)

	survey.Drafts[draftIndex].BallastWaterTanks =
		append(survey.Drafts[draftIndex].BallastWaterTanks, *bwt)

	if err = h.surveyRepository.Save(survey); err != nil {
		return err
	}

	return tadaptor.Render(c, templ.Join(
		components.BwTankItem(survey.ID, draftIndexStr, *bwt),
		tanks.BwAddRowForm(survey.ID, draftIndexStr, true)))
}

func (h *Handler) deleteBwTank(c *fiber.Ctx) error {
	id := c.Params("id")
	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return err
	}
	draftIndexStr := c.Params("draftIndex")
	draftIndex, err := strconv.Atoi(draftIndexStr)
	if err != nil {
		return err
	}
	tankID := c.Params("tankID")
	bwTanks := survey.Drafts[draftIndex].BallastWaterTanks

	i := slices.IndexFunc(bwTanks, func(tank types.BallastWaterTank) bool {
		return tank.ID == tankID
	})
	if i == -1 {
		return errors.New("tank not found")
	}
	bwTanks = slices.Delete(bwTanks, i, i+1)

	survey.Drafts[draftIndex].BallastWaterTanks = bwTanks
	if err := h.surveyRepository.Save(survey); err != nil {
		return err
	}

	totalWeight := format.WeightFormatted(calculation.TotalBallastWater(survey.Drafts[draftIndex].BallastWaterTanks))

	c.Status(http.StatusOK)
	return tadaptor.Render(c, tanks.BwTableHeaderForm(
		string(survey.Drafts[draftIndex].Type), totalWeight, true))
}

func (h *Handler) updateBwTank(c *fiber.Ctx) error {
	id := c.Params("id")
	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return err
	}
	draftIndexStr := c.Params("draftIndex")
	draftIndex, err := strconv.Atoi(draftIndexStr)
	if err != nil {
		return err
	}
	tankID := c.Params("tankID")
	bwTanks := survey.Drafts[draftIndex].BallastWaterTanks
	tankIndex := slices.IndexFunc(bwTanks, func(tank types.BallastWaterTank) bool {
		return tank.ID == tankID
	})
	if tankIndex == -1 {
		return errors.New("tank not found")
	}

	bwt := bwTanks[tankIndex]

	h.parseBwTank(c, &bwt)

	if bwt.Volume != nil && bwt.Density != nil {
		bwt.Weight = bwt.GetWeight()
	}

	survey.Drafts[draftIndex].BallastWaterTanks[tankIndex] = bwt

	if err := h.surveyRepository.Save(survey); err != nil {
		return err
	}

	totalWeight := format.WeightFormatted(calculation.TotalBallastWater(survey.Drafts[draftIndex].BallastWaterTanks))

	c.Status(http.StatusOK)
	return tadaptor.Render(c, templ.Join(
		components.BwTankItem(survey.ID, draftIndexStr, bwt),
		tanks.BwTableHeaderForm(string(survey.Drafts[draftIndex].Type), totalWeight, true)))
}

func (h *Handler) bwTanksCorrections(c *fiber.Ctx) error {
	id := c.Params("id")
	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return err
	}

	draftIndexStr := c.Params("draftIndex")
	draftIndex, err := strconv.Atoi(draftIndexStr)
	if err != nil {
		return err
	}

	tankID := c.Params("tankID")
	bwTanks := survey.Drafts[draftIndex].BallastWaterTanks
	tankIndex := slices.IndexFunc(bwTanks, func(tank types.BallastWaterTank) bool {
		return tank.ID == tankID
	})
	if tankIndex == -1 {
		return errors.New("tank not found")
	}

	bwt := bwTanks[tankIndex]
	if bwt.Correction != nil {
		//TODO: Implement loading logic with corrections struct parsing
	}
	//TODO: Pass List and Trim to props using calculation module

	c.Status(http.StatusOK)
	return tadaptor.Render(c, corrections.ModalForm(web.TanksCorrProps(survey, &bwt, nil, nil)))
}
