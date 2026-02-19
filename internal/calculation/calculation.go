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
	return PPCorrections{
		FWDCorrection: round3(dFwdDir * trim / LBM),
		MIDCorrection: round3(dMidDir * trim / LBM),
		AFTCorrection: round3(dAftDir * trim / LBM),
	}
}

func CalcDraftsWKeel(meanDraft MeanDraft, ppCorrections PPCorrections, vessel Vessel) DraftsWKeel {
	keelCorrectionFwd := -1 * vessel.KeelFWD / 1000
	keelCorrectionMid := -1 * vessel.KeelMID / 1000
	keelCorrectionAft := -1 * vessel.KeelAFT / 1000

	return DraftsWKeel{
		FWDDraftWKeel: round3(meanDraft.DraftFWDmean + ppCorrections.FWDCorrection + keelCorrectionFwd),
		MIDDraftWKeel: round3(meanDraft.DraftMIDmean + ppCorrections.MIDCorrection + keelCorrectionMid),
		AFTDraftWKeel: round3(meanDraft.DraftAFTmean + ppCorrections.AFTCorrection + keelCorrectionAft),
	}
}

func CalcMMC(draftsWKeel DraftsWKeel, vesselType VesselType) float64 {
	if vesselType == VesselTypeMarine {
		return round3((draftsWKeel.FWDDraftWKeel + 6*draftsWKeel.MIDDraftWKeel + draftsWKeel.AFTDraftWKeel) / 8)
	}

	if vesselType == VesselTypeRiver {
		return round3((draftsWKeel.FWDDraftWKeel + 4*draftsWKeel.MIDDraftWKeel + draftsWKeel.AFTDraftWKeel) / 6)
	}

	if vesselType == VesselTypeBarge {
		return round3((3*draftsWKeel.FWDDraftWKeel + 14*draftsWKeel.MIDDraftWKeel + 3*draftsWKeel.AFTDraftWKeel) / 20)
	}

	return 0
}

func Interpolate(fact, lowerDraft, lowerValue, upperDraft, upperValue float64) float64 {
	result := round3(lowerValue + ((fact - lowerDraft) * (upperValue - lowerValue) / (upperDraft - lowerDraft)))
	return result
}

func CalcHydrostatics(mmc float64, hr []HydrostaticRow) (displacement float64, tpc float64, lcf float64) {
	var lower, upper HydrostaticRow
	if hr[0].Draft < hr[1].Draft {
		lower = hr[0]
		upper = hr[1]
	} else {
		lower = hr[1]
		upper = hr[0]
	}
	displacement = Interpolate(mmc, lower.Draft, lower.Displacement, upper.Draft, upper.Displacement)
	tpc = Interpolate(mmc, lower.Draft, lower.TPC, upper.Draft, upper.TPC)
	lowerLcf := lower.LCF
	upperLcf := upper.LCF

	if lower.LCFDirection == LCFDirectionAft {
		lowerLcf *= -1
	}
	if upper.LCFDirection == LCFDirectionAft {
		upperLcf *= -1
	}

	lcf = Interpolate(mmc, lower.Draft, lowerLcf, upper.Draft, upperLcf)

	return displacement, tpc, lcf
}
