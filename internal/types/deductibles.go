package types

type OtherDeductibles struct {
	Others     float64 `json:"others"`
	OthersName string  `json:"others_name"`
}

type FreshWaterTank struct {
	Name     string  `json:"name"`
	Sounding float64 `json:"sounding"`
	Volume   float64 `json:"volume"`
}

func (fwt FreshWaterTank) GetWeight() float64 {
	const density = 1.0
	return fwt.Volume * density
}

type BallastWaterTank struct {
	Name     string  `json:"name"`
	Sounding float64 `json:"sounding"`
	Volume   float64 `json:"volume"`
	Density  float64 `json:"density"`
}

func (bwt BallastWaterTank) GetWeight() float64 {
	return bwt.Volume * bwt.Density
}

type Deductibles struct {
	HFO         float64 `json:"hfo"`
	MDO         float64 `json:"mdo"`
	LubOil      float64 `json:"lub_oil"`
	BilgeWater  float64 `json:"bilge_water"`
	SewageWater float64 `json:"sewage_water"`
	OtherDeductibles
}
