package types

type LCFDirection string

const (
	LCFDirectionForward LCFDirection = "F"
	LCFDirectionAft     LCFDirection = "A"
	LCFDirectionFromAP  LCFDirection = "AP"
)

type HydrostaticRow struct {
	Draft        *float64     `json:"draft"`
	Displacement *float64     `json:"displacement"`
	TPC          *float64     `json:"tpc"`
	LCF          *float64     `json:"lcf"`
	LCFDirection LCFDirection `json:"lcf_direction"`
}

type MTCRow struct {
	Draft *float64 `json:"draft"`
	MTC   *float64 `json:"mtc"`
}

type Hydrostatics struct {
	Displacement float64 `json:"displacement"`
	TPC          float64 `json:"tpc"`
	LCF          float64 `json:"lcf"`
}
