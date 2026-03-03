package types

import (
	"time"

	"github.com/AVZotov/draft-survey/internal/vessel"
)

type MeanDraft struct {
	DraftFwdMean float64 `json:"draft_fwd_mean"`
	DraftMidMean float64 `json:"draft_mid_mean"`
	DraftAftMean float64 `json:"draft_aft_mean"`
}

type PPCorrections struct {
	FwdCorrection float64 `json:"fwd_correction"`
	MidCorrection float64 `json:"mid_correction"`
	AftCorrection float64 `json:"aft_correction"`
}

type DraftsWKeel struct {
	FwdDraftWKeel float64 `json:"fwd_draft_w_keel"`
	MidDraftWKeel float64 `json:"mid_draft_w_keel"`
	AftDraftWKeel float64 `json:"aft_draft_w_keel"`
}

type InitialDraft struct {
	BallastWaterTanks []BallastWaterTank `json:"ballast_water_tanks"`
	FreshWaterTanks   []FreshWaterTank   `json:"fresh_water_tanks"`
	Deductibles       Deductibles        `json:"deductibles"`
	Marks             Marks              `json:"marks"`
	ConstantDeclared  float64            `json:"constant_declared"`
	Density           float64            `json:"density"`
	StartedAt         time.Time          `json:"started_at"`
	FinishedAt        time.Time          `json:"finished_at"`
	MTCRows           []MTCRow           `json:"mtc_rows"`
	HydrostaticRows   []HydrostaticRow   `json:"hydrostatic_rows"`
	TPCListPort       float64            `json:"tpc_list_port"`
	TPCListStarboard  float64            `json:"tpc_list_starboard"`
	SeaCondition      SeaCondition       `json:"sea_condition"`
}

type FinalDraft struct {
	BallastWaterTanks []BallastWaterTank `json:"ballast_water_tanks"`
	FreshWaterTanks   []FreshWaterTank   `json:"fresh_water_tanks"`
	Deductibles       Deductibles        `json:"deductibles"`
	Marks             Marks              `json:"marks"`
	CargoDeclared     float64            `json:"cargo_declared"`
	Density           float64            `json:"density"`
	StartedAt         time.Time          `json:"started_at"`
	FinishedAt        time.Time          `json:"finished_at"`
	MTCRows           []MTCRow           `json:"mtc_rows"`
	HydrostaticRows   []HydrostaticRow   `json:"hydrostatic_rows"`
	TPCListPort       float64            `json:"tpc_list_port"`
	TPCListStarboard  float64            `json:"tpc_list_starboard"`
	SeaCondition      SeaCondition       `json:"sea_condition"`
}

type Job struct {
	JobNumber int    `json:"job_number"`
	DSNumber  int    `json:"ds_number"`
	Principal string `json:"principal"`
}

type CargoOperation struct {
	PlaceOfInspection string `json:"place_of_inspection"`
	Destination       string `json:"destination"`
	Operation         string `json:"operation"`
	Cargo             string `json:"cargo"`
	Packing           string `json:"packing"`
}

type SurveyStatus string

const (
	SurveyStatusDraft      SurveyStatus = "draft"
	SurveyStatusInProgress SurveyStatus = "in_progress"
	SurveyStatusComplete   SurveyStatus = "complete"
)

type Survey struct {
	Surveyor       User              `json:"surveyor"`
	Status         SurveyStatus      `json:"status"`
	ID             string            `json:"id"`
	CreatedAt      time.Time         `json:"created_at"`
	InitialDraft   InitialDraft      `json:"initial_draft"`
	FinalDraft     FinalDraft        `json:"final_draft"`
	Job            Job               `json:"job"`
	CargoOperation CargoOperation    `json:"cargo_operation"`
	VesselData     vessel.VesselData `json:"vessel_data"`
	SeaCondition   SeaCondition      `json:"sea_condition"`
	Remarks        string            `json:"remarks"`
}
