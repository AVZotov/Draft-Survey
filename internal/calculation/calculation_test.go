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
		FWDPort:      7.1,
		FWDStarboard: 7.2,
		MIDPort:      7.3,
		MIDStarboard: 7.4,
		AFTPort:      7.5,
		AFTStarboard: 7.6,
	}
}

func getVessel() Vessel {
	return Vessel{
		DistancePPFWD: 1.400,
		DistancePPMID: 0.400,
		DistancePPAFT: 9.950,
		LBP:           182.000,
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
	draftFWDExpected := 7.15
	draftMIDExpected := 7.35
	draftAFTExpected := 7.55

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
	fwdCorrectionExpected := 0.002
	aftCorrectionExpected := -0.011

	meanDrafts := MeanDrafts(getMarks())
	vessel := getVessel()
	ppCorrections := CalcPPCorrections(meanDrafts, vessel)

	if fwdCorrectionExpected != ppCorrections.FWDCorrection {
		t.Errorf("Expected %f, got %f", fwdCorrectionExpected, ppCorrections.FWDCorrection)
	}
	if aftCorrectionExpected != ppCorrections.AFTCorrection {
		t.Errorf("Expected %f, got %f", aftCorrectionExpected, ppCorrections.AFTCorrection)
	}
}
