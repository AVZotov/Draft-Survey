package types

import (
	"time"

	"github.com/AVZotov/draft-survey/internal/vessel"
)

type MeanDraft struct {
	DraftFwdMean float64
	DraftMidMean float64
	DraftAftMean float64
}

type PPCorrections struct {
	FwdCorrection float64
	MidCorrection float64
	AftCorrection float64
}

type DraftsWKeel struct {
	FwdDraftWKeel float64
	MidDraftWKeel float64
	AftDraftWKeel float64
}

type InitialDraft struct {
	BallastWaterTanks []BallastWaterTank
	FreshWaterTanks   []FreshWaterTank
	Deductibles       Deductibles
	Marks             Marks
	ConstantDeclared  float64
	Density           float64
	StartedAt         time.Time
	FinishedAt        time.Time
	MTCRows           []MTCRow
	HydrostaticRows   []HydrostaticRow
	TPCListPort       float64
	TPCListStarboard  float64
	SeaCondition      SeaCondition
}

type FinalDraft struct {
	BallastWaterTanks []BallastWaterTank
	FreshWaterTanks   []FreshWaterTank
	Deductibles       Deductibles
	Marks             Marks
	CargoDeclared     float64
	Density           float64
	StartedAt         time.Time
	FinishedAt        time.Time
	MTCRows           []MTCRow
	HydrostaticRows   []HydrostaticRow
	TPCListPort       float64
	TPCListStarboard  float64
	SeaCondition      SeaCondition
}

type Job struct {
	JobNumber int
	DSNumber  int
	Principal string
}

type CargoOperation struct {
	PlaceOfInspection string
	Destination       string
	Operation         string
	Origin            string
	Cargo             string
	Packing           string
	Port              string
}

type Survey struct {
	Surveyor       *User
	ID             string
	InitialDraft   InitialDraft
	FinalDraft     FinalDraft
	Job            Job
	CargoOperation CargoOperation
	VesselData     vessel.VesselData
}
