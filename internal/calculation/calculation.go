package calculation

import "fmt"

func (fwt FreshWaterTank) GetWeight() float64 {
	return fwt.Volume
}

func (bwt BallastWaterTank) GetWeight() float64 {
	return bwt.Volume * bwt.Density
}

func TotalFreshWater(fwt []FreshWaterTank) float64 {
	var total float64
	for _, t := range fwt {
		total += round3(t.GetWeight())
	}
	return total
}

func TotalBallastWater(bwt []BallastWaterTank) float64 {
	var total float64
	for _, t := range bwt {
		total += round3(t.GetWeight())
	}
	return total
}

func MeanDrafts(m Marks) MeanDraft {
	return MeanDraft{
		DraftFWDmean: round3((m.FWDPort + m.FWDStarboard) / 2),
		DraftMIDmean: round3((m.MIDPort + m.MIDStarboard) / 2),
		DraftAFTmean: round3((m.AFTPort + m.AFTStarboard) / 2),
	}
}

func CalcFullLBPPPCorrections(m MeanDraft, v Vessel) PPCorrections {
	trim := m.DraftAFTmean - m.DraftFWDmean
	var dFwdDir, dMidDir, dAftDir float64

	if dFwdDir = v.DistancePPFWD; v.PPFWDDirection == PPDirectionAft {
		dFwdDir *= -1
	}
	if dMidDir = v.DistancePPMID; v.PPMIDDirection == PPDirectionAft {
		dMidDir *= -1
	}
	if dAftDir = v.DistancePPAFT; v.PPAFTDirection == PPDirectionAft {
		dAftDir *= -1
	}
	LBM := round3(v.LBP - dAftDir + dFwdDir)
	fmt.Println(LBM)
	return PPCorrections{
		FWDCorrection: round3(dFwdDir * trim / LBM),
		MIDCorrection: round3(dMidDir * trim / LBM),
		AFTCorrection: round3(dAftDir * trim / LBM),
	}
}

func CalcDraftsWKeel() {

}
