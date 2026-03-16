package types

var _ Tank = (*BallastWaterTank)(nil)
var _ Tank = (*FreshWaterTank)(nil)

type Tank interface {
	GetWeight() float64
}

type FreshWaterTank struct {
	Type       string                `json:"tank_type"`
	Name       string                `json:"tank_name"`
	ID         string                `json:"tank_id"`
	Sounding   *float64              `json:"tank_sounding"`
	Volume     *float64              `json:"tank_volume"`
	Correction *VolumeCorrectionData `json:"correction"`
	Weight     float64               `json:"tank_weight"`
}

func (fwt FreshWaterTank) GetWeight() float64 {
	if fwt.Volume == nil {
		return 0.0
	}
	const density = 1.000
	return *fwt.Volume * density
}

type BallastWaterTank struct {
	Type       string                `json:"tank_type"`
	Name       string                `json:"tank_name"`
	ID         string                `json:"tank_id"`
	Sounding   *float64              `json:"tank_sounding"`
	Volume     *float64              `json:"tank_volume"`
	Density    *float64              `json:"tank_density"`
	Correction *VolumeCorrectionData `json:"correction"`
	Weight     float64               `json:"tank_weight"`
}

func (bwt BallastWaterTank) GetWeight() float64 {
	if bwt.Volume == nil || bwt.Density == nil {
		return 0
	}

	return *bwt.Volume * *bwt.Density
}

type OtherDeductibles struct {
	Others     *float64 `json:"others"`
	OthersName string   `json:"others_name"`
}
type Deductibles struct {
	HFO         *float64 `json:"hfo"`
	MDO         *float64 `json:"mdo"`
	LubOil      *float64 `json:"lub_oil"`
	BilgeWater  *float64 `json:"bilge_water"`
	SewageWater *float64 `json:"sewage_water"`
	OtherDeductibles
}
