package handler

import (
	"strconv"

	"github.com/AVZotov/draft-survey/internal/constants"
	"github.com/AVZotov/draft-survey/internal/handler/tadaptor"
	"github.com/AVZotov/draft-survey/internal/types"
	"github.com/AVZotov/draft-survey/web"
	"github.com/AVZotov/draft-survey/web/components"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) tanks(c *fiber.Ctx) error {
	id := c.Params("id")
	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return err
	}

	draftIndex := c.Params("draftIndex")

	user, err := h.userRepository.Get()
	if err != nil {
		return err
	}

	props := web.TanksPageProps(user, survey)
	return tadaptor.Render(c, web.Tanks(props, draftIndex))
}

func (h *Handler) addNewBwTank(c *fiber.Ctx) error {
	tankID := uuid.New().String()
	id := c.Params("id")
	draftIndex, err := strconv.Atoi(c.Params("draftIndex"))
	if err != nil {
		return err
	}
	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return err
	}

	tankName, err := h.parseNewTankName(c, constants.NewBwtType, constants.NewBWTName)
	if err != nil {
		return err
	}

	bwt := types.BallastWaterTank{
		ID:   tankID,
		Name: tankName,
	}

	survey.Drafts[draftIndex].BallastWaterTanks =
		append(survey.Drafts[draftIndex].BallastWaterTanks, bwt)

	if err = h.surveyRepository.Save(survey); err != nil {
		return err
	}

	return tadaptor.Render(c, components.BwTankItem(bwt))
}
