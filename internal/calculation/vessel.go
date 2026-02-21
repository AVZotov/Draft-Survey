package calculation

type PPDirection string

const (
	PPDirectionForward PPDirection = "F"
	PPDirectionAft     PPDirection = "A"
)

type VesselGeometry struct {
	DistancePPFwd  float64
	DistancePPMid  float64
	DistancePPAft  float64
	PPFwdDirection PPDirection
	PPMidDirection PPDirection
	PPAftDirection PPDirection
	LBP            float64
	KeelFwd        float64
	KeelMid        float64
	KeelAft        float64
}
