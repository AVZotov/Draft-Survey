package calculation

import (
	"time"
)

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

type PPDirection string

const (
	PPDirectionForward PPDirection = "F"
	PPDirectionAft     PPDirection = "A"
)

type Vessel struct {
	Name              string
	Flag              string
	HomePort          string
	IMO               string
	BuiltCountry      string
	BuiltYear         int
	HydrostaticTables string
	DistancePPFWD     float64
	DistancePPMID     float64
	DistancePPAFT     float64
	PPFWDDirection    PPDirection
	PPMIDDirection    PPDirection
	PPAFTDirection    PPDirection
	LBP               float64
	KeelFWD           float64
	KeelMID           float64
	KeelAFT           float64
	VesselType        VesselType
	CorrectionMethod  CorrectionMethod
}

type Survey struct {
	InitialDraft   InitialDraft
	FinalDraft     FinalDraft
	Job            Job
	CargoOperation CargoOperation
	Vessel         Vessel
}

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

type VesselType string

const (
	VesselTypeMarine VesselType = "marine"
	VesselTypeRiver  VesselType = "river"
	VesselTypeBarge  VesselType = "barge"
)

type CorrectionMethod string

const (
	CorrectionMethodFullLBP CorrectionMethod = "Full LBP"
	CorrectionMethodHalfLBP CorrectionMethod = "Half LBP"
)
