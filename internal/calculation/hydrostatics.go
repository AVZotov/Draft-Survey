package calculation

type LCFDirection string

const (
	LCFDirectionForward LCFDirection = "F"
	LCFDirectionAft     LCFDirection = "A"
	LCFDirectionFromAP  LCFDirection = "AP"
)

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

type Hydrostatics struct {
	Displacement float64
	TPC          float64
	LCF          float64
}
