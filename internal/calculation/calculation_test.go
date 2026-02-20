package calculation

import (
	"testing"
)

func getFreshWaterTank() FreshWaterTank {
	return FreshWaterTank{
		Name:     "test",
		Sounding: 3.5,
		Volume:   3.5,
	}
}

func getInitFreshWaterTanks() []FreshWaterTank {
	return []FreshWaterTank{
		{Name: "FW P", Sounding: 364.000, Volume: 364.000},
	}
}

func getBallastWaterTank() BallastWaterTank {
	return BallastWaterTank{
		Name:     "test",
		Sounding: 3.5,
		Volume:   3.5,
		Density:  1.025,
	}
}

func getInitBallastWaterTanks() []BallastWaterTank {
	return []BallastWaterTank{
		{Name: "FPT", Sounding: 10347.899, Volume: 10347.899, Density: 1.025},
	}
}

func getInitDraftData() InitialDraft {
	return InitialDraft{
		TPCListPort:      49.665,
		TPCListStarboard: 49.688,
		Density:          1.023,
	}
}

func getMarks() Marks {
	return Marks{
		FWDPort:      3.41,
		FWDStarboard: 3.41,
		MIDPort:      4.51,
		MIDStarboard: 4.54,
		AFTPort:      5.69,
		AFTStarboard: 5.70,
	}
}

func getVessel() Vessel {
	return Vessel{
		DistancePPFWD:  1.400,
		DistancePPMID:  0.400,
		DistancePPAFT:  9.950,
		PPFWDDirection: PPDirectionAft,
		PPMIDDirection: PPDirectionAft,
		PPAFTDirection: PPDirectionForward,
		LBP:            182.000,
		KeelFWD:        0.000,
		KeelMID:        0.000,
		KeelAFT:        0.000,
		VesselType:     VesselTypeMarine,
	}
}

func getInitHydrostaticRows() []HydrostaticRow {
	return []HydrostaticRow{
		{Draft: 4.54, Displacement: 21226, TPC: 49.7, LCF: 6.93, LCFDirection: LCFDirectionForward},
		{Draft: 4.55, Displacement: 21276, TPC: 49.7, LCF: 6.92, LCFDirection: LCFDirectionForward},
	}
}

func getInitMtcRows() []MTCRow {
	return []MTCRow{
		{Draft: 4.04, MTC: 529.4},
		{Draft: 5.04, MTC: 548.0},
	}
}

func getInitDeductibles() Deductibles {
	return Deductibles{
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
	var ballastWaterTanks []BallastWaterTank

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

	if meanDrafts.DraftFWDmean != draftFWDExpected {
		t.Errorf("Expected %f, got %f", draftFWDExpected, meanDrafts.DraftFWDmean)
	}
	if meanDrafts.DraftMIDmean != draftMIDExpected {
		t.Errorf("Expected %f, got %f", draftMIDExpected, meanDrafts.DraftMIDmean)
	}
	if meanDrafts.DraftAFTmean != draftAFTExpected {
		t.Errorf("Expected %f, got %f", draftAFTExpected, meanDrafts.DraftAFTmean)
	}
}

func TestCalcPPCorrections(t *testing.T) {
	fwdCorrectionExpected := -0.019
	midCorrectionExpected := -0.005
	aftCorrectionExpected := 0.133

	meanDrafts := MeanDrafts(getMarks())
	vessel := getVessel()
	ppCorrections := CalcFullLBPPPCorrections(meanDrafts, vessel)

	if fwdCorrectionExpected != ppCorrections.FWDCorrection {
		t.Errorf("Expected %f, got %f", fwdCorrectionExpected, ppCorrections.FWDCorrection)
	}
	if midCorrectionExpected != ppCorrections.MIDCorrection {
		t.Errorf("Expected %f, got %f", midCorrectionExpected, ppCorrections.MIDCorrection)
	}
	if aftCorrectionExpected != ppCorrections.AFTCorrection {
		t.Errorf("Expected %f, got %f", aftCorrectionExpected, ppCorrections.AFTCorrection)
	}
}

func TestCalcDraftsWKeel(t *testing.T) {
	FWDDraftsWKeelExpected := 3.391
	MIDDraftsWKeelExpected := 4.520
	AFTDraftsWKeelExpected := 5.828

	marks := getMarks()
	meanDraft := MeanDrafts(marks)
	vessel := getVessel()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vessel)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vessel)

	if FWDDraftsWKeelExpected != draftsWKeel.FWDDraftWKeel {
		t.Errorf("Expected %f, got %f", FWDDraftsWKeelExpected, draftsWKeel.FWDDraftWKeel)
	}
	if MIDDraftsWKeelExpected != draftsWKeel.MIDDraftWKeel {
		t.Errorf("Expected %f, got %f", MIDDraftsWKeelExpected, draftsWKeel.MIDDraftWKeel)
	}
	if AFTDraftsWKeelExpected != draftsWKeel.AFTDraftWKeel {
		t.Errorf("Expected %f, got %f", AFTDraftsWKeelExpected, draftsWKeel.AFTDraftWKeel)
	}
}

