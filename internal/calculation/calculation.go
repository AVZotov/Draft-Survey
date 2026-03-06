package calculation

import (
	"math"

	"github.com/AVZotov/draft-survey/internal/types"
	"github.com/AVZotov/draft-survey/internal/vessel"
)

func TotalFreshWater(fwt []types.FreshWaterTank) float64 {
	var total float64
	for _, t := range fwt {
		total += round3(t.GetWeight())
	}
	return total
}

func TotalBallastWater(bwt []types.BallastWaterTank) float64 {
	var total float64
	for _, t := range bwt {
		total += round3(t.GetWeight())
	}
	return total
}

func markVal(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}

func MeanDrafts(m types.Marks) types.MeanDraft {
	return types.MeanDraft{
		DraftFwdMean: round3((markVal(m.FwdPort.Value) + markVal(m.FwdStarboard.Value)) / 2),
		DraftMidMean: round3((markVal(m.MidPort.Value) + markVal(m.MidStarboard.Value)) / 2),
		DraftAftMean: round3((markVal(m.AftPort.Value) + markVal(m.AftStarboard.Value)) / 2),
	}
}

func CalcFullLBPPPCorrections(m types.MeanDraft, draft types.Draft, lbp float64) types.PPCorrections {
	trim := m.DraftAftMean - m.DraftFwdMean
	var dFwdDir, dMidDir, dAftDir float64

	if dFwdDir = markVal(draft.DistancePPFwd); draft.PPFwdDirection == "A" {
		dFwdDir *= -1
	}
	if dMidDir = markVal(draft.DistancePPMid); draft.PPMidDirection == "A" {
		dMidDir *= -1
	}
	if dAftDir = markVal(draft.DistancePPAft); draft.PPAftDirection == "A" {
		dAftDir *= -1
	}
	lbm := round3(lbp - dAftDir + dFwdDir)
	return types.PPCorrections{
		FwdCorrection: round3(dFwdDir * trim / lbm),
		MidCorrection: round3(dMidDir * trim / lbm),
		AftCorrection: round3(dAftDir * trim / lbm),
	}
}

func CalcHalfLBPPPCorrections(m types.MeanDraft, draft types.Draft, lbp float64) types.PPCorrections {
	var dFwdDir, dMidDir, dAftDir float64

	if dFwdDir = markVal(draft.DistancePPFwd); draft.PPFwdDirection == "A" {
		dFwdDir *= -1
	}
	if dMidDir = markVal(draft.DistancePPMid); draft.PPMidDirection == "A" {
		dMidDir *= -1
	}
	if dAftDir = markVal(draft.DistancePPAft); draft.PPAftDirection == "A" {
		dAftDir *= -1
	}

	lbmMidFwd := round3((lbp / 2) - dMidDir - dFwdDir)
	lbmAftMid := round3((lbp / 2) - dAftDir - dMidDir)

	fwdCorr := round3(dFwdDir * (m.DraftMidMean - m.DraftFwdMean) / lbmMidFwd)
	midCorr := round3(dMidDir * (m.DraftMidMean - m.DraftFwdMean) / lbmMidFwd)
	midWKeel := round3(m.DraftMidMean + midCorr - (markVal(draft.KeelMid) / 1000))
	aftCorr := round3(dAftDir * (m.DraftAftMean - midWKeel) / lbmAftMid)

	return types.PPCorrections{
		FwdCorrection: fwdCorr,
		MidCorrection: midCorr,
		AftCorrection: aftCorr,
	}
}

func CalcDraftsWKeel(
	meanDraft types.MeanDraft, ppCorrections types.PPCorrections, draft types.Draft) types.DraftsWKeel {
	keelCorrectionFwd := -1 * markVal(draft.KeelFwd) / 1000
	keelCorrectionMid := -1 * markVal(draft.KeelMid) / 1000
	keelCorrectionAft := -1 * markVal(draft.KeelAft) / 1000

	return types.DraftsWKeel{
		FwdDraftWKeel: round3(meanDraft.DraftFwdMean + ppCorrections.FwdCorrection + keelCorrectionFwd),
		MidDraftWKeel: round3(meanDraft.DraftMidMean + ppCorrections.MidCorrection + keelCorrectionMid),
		AftDraftWKeel: round3(meanDraft.DraftAftMean + ppCorrections.AftCorrection + keelCorrectionAft),
	}
}

func CalcMMC(draftsWKeel types.DraftsWKeel, v vessel.VesselData) float64 {
	if v.VesselType == vessel.VesselTypeMarine {
		return round3((draftsWKeel.FwdDraftWKeel + round3(6*draftsWKeel.MidDraftWKeel) + draftsWKeel.AftDraftWKeel) / 8)
	}

	if v.VesselType == vessel.VesselTypeRiver {
		return round3((draftsWKeel.FwdDraftWKeel + round3(4*draftsWKeel.MidDraftWKeel) + draftsWKeel.AftDraftWKeel) / 6)
	}

	if v.VesselType == vessel.VesselTypeBarge {
		return round3((round3(3*draftsWKeel.FwdDraftWKeel) + round3(14*draftsWKeel.MidDraftWKeel) + round3(3*draftsWKeel.AftDraftWKeel)) / 20)
	}

	return 0
}

