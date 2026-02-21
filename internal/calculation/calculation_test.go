package calculation

import (
	"testing"

	"github.com/AVZotov/draft-survey/internal/types"
	"github.com/AVZotov/draft-survey/internal/vessel"
)

func getFreshWaterTank() types.FreshWaterTank {
	return types.FreshWaterTank{
		Name:     "test",
		Sounding: 3.5,
		Volume:   3.5,
	}
}

func getInitFreshWaterTanks() []types.FreshWaterTank {
	return []types.FreshWaterTank{
		{Name: "FW P", Sounding: 364.000, Volume: 364.000},
	}
}

func getBallastWaterTank() types.BallastWaterTank {
	return types.BallastWaterTank{
		Name:     "test",
		Sounding: 3.5,
		Volume:   3.5,
		Density:  1.025,
	}
}

func getInitBallastWaterTanks() []types.BallastWaterTank {
	return []types.BallastWaterTank{
		{Name: "FPT", Sounding: 10347.899, Volume: 10347.899, Density: 1.025},
	}
}

func getInitDraftData() types.InitialDraft {
	return types.InitialDraft{
		TPCListPort:      49.665,
		TPCListStarboard: 49.688,
		Density:          1.023,
	}
}

func getMarks() types.Marks {
	return types.Marks{
		FwdPort:      types.Mark{Value: 3.41},
		FwdStarboard: types.Mark{Value: 3.41},
		MidPort:      types.Mark{Value: 4.51},
		MidStarboard: types.Mark{Value: 4.54},
		AftPort:      types.Mark{Value: 5.69},
		AftStarboard: types.Mark{Value: 5.70},
	}
}

func getVesselData() vessel.VesselData {
	return vessel.VesselData{
		DistancePPFwd:  1.400,
		DistancePPMid:  0.400,
		DistancePPAft:  9.950,
		PPFwdDirection: vessel.PPDirectionAft,
		PPMidDirection: vessel.PPDirectionAft,
		PPAftDirection: vessel.PPDirectionForward,
		LBP:            182.000,
		KeelFwd:        0.000,
		KeelMid:        0.000,
		KeelAft:        0.000,
		VesselType:     vessel.VesselTypeMarine,
	}
}

func getInitHydrostaticRows() []types.HydrostaticRow {
	return []types.HydrostaticRow{
		{Draft: 4.54, Displacement: 21226, TPC: 49.7, LCF: 6.93, LCFDirection: types.LCFDirectionForward},
		{Draft: 4.55, Displacement: 21276, TPC: 49.7, LCF: 6.92, LCFDirection: types.LCFDirectionForward},
	}
}

func getInitMtcRows() []types.MTCRow {
	return []types.MTCRow{
		{Draft: 4.04, MTC: 529.4},
		{Draft: 5.04, MTC: 548.0},
	}
}

func getInitDeductibles() types.Deductibles {
	return types.Deductibles{
		HFO: 683.868,
		MDO: 89.130,
	}
}

func TestFreshWaterTank_GetWeight(t *testing.T) {
	const weight = 3.5
	tank := getFreshWaterTank()
	tankWeight := tank.GetWeight()
	if tankWeight != weight {
		t.Errorf("Expected %f, got %f", weight, tankWeight)
	}
}

func TestBallastWaterTank_GetWeight(t *testing.T) {
	const weight = 3.587
	tank := getBallastWaterTank()
	tankWeight := round3(tank.GetWeight())
	if tankWeight != weight {
		t.Errorf("Expected %f, got %f", weight, tankWeight)
	}
}

func TestTotalFreshWater(t *testing.T) {
	const weight = 364.000
	tanks := getInitFreshWaterTanks()

	totalWeight := TotalFreshWater(tanks)
	if totalWeight != weight {
		t.Errorf("Expected %f, got %f", weight, totalWeight)
	}
}

func TestTotalBallastWater(t *testing.T) {
	const weight = 17.935
	tank := getBallastWaterTank()
	var ballastWaterTanks []types.BallastWaterTank

	for i := 0; i < 5; i++ {
		ballastWaterTanks = append(ballastWaterTanks, tank)
	}
	totalWeight := round3(TotalBallastWater(ballastWaterTanks))
	if totalWeight != weight {
		t.Errorf("Expected %f, got %f", weight, totalWeight)
	}
}

func TestMeanDrafts(t *testing.T) {
	draftFWDExpected := 3.410
	draftMIDExpected := 4.525
	draftAFTExpected := 5.695

	marks := getMarks()
	meanDrafts := MeanDrafts(marks)

	if meanDrafts.DraftFwdMean != draftFWDExpected {
		t.Errorf("Expected %f, got %f", draftFWDExpected, meanDrafts.DraftFwdMean)
	}
	if meanDrafts.DraftMidMean != draftMIDExpected {
		t.Errorf("Expected %f, got %f", draftMIDExpected, meanDrafts.DraftMidMean)
	}
	if meanDrafts.DraftAftMean != draftAFTExpected {
		t.Errorf("Expected %f, got %f", draftAFTExpected, meanDrafts.DraftAftMean)
	}
}

