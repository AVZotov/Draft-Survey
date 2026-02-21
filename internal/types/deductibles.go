package types

type OtherDeductibles struct {
	Others     float64
	OthersName string
}

type FreshWaterTank struct {
	Name     string
	Sounding float64
	Volume   float64
}

func (fwt FreshWaterTank) GetWeight() float64 {
	const density = 1.0
	return fwt.Volume * density
}

type BallastWaterTank struct {
	Name     string
	Sounding float64
	Volume   float64
	Density  float64
}

func (bwt BallastWaterTank) GetWeight() float64 {
	return bwt.Volume * bwt.Density
}

type Deductibles struct {
	HFO         float64
	MDO         float64
	LubOil      float64
	BilgeWater  float64
	SewageWater float64
	OtherDeductibles
}
