package handler

import (
	"errors"
	"net/http"
	"slices"
	"strconv"

	"github.com/AVZotov/draft-survey/internal/calculation"
	"github.com/AVZotov/draft-survey/internal/handler/tadaptor"
	"github.com/AVZotov/draft-survey/internal/types"
	"github.com/AVZotov/draft-survey/web"
	"github.com/AVZotov/draft-survey/web/components"
	"github.com/AVZotov/draft-survey/web/templates/pages"
	"github.com/AVZotov/draft-survey/web/widgets/tanks"
	"github.com/AVZotov/draft-survey/web/widgets/tanks/corrections"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type props struct {
	surveyID   string
	draftIndex int
	survey     *types.Survey
	user       *types.User
	tankID     string
	tankIndex  int
	tanks      []types.Tank
}

func (h *Handler) tanks(c *fiber.Ctx) error {
	props, err := getProps(h, c)
	if err != nil {
		return err
	}

	tanksLayoutProps := web.TanksLayoutProps(props.user)
	tanksPageProps := web.TanksPageProps(*props.survey, props.draftIndex)

	return tadaptor.Render(c, pages.Tanks(tanksLayoutProps, tanksPageProps))
}

func (h *Handler) newBwTank(c *fiber.Ctx) error {
	tankID := uuid.New().String()
	p, err := getProps(h, c)
	if err != nil {
		return err
	}

	bwt := types.Tank{
		ID: tankID,
	}
	h.parseBwTank(c, &bwt)

	p.survey.Drafts[p.draftIndex].BallastWaterTanks =
		append(p.survey.Drafts[p.draftIndex].BallastWaterTanks, bwt)

	if err = h.surveyRepository.Save(p.survey); err != nil {
		return err
	}

	return tadaptor.Render(c, templ.Join(
		components.TankItem(p.survey.ID, p.draftIndex, bwt),
		tanks.BwAddRowForm(p.survey.ID, p.draftIndex, true)))
}

func (h *Handler) deleteBwTank(c *fiber.Ctx) error {
	p, err := getProps(h, c)
	if err != nil {
		return err
	}

	p.tanks = slices.Delete(p.tanks, p.tankIndex, p.tankIndex+1)

	p.survey.Drafts[p.draftIndex].BallastWaterTanks = p.tanks
	if err := h.surveyRepository.Save(p.survey); err != nil {
		return err
	}

	tanksProps := web.TanksPageProps(*p.survey, p.draftIndex)

	c.Status(http.StatusOK)
	return tadaptor.Render(c, tanks.BwTableHeaderForm(tanksProps, true))
}

func (h *Handler) updateBwTank(c *fiber.Ctx) error {
	p, err := getProps(h, c)
	if err != nil {
		return err
	}

	tank := p.tanks[p.tankIndex]
	h.parseBwTank(c, &tank)

	p.survey.Drafts[p.draftIndex].BallastWaterTanks[p.tankIndex] = tank

	if err := h.surveyRepository.Save(p.survey); err != nil {
		return err
	}

	tanksProps := web.TanksPageProps(*p.survey, p.draftIndex)

	c.Status(http.StatusOK)
	return tadaptor.Render(c, templ.Join(
		components.TankItem(p.survey.ID, p.draftIndex, tank),
		tanks.BwTableHeaderForm(tanksProps, true)))
}

func (h *Handler) tanksCorrections(c *fiber.Ctx) error {
	p, err := getProps(h, c)
	if err != nil {
		return err
	}

	tank := p.tanks[p.tankIndex]
	//TODO: Implement loading logic with corrections struct parsing
	draftCalcs := calculation.CalcDraft(p.survey.Drafts[p.draftIndex], p.survey.VesselData)
	tanksProps := web.TanksPageProps(*p.survey, p.draftIndex)
	callibrationProps := web.CalibrationPageProps(tank, draftCalcs.TrueTrim, draftCalcs.ListDegrees)

	c.Status(http.StatusOK)
	return tadaptor.Render(c, corrections.ModalForm(tanksProps, callibrationProps))
}

func getProps(h *Handler, c *fiber.Ctx) (*props, error) {
	id := c.Params("id")
	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		return nil, err
	}

	draftIndexStr := c.Params("draftIndex")
	draftIndex, err := strconv.Atoi(draftIndexStr)
	if err != nil {
		return nil, err
	}

	user, err := h.userRepository.Get()
	if err != nil {
		return nil, err
	}

	tankID := c.Params("tankID")

	if tankID == "" {
		return &props{
			surveyID:   id,
			draftIndex: draftIndex,
			survey:     survey,
			user:       user,
		}, nil
	}

	tanks := survey.Drafts[draftIndex].BallastWaterTanks
	tankIndex := slices.IndexFunc(tanks, func(tank types.Tank) bool {
		return tank.ID == tankID
	})
	if tankIndex == -1 {
		return nil, errors.New("tank not found")
	}

	return &props{
		surveyID:   id,
		draftIndex: draftIndex,
		survey:     survey,
		user:       user,
		tankID:     tankID,
		tankIndex:  tankIndex,
		tanks:      tanks,
	}, nil
}
