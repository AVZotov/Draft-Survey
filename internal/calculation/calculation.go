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

func MeanDrafts(m types.Marks) types.MeanDraft {
	return types.MeanDraft{
		DraftFwdMean: round3((m.FwdPort.Value + m.FwdStarboard.Value) / 2),
		DraftMidMean: round3((m.MidPort.Value + m.MidStarboard.Value) / 2),
		DraftAftMean: round3((m.AftPort.Value + m.AftStarboard.Value) / 2),
	}
}

func CalcFullLBPPPCorrections(m types.MeanDraft, v vessel.VesselData) types.PPCorrections {
	trim := m.DraftAftMean - m.DraftFwdMean
	var dFwdDir, dMidDir, dAftDir float64

	if dFwdDir = v.DistancePPFwd; v.PPFwdDirection == vessel.PPDirectionAft {
		dFwdDir *= -1
	}
	if dMidDir = v.DistancePPMid; v.PPMidDirection == vessel.PPDirectionAft {
		dMidDir *= -1
	}
	if dAftDir = v.DistancePPAft; v.PPAftDirection == vessel.PPDirectionAft {
		dAftDir *= -1
	}
	lbm := round3(v.LBP - dAftDir + dFwdDir)
	return types.PPCorrections{
		FwdCorrection: round3(dFwdDir * trim / lbm),
		MidCorrection: round3(dMidDir * trim / lbm),
		AftCorrection: round3(dAftDir * trim / lbm),
	}
}

func CalcHalfLBPPPCorrections(m types.MeanDraft, v vessel.VesselData) types.PPCorrections {
	var dFwdDir, dMidDir, dAftDir float64

	if dFwdDir = v.DistancePPFwd; v.PPFwdDirection == vessel.PPDirectionAft {
		dFwdDir *= -1
	}
	if dMidDir = v.DistancePPMid; v.PPMidDirection == vessel.PPDirectionAft {
		dMidDir *= -1
	}
	if dAftDir = v.DistancePPAft; v.PPAftDirection == vessel.PPDirectionAft {
		dAftDir *= -1
	}

	lbmMidFwd := round3((v.LBP / 2) - dMidDir - dFwdDir)
	lbmAftMid := round3((v.LBP / 2) - dAftDir - dMidDir)

	fwdCorr := round3(dFwdDir * (m.DraftMidMean - m.DraftFwdMean) / lbmMidFwd)
	midCorr := round3(dMidDir * (m.DraftMidMean - m.DraftFwdMean) / lbmMidFwd)
	midWKeel := round3(m.DraftMidMean + midCorr - (v.KeelMid / 1000))
	aftCorr := round3(dAftDir * (m.DraftAftMean - midWKeel) / lbmAftMid)

	return types.PPCorrections{
		FwdCorrection: fwdCorr,
		MidCorrection: midCorr,
		AftCorrection: aftCorr,
	}
}

func CalcDraftsWKeel(
	meanDraft types.MeanDraft, ppCorrections types.PPCorrections, v vessel.VesselData) types.DraftsWKeel {
	keelCorrectionFwd := -1 * v.KeelFwd / 1000
	keelCorrectionMid := -1 * v.KeelMid / 1000
	keelCorrectionAft := -1 * v.KeelAft / 1000

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

	if lower.LCFDirection == types.LCFDirectionFromAP || lower.LCF > v.LBP*k3 {
		lowerLcf = (v.LBP / 2) - lower.LCF
		upperLcf = (v.LBP / 2) - upper.LCF
	} else {
		if lower.LCFDirection == types.LCFDirectionForward {
			lowerLcf *= -1
		}
		if upper.LCFDirection == types.LCFDirectionForward {
			upperLcf *= -1
		}
	}

	lcf := Interpolate(mmc, lower.Draft, lowerLcf, upper.Draft, upperLcf)

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
	if mtcRows[0].Draft < mtcRows[1].Draft {
		lowerMtcRow = mtcRows[0]
		upperMtcRow = mtcRows[1]
	} else {
		lowerMtcRow = mtcRows[1]
		upperMtcRow = mtcRows[0]
	}

	deltaMtc := upperMtcRow.MTC - lowerMtcRow.MTC
	trueTrim := dwk.AftDraftWKeel - dwk.FwdDraftWKeel

	return round3(50 * math.Pow(trueTrim, 2) * deltaMtc / lbp)
}

func CalcListCorrection(marks types.Marks, tpcListPort, tpcListStarboard float64) float64 {
	if marks.MidPort == marks.MidStarboard {
		return 0.0
	}
	return round3(6 * math.Abs(marks.MidPort.Value-marks.MidStarboard.Value) * math.Abs(tpcListPort-tpcListStarboard))
}

func CalcDensityCorrection(displacement float64, firstTrim float64, secondTrim float64, listCorrection float64, density float64) float64 {
	displacementCorrected := round3(displacement + firstTrim + secondTrim + listCorrection)
	return round3(displacementCorrected * (density - 1.025) / 1.025)
}

func CalcTotalDeductibles(bwt []types.BallastWaterTank, fwt []types.FreshWaterTank, d types.Deductibles) float64 {
	tbw := TotalBallastWater(bwt)
	tfw := TotalFreshWater(fwt)

	return round3(tbw + tfw + d.HFO + d.MDO + d.LubOil + d.BilgeWater + d.SewageWater + d.Others)
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
