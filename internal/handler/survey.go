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

func (h *Handler) saveSurvey(c *fiber.Ctx) error {
	survey, err := h.getNewSurvey(c)
	if err != nil {
		return err
	}

	if err = h.surveyRepository.Save(survey); err != nil {
		return err
	}

	switch c.FormValue("next") {
	case "dashboard":
		c.Set("HX-Redirect", "/")
	case "draft":
		c.Set("HX-Redirect", "/survey/"+survey.ID+"/draft")
	case "stay":
		c.Set("HX-Redirect", "/survey/"+survey.ID)
	default:
		c.Set("HX-Redirect", "/")
	}
	return c.SendStatus(http.StatusOK)
}

func (h *Handler) getSurvey(c *fiber.Ctx) error {
	id := c.Params("id")
	survey, err := h.surveyRepository.Get(id)
	if err != nil {
		// TODO: Show error page
		return err
	}
	user, err := h.userRepository.Get()
	if err != nil {
		return err
	}
	props := web.NewSurveyProps(user, survey)
	return tadaptor.Render(c, web.NewSurvey(props))
}

/*
Helper function to parse survey data from the context recieved.
This function not a part of API
*/
func (h *Handler) getNewSurvey(c *fiber.Ctx) (*types.Survey, error) {
	user, err := h.userRepository.Get()
	if err != nil {
		return nil, err
	}

	createdAt := time.Now()
	if v := c.FormValue("created_at"); v != "" {
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			if !t.IsZero() {
				createdAt = t
			}

		}
	}

	job := types.Job{
		JobNumber: c.FormValue("job_no"),
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
		VesselType:       vessel.VesselType(c.FormValue("mmc_method")),
		CorrectionMethod: vessel.CorrectionMethod(c.FormValue("corr_method")),
	}

	seaCondition := types.SeaCondition{
		Type: types.SeaConditionType(c.FormValue("sea_type")),
		Wave: types.WaveCondition(c.FormValue("sea_condition")),
		Ice:  types.IceCondition(c.FormValue("ice_condition")),
	}
	survey := &types.Survey{
		Surveyor:       *user,
		Status:         types.SurveyStatusDraft,
		ID:             c.FormValue("survey_id"),
		CreatedAt:      createdAt,
		Job:            job,
		CargoOperation: cargoOperation,
		VesselData:     vesselData,
		SeaCondition:   seaCondition,
		Remarks:        c.FormValue("remarks"),
	}

	//TODO: Validation of fields

	return survey, nil
}
