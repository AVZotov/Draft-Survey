package calculation

import (
	"time"
)

type FreshWaterTank struct {
	Name     string
	Sounding float64
	Volume   float64
}

type BallastWaterTank struct {
	Name     string
	Sounding float64
	Volume   float64
	Density  float64
}

type Deductibles struct {
	HFO         float64
	MDO         float64
	Luboil      float64
	BilgeWater  float64
	SewageWater float64
	Others      float64
	OthersName  string
}

type Marks struct {
	FWDPort      float64
	FWDStarboard float64
	MIDPort      float64
	MIDStarboard float64
	AFTPort      float64
	AFTStarboard float64
}

type HydrostaticRow struct {
	Draft        float64
	Displacement float64
	TPC          float64
	LCF          float64
	LCFDirection LCFDirection
}

type MTCRow struct {
	Draft float64
	MTC   float64
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
}

type LCFDirection string

const (
	LCFDirectionForward LCFDirection = "F"
	LCFDirectionAft     LCFDirection = "A"
)

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