func TestCalcFullLBPPPCorrections(t *testing.T) {
	fwdCorrectionExpected := -0.019
	midCorrectionExpected := -0.005
	aftCorrectionExpected := 0.133

	meanDrafts := MeanDrafts(getMarks())
	vesselData := getVesselData()
	ppCorrections := CalcFullLBPPPCorrections(meanDrafts, vesselData)

	if fwdCorrectionExpected != ppCorrections.FwdCorrection {
		t.Errorf("Expected %f, got %f", fwdCorrectionExpected, ppCorrections.FwdCorrection)
	}
	if midCorrectionExpected != ppCorrections.MidCorrection {
		t.Errorf("Expected %f, got %f", midCorrectionExpected, ppCorrections.MidCorrection)
	}
	if aftCorrectionExpected != ppCorrections.AftCorrection {
		t.Errorf("Expected %f, got %f", aftCorrectionExpected, ppCorrections.AftCorrection)
	}
}

func TestCalcHalfLBPPPCorrections(t *testing.T) {
	fwdCorrectionExpected := -0.017
	midCorrectionExpected := -0.005
	aftCorrectionExpected := 0.144

	meanDrafts := MeanDrafts(getMarks())
	vesselData := getVesselData()
	ppCorrections := CalcHalfLBPPPCorrections(meanDrafts, vesselData)

	if fwdCorrectionExpected != ppCorrections.FwdCorrection {
		t.Errorf("Expected %f, got %f", fwdCorrectionExpected, ppCorrections.FwdCorrection)
	}
	if midCorrectionExpected != ppCorrections.MidCorrection {
		t.Errorf("Expected %f, got %f", midCorrectionExpected, ppCorrections.MidCorrection)
	}
	if aftCorrectionExpected != ppCorrections.AftCorrection {
		t.Errorf("Expected %f, got %f", aftCorrectionExpected, ppCorrections.AftCorrection)
	}
}

func TestCalcDraftsWKeel(t *testing.T) {
	FWDDraftsWKeelExpected := 3.391
	MIDDraftsWKeelExpected := 4.520
	AFTDraftsWKeelExpected := 5.828

	marks := getMarks()
	meanDraft := MeanDrafts(marks)
	vesselData := getVesselData()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vesselData)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vesselData)

	if FWDDraftsWKeelExpected != draftsWKeel.FwdDraftWKeel {
		t.Errorf("Expected %f, got %f", FWDDraftsWKeelExpected, draftsWKeel.FwdDraftWKeel)
	}
	if MIDDraftsWKeelExpected != draftsWKeel.MidDraftWKeel {
		t.Errorf("Expected %f, got %f", MIDDraftsWKeelExpected, draftsWKeel.MidDraftWKeel)
	}
	if AFTDraftsWKeelExpected != draftsWKeel.AftDraftWKeel {
		t.Errorf("Expected %f, got %f", AFTDraftsWKeelExpected, draftsWKeel.AftDraftWKeel)
	}
}

func TestCalcMMC(t *testing.T) {
	mmcExpected := 4.542
	marks := getMarks()
	meanDraft := MeanDrafts(marks)
	vesselData := getVesselData()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vesselData)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vesselData)
	mmc := CalcMMC(draftsWKeel, vesselData)

	if mmcExpected != mmc {
		t.Errorf("Expected %f, got %f", mmcExpected, mmc)
	}
}

func TestInterpolate(t *testing.T) {
	expected := 21236.000
	got := Interpolate(4.542, 4.540, 21226.000, 4.550, 21276.000)

	if expected != got {
		t.Errorf("Expected %f, got %f", expected, got)
	}
}

func TestCalcHydrostatics(t *testing.T) {
	displacementExpected := 21236.000
	tpcExpected := 49.700
	lcfExpected := -6.928
	marks := getMarks()
	meanDraft := MeanDrafts(marks)
	vesselData := getVesselData()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vesselData)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vesselData)
	mmc := CalcMMC(draftsWKeel, vesselData)
	hr := getInitHydrostaticRows()
	hydrostatics := CalcHydrostatics(mmc, hr, vesselData)

	if displacementExpected != hydrostatics.Displacement {
		t.Errorf("Expected %f, got %f", displacementExpected, hydrostatics.Displacement)
	}
	if tpcExpected != hydrostatics.TPC {
		t.Errorf("Expected %f, got %f", tpcExpected, hydrostatics.TPC)
	}
	if lcfExpected != hydrostatics.LCF {
		t.Errorf("Expected %f, got %f", lcfExpected, hydrostatics.LCF)
	}
}

func TestCalcFirstTrimCorrection(t *testing.T) {
	firstTrimCorrectionExpected := -461.050
	marks := getMarks()
	meanDraft := MeanDrafts(marks)
	vesselData := getVesselData()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vesselData)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vesselData)
	mmc := CalcMMC(draftsWKeel, vesselData)
	hr := getInitHydrostaticRows()
	hydrostatics := CalcHydrostatics(mmc, hr, vesselData)
	firstTrimCorrectionGot := CalcFirstTrimCorrection(draftsWKeel, hydrostatics.TPC, hydrostatics.LCF, vesselData.LBP)

	if firstTrimCorrectionExpected != firstTrimCorrectionGot {
		t.Errorf("Expected %f, got %f", firstTrimCorrectionExpected, firstTrimCorrectionGot)
	}
}

