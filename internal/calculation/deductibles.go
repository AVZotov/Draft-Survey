package calculation

type OtherDeductibles struct {
	Others     float64
	OthersName string
}

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
	LubOil      float64
	BilgeWater  float64
	SewageWater float64
	OtherDeductibles
}