func Interpolate(fact, lowerDraft, lowerValue, upperDraft, upperValue float64) float64 {
	result := round3(lowerValue + ((fact - lowerDraft) * (upperValue - lowerValue) / (upperDraft - lowerDraft)))
	return result
}

func CalcHydrostatics(mmc float64, hr []types.HydrostaticRow, v vessel.VesselData) types.Hydrostatics {
	var lower, upper types.HydrostaticRow
	if markVal(hr[0].Draft) < markVal(hr[1].Draft) {
		lower = hr[0]
		upper = hr[1]
	} else {
		lower = hr[1]
		upper = hr[0]
	}
	displacement := Interpolate(mmc, markVal(lower.Draft), markVal(lower.Displacement), markVal(upper.Draft), markVal(upper.Displacement))
	tpc := Interpolate(mmc, markVal(lower.Draft), markVal(lower.TPC), markVal(upper.Draft), markVal(upper.TPC))
	const k3 = 0.045
	lowerLcf := markVal(lower.LCF)
	upperLcf := markVal(upper.LCF)

	if lower.LCFDirection == types.LCFDirectionFromAP || markVal(lower.LCF) > v.LBP*k3 {
		lowerLcf = (v.LBP / 2) - markVal(lower.LCF)
		upperLcf = (v.LBP / 2) - markVal(upper.LCF)
	} else {
		if lower.LCFDirection == types.LCFDirectionForward {
			lowerLcf *= -1
		}
		if upper.LCFDirection == types.LCFDirectionForward {
			upperLcf *= -1
		}
	}

	lcf := Interpolate(mmc, markVal(lower.Draft), lowerLcf, markVal(upper.Draft), upperLcf)

	return types.Hydrostatics{
		Displacement: displacement,
		TPC:          tpc,
		LCF:          lcf,
	}
}

func CalcFirstTrimCorrection(dwk types.DraftsWKeel, tpc float64, lcf float64, lbp float64) float64 {
	trueTrim := dwk.AftDraftWKeel - dwk.FwdDraftWKeel
	var firstTrimCorrection float64

	if trueTrim < 0 && lcf >= 0 || trueTrim > 0 && lcf <= 0 {
		firstTrimCorrection = -1 * (math.Abs(trueTrim * tpc * lcf * 100 / lbp))
	} else {
		firstTrimCorrection = math.Abs(trueTrim * tpc * lcf * 100 / lbp)
	}

	return round3(firstTrimCorrection)
}

func CalcSecondTrimCorrection(dwk types.DraftsWKeel, mtcRows []types.MTCRow, lbp float64) float64 {
	var lowerMtcRow, upperMtcRow types.MTCRow
	if markVal(mtcRows[0].Draft) < markVal(mtcRows[1].Draft) {
		lowerMtcRow = mtcRows[0]
		upperMtcRow = mtcRows[1]
	} else {
		lowerMtcRow = mtcRows[1]
		upperMtcRow = mtcRows[0]
	}

	deltaMtc := markVal(upperMtcRow.MTC) - markVal(lowerMtcRow.MTC)
	trueTrim := dwk.AftDraftWKeel - dwk.FwdDraftWKeel

	return round3(50 * math.Pow(trueTrim, 2) * deltaMtc / lbp)
}

func CalcListCorrection(marks types.Marks, tpcListPort, tpcListStarboard float64) float64 {
	if markVal(marks.MidPort.Value) == markVal(marks.MidStarboard.Value) {
		return 0.0
	}
	return round3(6 * math.Abs(markVal(marks.MidPort.Value)-markVal(marks.MidStarboard.Value)) * math.Abs(tpcListPort-tpcListStarboard))
}

func CalcDensityCorrection(displacement float64, firstTrim float64, secondTrim float64, listCorrection float64, density float64) float64 {
	displacementCorrected := round3(displacement + firstTrim + secondTrim + listCorrection)
	return round3(displacementCorrected * (density - 1.025) / 1.025)
}

func CalcTotalDeductibles(bwt []types.BallastWaterTank, fwt []types.FreshWaterTank, d types.Deductibles) float64 {
	tbw := TotalBallastWater(bwt)
	tfw := TotalFreshWater(fwt)

	return round3(tbw + tfw + markVal(d.HFO) + markVal(d.MDO) + markVal(d.LubOil) + markVal(d.BilgeWater) + markVal(d.SewageWater) + markVal(d.Others))
}

func CalcNetDisplacement(displacement, firstTrim, secondTrim, listCorrection, densityCorrection, totalDeductibles float64) float64 {
	displCorrToDensity := round3(displacement + firstTrim + secondTrim + listCorrection + densityCorrection)
	return round3(displCorrToDensity - totalDeductibles)
}

func CalcCargoWeight(netDisplacementIni, netDisplacementFin float64) float64 {
	return round3(math.Abs(netDisplacementFin - netDisplacementIni))
}

func CalcConstant(netDisplacementIni float64, lightship float64) float64 {
	return round3(netDisplacementIni - lightship)
}

func CalcCurrentDWT(displCorrToDensity float64, lightship float64) float64 {
	return round3(displCorrToDensity - lightship)
}
