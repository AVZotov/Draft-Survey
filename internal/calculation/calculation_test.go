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

func getBallastWaterTank() BallastWaterTank {
	return BallastWaterTank{
		Name:     "test",
		Sounding: 3.5,
		Volume:   3.5,
		Density:  1.025,
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
		{Draft: 4.54, Displacement: 21226},
		{Draft: 4.55, Displacement: 21227},
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
	const weight = 17.5
	tank := getFreshWaterTank()
	var freshWaterTanks []FreshWaterTank

	for i := 0; i < 5; i++ {
		freshWaterTanks = append(freshWaterTanks, tank)
	}
	totalWeight := TotalFreshWater(freshWaterTanks)
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