func TestCalcMMC(t *testing.T) {
	MMCExpected := 4.542
	marks := getMarks()
	meanDraft := MeanDrafts(marks)
	vessel := getVessel()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vessel)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vessel)
	MMC := CalcMMC(draftsWKeel, vessel.VesselType)

	if MMCExpected != MMC {
		t.Errorf("Expected %f, got %f", MMCExpected, MMC)
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
	displasementExpected := 21236.000
	tpcExpected := 49.700
	lcfExpected := -6.928
	marks := getMarks()
	meanDraft := MeanDrafts(marks)
	vessel := getVessel()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vessel)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vessel)
	mmc := CalcMMC(draftsWKeel, vessel.VesselType)
	hr := getInitHydrostaticRows()
	displasementGot, tpcGot, lcfGot := CalcHydrostatics(mmc, hr)

	if displasementExpected != displasementGot {
		t.Errorf("Expected %f, got %f", displasementExpected, displasementGot)
	}
	if tpcExpected != tpcGot {
		t.Errorf("Expected %f, got %f", tpcExpected, tpcGot)
	}
	if lcfExpected != lcfGot {
		t.Errorf("Expected %f, got %f", lcfExpected, lcfGot)
	}
}

func TestCalcFirstTrimCorrection(t *testing.T) {
	firstTrimCorrectionExpected := -461.050
	marks := getMarks()
	meanDraft := MeanDrafts(marks)
	vessel := getVessel()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vessel)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vessel)
	mmc := CalcMMC(draftsWKeel, vessel.VesselType)
	hr := getInitHydrostaticRows()
	_, tpc, lcf := CalcHydrostatics(mmc, hr)
	firstTrimCorrectionGot := CalcFirstTrimCorrection(draftsWKeel, tpc, lcf, vessel.LBP)

	if firstTrimCorrectionExpected != firstTrimCorrectionGot {
		t.Errorf("Expected %f, got %f", firstTrimCorrectionExpected, firstTrimCorrectionGot)
	}
}

func TestCalcSecondTrimCorrection(t *testing.T) {
	secondTrimCorrectionExpected := 30.347
	marks := getMarks()
	meanDraft := MeanDrafts(marks)
	vessel := getVessel()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vessel)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vessel)
	mtcRows := getInitMtcRows()
	secondTrimCorrectionGot := CalcSecondTrimCorrection(draftsWKeel, mtcRows, vessel.LBP)

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
	vessel := getVessel()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vessel)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vessel)
	mmc := CalcMMC(draftsWKeel, vessel.VesselType)
	hr := getInitHydrostaticRows()
	displacement, tpc, lcf := CalcHydrostatics(mmc, hr)
	mtcRows := getInitMtcRows()
	initDS := getInitDraftData()
	firstTrim := CalcFirstTrimCorrection(draftsWKeel, tpc, lcf, vessel.LBP)
	secondTrim := CalcSecondTrimCorrection(draftsWKeel, mtcRows, vessel.LBP)
	listCorrection := CalcListCorrection(marks, initDS.TPCListPort, initDS.TPCListStarboard)
	densityCorrGot := CalcDensityCorrection(displacement, firstTrim, secondTrim, listCorrection, initDS.Density)

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
	vessel := getVessel()
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vessel)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vessel)
	mmc := CalcMMC(draftsWKeel, vessel.VesselType)
	hr := getInitHydrostaticRows()
	displacement, tpc, lcf := CalcHydrostatics(mmc, hr)
	mtcRows := getInitMtcRows()
	initDS := getInitDraftData()
	firstTrim := CalcFirstTrimCorrection(draftsWKeel, tpc, lcf, vessel.LBP)
	secondTrim := CalcSecondTrimCorrection(draftsWKeel, mtcRows, vessel.LBP)
	listCorrection := CalcListCorrection(marks, initDS.TPCListPort, initDS.TPCListStarboard)
	densityCorr := CalcDensityCorrection(displacement, firstTrim, secondTrim, listCorrection, initDS.Density)
	netDisplacementGot := CalcNetDisplacement(displacement, firstTrim, secondTrim, listCorrection, densityCorr, totalDeductibles)
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
