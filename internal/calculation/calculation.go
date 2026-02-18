package calculation

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
