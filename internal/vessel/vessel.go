package vessel

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

type PPDirection string

const (
	PPDirectionForward PPDirection = "F"
	PPDirectionAft     PPDirection = "A"
)

type VesselData struct {
	Name                 string
	Flag                 string
	HomePort             string
	IMO                  string
	BuiltCountry         string
	BuiltYear            int
	HydrostaticDocsPhoto string
	Lightship            float64 // вес порожнем, MT
	Breadth              float64 // ширина корпуса, м
	Depth                float64 // высота корпуса, м
	LBP                  float64 // длина между перпендикулярами, м
	SummerDraft          float64
	SummerDWT            float64
	SummerTPC            float64
	SummerFreeboard      float64
	DistancePPFwd        float64
	PPFwdDirection       PPDirection
	DistancePPMid        float64
	PPMidDirection       PPDirection
	DistancePPAft        float64
	PPAftDirection       PPDirection
	KeelFwd              float64
	KeelMid              float64
	KeelAft              float64
	VesselType           VesselType
	CorrectionMethod     CorrectionMethod
}
