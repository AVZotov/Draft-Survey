package calculation

import "testing"

func getPolarStarTrimnoListVessel() Vessel {
	return Vessel{
		LBP:            183.000,
		DistancePPFWD:  4.800,
		DistancePPMID:  0.500,
		DistancePPAFT:  1.200,
		PPFWDDirection: PPDirectionAft,
		PPMIDDirection: PPDirectionAft,
		PPAFTDirection: PPDirectionAft,
		KeelFWD:        0,
		KeelMID:        0,
		KeelAFT:        0,
		VesselType:     VesselTypeMarine,
	}
}

func getPolarStarTrimnoListMarks() Marks {
	return Marks{
		FWDPort: 3.33, FWDStarboard: 3.33,
		MIDPort: 4.64, MIDStarboard: 4.64,
		AFTPort: 6.12, AFTStarboard: 6.12,
	}
}

func getPolarStarTrimnoListHydrostaticRows() []HydrostaticRow {
	return []HydrostaticRow{
		{Draft: 4.617, Displacement: 19182.7, TPC: 45.2, LCF: 98.457},
		{Draft: 4.667, Displacement: 19409.0, TPC: 45.3, LCF: 98.405},
	}
}

func getPolarStarTrimnoListMTCRows() []MTCRow {
	return []MTCRow{
		{Draft: 4.167, MTC: 500.2},
		{Draft: 5.167, MTC: 526.9},
	}
}

// POLAR STAR - TrimnoList tests

func TestPolarStar_TrimnoList_MeanDrafts(t *testing.T) {
	marks := getPolarStarTrimnoListMarks()
	got := MeanDrafts(marks)

	if got.DraftFWDmean != 3.330 {
		t.Errorf("meanF: expected 3.330, got %f", got.DraftFWDmean)
	}
	if got.DraftMIDmean != 4.640 {
		t.Errorf("meanM: expected 4.640, got %f", got.DraftMIDmean)
	}
	if got.DraftAFTmean != 6.120 {
		t.Errorf("meanA: expected 6.120, got %f", got.DraftAFTmean)
	}
}

func TestPolarStar_TrimnoList_PPCorrections(t *testing.T) {
	marks := getPolarStarTrimnoListMarks()
	vessel := getPolarStarTrimnoListVessel()
	got := CalcFullLBPPPCorrections(MeanDrafts(marks), vessel)

	if got.FWDCorrection != -0.075 {
		t.Errorf("FWD corr: expected -0.075, got %f", got.FWDCorrection)
	}
	if got.MIDCorrection != -0.008 {
		t.Errorf("MID corr: expected -0.008, got %f", got.MIDCorrection)
	}
	if got.AFTCorrection != -0.019 {
		t.Errorf("AFT corr: expected -0.019, got %f", got.AFTCorrection)
	}
}

func TestPolarStar_TrimnoList_DraftsWKeel(t *testing.T) {
	marks := getPolarStarTrimnoListMarks()
	vessel := getPolarStarTrimnoListVessel()
	meanDraft := MeanDrafts(marks)
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vessel)
	got := CalcDraftsWKeel(meanDraft, ppCorrections, vessel)

	if got.FWDDraftWKeel != 3.255 {
		t.Errorf("FWD wKeel: expected 3.255, got %f", got.FWDDraftWKeel)
	}
	if got.MIDDraftWKeel != 4.632 {
		t.Errorf("MID wKeel: expected 4.632, got %f", got.MIDDraftWKeel)
	}
	if got.AFTDraftWKeel != 6.101 {
		t.Errorf("AFT wKeel: expected 6.101, got %f", got.AFTDraftWKeel)
	}
}

func TestPolarStar_TrimnoList_MMC(t *testing.T) {
	marks := getPolarStarTrimnoListMarks()
	vessel := getPolarStarTrimnoListVessel()
	meanDraft := MeanDrafts(marks)
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vessel)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vessel)
	got := CalcMMC(draftsWKeel, vessel.VesselType)

	if got != 4.644 {
		t.Errorf("MMC: expected 4.644, got %f", got)
	}
}
