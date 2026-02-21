package types

import (
	"time"
)

type MeanDraft struct {
	DraftFWDmean float64
	DraftMIDmean float64
	DraftAFTmean float64
}

type PPCorrections struct {
	FWDCorrection float64
	MIDCorrection float64
	AFTCorrection float64
}

type DraftsWKeel struct {
	FWDDraftWKeel float64
	MIDDraftWKeel float64
	AFTDraftWKeel float64
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
	InitialDraft   InitialDraft
	FinalDraft     FinalDraft
	Job            Job
	CargoOperation CargoOperation
	VesselGeometry VesselGeometry
}
