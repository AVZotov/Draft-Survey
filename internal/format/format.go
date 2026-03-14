package format

import (
	"fmt"

	"github.com/AVZotov/draft-survey/internal/vessel"
)

func Mark(p *float64) string {
	if p == nil {
		return "—"
	}
	return fmt.Sprintf("%.3f", *p)
}

func FloatOrEmpty(v *float64) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%g", *v)
}

func Weight(v float64) string {
	return fmt.Sprintf("%.3f MT", v)
}

func Draft(v float64) string {
	return fmt.Sprintf("%.3f m", v)
}

func LBPCorrection(correction vessel.CorrectionMethod) string {
	lbpCorrections := vessel.CorrectionMethodFullLBP
	if correction == vessel.CorrectionMethodHalfLBP {
		lbpCorrections = vessel.CorrectionMethodHalfLBP
	}

	return string(lbpCorrections)
}

func MMCShortFormula(vesselType vessel.VesselType) string {
	mmcCorrections := "(6m/8)"
	if vesselType == vessel.VesselTypeRiver {
		mmcCorrections = "(4m/6)"
	}
	if vesselType == vessel.VesselTypeBarge {
		mmcCorrections = "(14m/20)"
	}

	return mmcCorrections
}
