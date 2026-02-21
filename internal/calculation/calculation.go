package calculation

import "math"

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
	lbm := round3(v.LBP - dAftDir + dFwdDir)
	return PPCorrections{
		FWDCorrection: round3(dFwdDir * trim / lbm),
		MIDCorrection: round3(dMidDir * trim / lbm),
		AFTCorrection: round3(dAftDir * trim / lbm),
	}
}

func CalcHalfLBPPPCorrections(m MeanDraft, v Vessel) PPCorrections {
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

	lbmMidFwd := round3((v.LBP / 2) - dMidDir - dFwdDir)
	lbmAftMid := round3((v.LBP / 2) - dAftDir - dMidDir)

	fwdCorr := round3(dFwdDir * (m.DraftMIDmean - m.DraftFWDmean) / lbmMidFwd)
	midCorr := round3(dMidDir * (m.DraftMIDmean - m.DraftFWDmean) / lbmMidFwd)
	midWKeel := round3(m.DraftMIDmean + midCorr - (v.KeelMID / 1000))
	aftCorr := round3(dAftDir * (m.DraftAFTmean - midWKeel) / lbmAftMid)

	return PPCorrections{
		FWDCorrection: fwdCorr,
		MIDCorrection: midCorr,
		AFTCorrection: aftCorr,
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
		return round3((draftsWKeel.FWDDraftWKeel + round3(6*draftsWKeel.MIDDraftWKeel) + draftsWKeel.AFTDraftWKeel) / 8)
	}

	if vesselType == VesselTypeRiver {
		return round3((draftsWKeel.FWDDraftWKeel + round3(4*draftsWKeel.MIDDraftWKeel) + draftsWKeel.AFTDraftWKeel) / 6)
	}

	if vesselType == VesselTypeBarge {
		return round3((round3(3*draftsWKeel.FWDDraftWKeel) + round3(14*draftsWKeel.MIDDraftWKeel) + round3(3*draftsWKeel.AFTDraftWKeel)) / 20)
	}

	return 0
}

func Interpolate(fact, lowerDraft, lowerValue, upperDraft, upperValue float64) float64 {
	result := round3(lowerValue + ((fact - lowerDraft) * (upperValue - lowerValue) / (upperDraft - lowerDraft)))
	return result
}

func CalcHydrostatics(mmc float64, hr []HydrostaticRow, vessel Vessel) Hydrostatics {
	var lower, upper HydrostaticRow
	if hr[0].Draft < hr[1].Draft {
		lower = hr[0]
		upper = hr[1]
	} else {
		lower = hr[1]
		upper = hr[0]
	}
	displacement := Interpolate(mmc, lower.Draft, lower.Displacement, upper.Draft, upper.Displacement)
	tpc := Interpolate(mmc, lower.Draft, lower.TPC, upper.Draft, upper.TPC)
	const k3 = 0.045
	lowerLcf := lower.LCF
	upperLcf := upper.LCF

	if lower.LCFDirection == LCFDirectionFromAP || lower.LCF > vessel.LBP*k3 {
		lowerLcf = (vessel.LBP / 2) - lower.LCF
		upperLcf = (vessel.LBP / 2) - upper.LCF
	} else {
		if lower.LCFDirection == LCFDirectionForward {
			lowerLcf *= -1
		}
		if upper.LCFDirection == LCFDirectionForward {
			upperLcf *= -1
		}
	}

	lcf := Interpolate(mmc, lower.Draft, lowerLcf, upper.Draft, upperLcf)

	return Hydrostatics{
		Displacement: displacement,
		TPC:          tpc,
		LCF:          lcf,
	}
}

func CalcFirstTrimCorrection(dwk DraftsWKeel, tpc float64, lcf float64, lbp float64) float64 {
	trueTrim := dwk.AFTDraftWKeel - dwk.FWDDraftWKeel
	var firstTrimCorrection float64

	if trueTrim < 0 && lcf >= 0 || trueTrim > 0 && lcf <= 0 {
		firstTrimCorrection = -1 * (math.Abs(trueTrim * tpc * lcf * 100 / lbp))
	} else {
		firstTrimCorrection = math.Abs(trueTrim * tpc * lcf * 100 / lbp)
	}

	return round3(firstTrimCorrection)
}

func CalcSecondTrimCorrection(dwk DraftsWKeel, mtcRows []MTCRow, lbp float64) float64 {
	var lowerMtcRow, upperMtcRow MTCRow
	if mtcRows[0].Draft < mtcRows[1].Draft {
		lowerMtcRow = mtcRows[0]
		upperMtcRow = mtcRows[1]
	} else {
		lowerMtcRow = mtcRows[1]
		upperMtcRow = mtcRows[0]
	}

	deltaMtc := upperMtcRow.MTC - lowerMtcRow.MTC
	trueTrim := dwk.AFTDraftWKeel - dwk.FWDDraftWKeel

	return round3(50 * math.Pow(trueTrim, 2) * deltaMtc / lbp)
}

func CalcListCorrection(marks Marks, tpcListPort, tpcListStarboard float64) float64 {
	if marks.MIDPort == marks.MIDStarboard {
		return 0.0
	}
	return round3(6 * math.Abs(marks.MIDPort-marks.MIDStarboard) * math.Abs(tpcListPort-tpcListStarboard))
}

func CalcDensityCorrection(displacement float64, firstTrim float64, secondTrim float64, listCorrection float64, density float64) float64 {
	displacementCorrected := round3(displacement + firstTrim + secondTrim + listCorrection)
	return round3(displacementCorrected * (density - 1.025) / 1.025)
}

func CalcTotalDeductibles(bwt []BallastWaterTank, fwt []FreshWaterTank, d Deductibles) float64 {
	tbw := TotalBallastWater(bwt)
	tfw := TotalFreshWater(fwt)

	return round3(tbw + tfw + d.HFO + d.MDO + d.Luboil + d.BilgeWater + d.SewageWater + d.Others)
}

func CalcNetDisplacement(displacement, firstTrim, secondTrim, listCorrection, densityCorrection, totalDeductibles float64) float64 {
	displCorrToDensity := round3(displacement + firstTrim + secondTrim + listCorrection + densityCorrection)
	return round3(displCorrToDensity - totalDeductibles)
}

func CalcCargoWeight(netDisplacementIni, netDisplacementFin float64) float64 {
	return round3(math.Abs(netDisplacementFin - netDisplacementIni))
}
