package handler

import (
	"net/http"
	"time"

	"github.com/AVZotov/draft-survey/internal/handler/tadaptor"
	"github.com/AVZotov/draft-survey/internal/types"
	"github.com/AVZotov/draft-survey/internal/vessel"
	"github.com/AVZotov/draft-survey/web"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) newSurvey(c *fiber.Ctx) error {
	id := uuid.New().String()
	survey := &types.Survey{ID: id}
	user, err := h.userRepository.Get()
	if err != nil {
		return err
	}
	props := web.NewSurveyProps(user, survey)
	component := web.NewSurvey(props)
	return tadaptor.Render(c, component)
}

func (h *Handler) createSurvey(c *fiber.Ctx) error {
	survey, err := h.getNewSurvey(c)
	if err = h.surveyRepository.Save(survey); err != nil {
		return err
	}

	c.Set("HX-Redirect", "/")
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) getNewSurvey(c *fiber.Ctx) (*types.Survey, error) {
	user, err := h.userRepository.Get()
	if err != nil {
		return nil, err
	}
	job := types.Job{
		JobNumber: parseInt(c, "job_no"),
		DSNumber:  parseInt(c, "ds_no"),
		Principal: c.FormValue("client"),
	}
	cargoOperation := types.CargoOperation{
		PlaceOfInspection: c.FormValue("port"),
		Destination:       c.FormValue("destination"),
		Operation:         c.FormValue("cargo_op"),
		Cargo:             c.FormValue("cargo"),
		Packing:           c.FormValue("packing"),
	}
	vesselData := vessel.VesselData{
		Name:             c.FormValue("vessel_name"),
		Flag:             c.FormValue("flag"),
		IMO:              c.FormValue("imo"),
		BuiltYear:        parseInt(c, "built"),
		Lightship:        parseFloat(c, "lightship"),
		Breadth:          parseFloat(c, "breadth"),
		Depth:            parseFloat(c, "depth"),
		LBP:              parseFloat(c, "lbp"),
		SummerDraft:      parseFloat(c, "summer_draft"),
		SummerDWT:        parseFloat(c, "summer_dwt"),
		SummerTPC:        parseFloat(c, "summer_tpc"),
		SummerFreeboard:  parseFloat(c, "summer_freeboard"),
		DistancePPFwd:    0,
		PPFwdDirection:   "",
		DistancePPMid:    0,
		PPMidDirection:   "",
		DistancePPAft:    0,
		PPAftDirection:   "",
		KeelFwd:          0,
		KeelMid:          0,
		KeelAft:          0,
		VesselType:       vessel.VesselType(c.FormValue("mmc_method")),
		CorrectionMethod: vessel.CorrectionMethod(c.FormValue("corr_method")),
	}
	survey := &types.Survey{
		Surveyor:       *user,
		Status:         types.SurveyStatusDraft,
		ID:             c.FormValue("survey_id"),
		CreatedAt:      time.Now(),
		Job:            job,
		CargoOperation: cargoOperation,
		VesselData:     vesselData,
	}

	//TODO: Validation of fields

	return survey, nil
}