func TestCalcSecondTrimCorrection(t *testing.T) {
	secondTrimCorrectionExpected := 30.347
	marks := getMarks()
	meanDraft := MeanDrafts(marks)
	vesselData := getVesselData()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vesselData)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vesselData)
	mtcRows := getInitMtcRows()
	secondTrimCorrectionGot := CalcSecondTrimCorrection(draftsWKeel, mtcRows, vesselData.LBP)

	if secondTrimCorrectionExpected != secondTrimCorrectionGot {
		t.Errorf("Expected %f, got %f", secondTrimCorrectionExpected, secondTrimCorrectionGot)
	}
}

func TestCalcListCorrection(t *testing.T) {
	listCorrectionExpected := 0.004
	marks := getMarks()
	initDS := getInitDraftData()
	listCorrectionGot := CalcListCorrection(marks, initDS.TPCListPort, initDS.TPCListStarboard)
	if listCorrectionExpected != listCorrectionGot {
		t.Errorf("Expected %f, got %f", listCorrectionExpected, listCorrectionGot)
	}
}

func TestCalcDensityCorrection(t *testing.T) {
	densityCorrExpected := -40.596
	marks := getMarks()
	meanDraft := MeanDrafts(marks)
	vesselData := getVesselData()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vesselData)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vesselData)
	mmc := CalcMMC(draftsWKeel, vesselData)
	hr := getInitHydrostaticRows()
	hydrostatics := CalcHydrostatics(mmc, hr, vesselData)
	mtcRows := getInitMtcRows()
	initDS := getInitDraftData()
	firstTrim := CalcFirstTrimCorrection(draftsWKeel, hydrostatics.TPC, hydrostatics.LCF, vesselData.LBP)
	secondTrim := CalcSecondTrimCorrection(draftsWKeel, mtcRows, vesselData.LBP)
	listCorrection := CalcListCorrection(marks, initDS.TPCListPort, initDS.TPCListStarboard)
	densityCorrGot := CalcDensityCorrection(hydrostatics.Displacement, firstTrim, secondTrim, listCorrection, initDS.Density)

	if densityCorrExpected != densityCorrGot {
		t.Errorf("Expected %f, got %f", densityCorrExpected, densityCorrGot)
	}
}

func TestCalcTotalDeductibles(t *testing.T) {
	totalDeductiblesExpected := 11743.594
	bwt := getInitBallastWaterTanks()
	fwt := getInitFreshWaterTanks()
	d := getInitDeductibles()
	totalDeductiblesGot := CalcTotalDeductibles(bwt, fwt, d)

	if totalDeductiblesExpected != totalDeductiblesGot {
		t.Errorf("Expected %f, got %f", totalDeductiblesExpected, totalDeductiblesGot)
	}
}

func TestCalcNetDisplacement(t *testing.T) {
	netDisplacementExpected := 9021.111
	bwt := getInitBallastWaterTanks()
	fwt := getInitFreshWaterTanks()
	d := getInitDeductibles()
	totalDeductibles := CalcTotalDeductibles(bwt, fwt, d)
	marks := getMarks()
	meanDraft := MeanDrafts(marks)
	vesselData := getVesselData()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vesselData)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vesselData)
	mmc := CalcMMC(draftsWKeel, vesselData)
	hr := getInitHydrostaticRows()
	hydrostatics := CalcHydrostatics(mmc, hr, vesselData)
	mtcRows := getInitMtcRows()
	initDS := getInitDraftData()
	firstTrim := CalcFirstTrimCorrection(draftsWKeel, hydrostatics.TPC, hydrostatics.LCF, vesselData.LBP)
	secondTrim := CalcSecondTrimCorrection(draftsWKeel, mtcRows, vesselData.LBP)
	listCorrection := CalcListCorrection(marks, initDS.TPCListPort, initDS.TPCListStarboard)
	densityCorr := CalcDensityCorrection(hydrostatics.Displacement, firstTrim, secondTrim, listCorrection, initDS.Density)
	netDisplacementGot := CalcNetDisplacement(hydrostatics.Displacement, firstTrim, secondTrim, listCorrection, densityCorr, totalDeductibles)
	if netDisplacementExpected != netDisplacementGot {
		t.Errorf("Expected %f, got %f", netDisplacementExpected, netDisplacementGot)
	}
}

func TestCalcCargoWeight(t *testing.T) {
	netDisplacementIni := 9000.000
	netDisplacementFin := 49000.000
	cargoWeightExpected := 40000.000
	cargoWeightGot := CalcCargoWeight(netDisplacementIni, netDisplacementFin)
	if cargoWeightExpected != cargoWeightGot {
		t.Errorf("Expected %f, got %f", cargoWeightExpected, cargoWeightGot)
	}
}
