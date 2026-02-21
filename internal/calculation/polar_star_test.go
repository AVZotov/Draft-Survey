package calculation

import "testing"

// POLAR STAR - TrimnoList tests

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
		FwdPort: 3.33, FwdStarboard: 3.33,
		MidPort: 4.64, MidStarboard: 4.64,
		AftPort: 6.12, AftStarboard: 6.12,
	}
}

func getPolarStarTrimnoListHydrostaticRows() []HydrostaticRow {
	return []HydrostaticRow{
		{Draft: 4.617, Displacement: 19182.7, TPC: 45.2, LCF: 98.457, LCFDirection: LCFDirectionFromAP},
		{Draft: 4.667, Displacement: 19409.0, TPC: 45.3, LCF: 98.405, LCFDirection: LCFDirectionFromAP},
	}
}

func getPolarStarTrimnoListMTCRows() []MTCRow {
	return []MTCRow{
		{Draft: 4.167, MTC: 500.2},
		{Draft: 5.167, MTC: 526.9},
	}
}

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

func TestPolarStar_TrimnoListCalcHydrostatics(t *testing.T) {
	displasementExpected := 19304.902
	tpcExpected := 45.254
	lcfExpected := -6.929
	marks := getPolarStarTrimnoListMarks()
	vessel := getPolarStarTrimnoListVessel()
	meanDraft := MeanDrafts(marks)
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vessel)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vessel)
	mmc := CalcMMC(draftsWKeel, vessel.VesselType)
	hr := getPolarStarTrimnoListHydrostaticRows()

	hydrostatics := CalcHydrostatics(mmc, hr, vessel)

	if displasementExpected != hydrostatics.Displacement {
		t.Errorf("Expected %f, got %f", displasementExpected, hydrostatics.Displacement)
	}
	if tpcExpected != hydrostatics.TPC {
		t.Errorf("Expected %f, got %f", tpcExpected, hydrostatics.TPC)
	}
	if lcfExpected != hydrostatics.LCF {
		t.Errorf("Expected %f, got %f", lcfExpected, hydrostatics.LCF)
	}
}

func TestPolarStar_TrimnoList_FirstTrimCorrection(t *testing.T) {
	marks := getPolarStarTrimnoListMarks()
	vessel := getPolarStarTrimnoListVessel()
	meanDraft := MeanDrafts(marks)
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vessel)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vessel)
	mmc := CalcMMC(draftsWKeel, vessel.VesselType)
	hr := getPolarStarTrimnoListHydrostaticRows()
	hydrostatics := CalcHydrostatics(mmc, hr, vessel)
	got := CalcFirstTrimCorrection(draftsWKeel, hydrostatics.TPC, hydrostatics.LCF, vessel.LBP)

	if got != -487.653 {
		t.Errorf("1st trim: expected -487.653, got %f", got)
	}
}

func TestPolarStar_TrimnoList_SecondTrimCorrection(t *testing.T) {
	marks := getPolarStarTrimnoListMarks()
	vessel := getPolarStarTrimnoListVessel()
	meanDraft := MeanDrafts(marks)
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vessel)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vessel)
	mtcRows := getPolarStarTrimnoListMTCRows()
	got := CalcSecondTrimCorrection(draftsWKeel, mtcRows, vessel.LBP)

	if got != 59.088 {
		t.Errorf("2nd trim: expected 59.088, got %f", got)
	}
}

func TestPolarStar_TrimnoList_ListCorrection(t *testing.T) {
	marks := getPolarStarTrimnoListMarks()
	got := CalcListCorrection(marks, 0, 0)

	if got != 0 {
		t.Errorf("List corr: expected 0, got %f", got)
	}
}

func TestPolarStar_TrimnoList_DensityCorrection(t *testing.T) {
	marks := getPolarStarTrimnoListMarks()
	vessel := getPolarStarTrimnoListVessel()
	meanDraft := MeanDrafts(marks)
	ppCorrections := CalcFullLBPPPCorrections(meanDraft, vessel)
	draftsWKeel := CalcDraftsWKeel(meanDraft, ppCorrections, vessel)
	mmc := CalcMMC(draftsWKeel, vessel.VesselType)
	hr := getPolarStarTrimnoListHydrostaticRows()
	hydrostatics := CalcHydrostatics(mmc, hr, vessel)
	mtcRows := getPolarStarTrimnoListMTCRows()
	firstTrim := CalcFirstTrimCorrection(draftsWKeel, hydrostatics.TPC, hydrostatics.LCF, vessel.LBP)
	secondTrim := CalcSecondTrimCorrection(draftsWKeel, mtcRows, vessel.LBP)
	listCorrection := CalcListCorrection(marks, 0, 0)
	got := CalcDensityCorrection(hydrostatics.Displacement, firstTrim, secondTrim, listCorrection, 1.017)

	if got != -147.328 {
		t.Errorf("Density corr: expected -147.328, got %f", got)
	}
}

// POLAR STAR - TrimList tests

func getPolarStarTrimListVessel() Vessel {
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

func getPolarStarTrimListMarks() Marks {
	return Marks{
		FwdPort: 3.39, FwdStarboard: 3.36,
		MidPort: 4.64, MidStarboard: 4.54,
		AftPort: 6.12, AftStarboard: 6.12,
	}
}

func getPolarStarTrimListHydrostaticRows() []HydrostaticRow {
	return []HydrostaticRow{
		{Draft: 4.567, Displacement: 18956.7, TPC: 45.2, LCF: 98.509, LCFDirection: LCFDirectionFromAP},
		{Draft: 4.617, Displacement: 19182.7, TPC: 45.2, LCF: 98.457, LCFDirection: LCFDirectionFromAP},
	}
}

func getPolarStarTrimListMTCRows() []MTCRow {
	return []MTCRow{
		{Draft: 4.117, MTC: 498.8},
		{Draft: 5.117, MTC: 525.7},
	}
}

func TestPolarStar_TrimList_MeanDrafts(t *testing.T) {
	got := MeanDrafts(getPolarStarTrimListMarks())
	if got.DraftFWDmean != 3.375 {
		t.Errorf("meanF: expected 3.375, got %f", got.DraftFWDmean)
	}
	if got.DraftMIDmean != 4.590 {
		t.Errorf("meanM: expected 4.590, got %f", got.DraftMIDmean)
	}
	if got.DraftAFTmean != 6.120 {
		t.Errorf("meanA: expected 6.120, got %f", got.DraftAFTmean)
	}
}

func TestPolarStar_TrimList_PPCorrections(t *testing.T) {
	vessel := getPolarStarTrimListVessel()
	got := CalcFullLBPPPCorrections(MeanDrafts(getPolarStarTrimListMarks()), vessel)
	if got.FWDCorrection != -0.073 {
		t.Errorf("FWD: expected -0.073, got %f", got.FWDCorrection)
	}
	if got.MIDCorrection != -0.008 {
		t.Errorf("MID: expected -0.008, got %f", got.MIDCorrection)
	}
	if got.AFTCorrection != -0.018 {
		t.Errorf("AFT: expected -0.018, got %f", got.AFTCorrection)
	}
}

func TestPolarStar_TrimList_DraftsWKeel(t *testing.T) {
	vessel := getPolarStarTrimListVessel()
	meanDraft := MeanDrafts(getPolarStarTrimListMarks())
	got := CalcDraftsWKeel(meanDraft, CalcFullLBPPPCorrections(meanDraft, vessel), vessel)
	if got.FWDDraftWKeel != 3.302 {
		t.Errorf("FWD wKeel: expected 3.302, got %f", got.FWDDraftWKeel)
	}
	if got.MIDDraftWKeel != 4.582 {
		t.Errorf("MID wKeel: expected 4.582, got %f", got.MIDDraftWKeel)
	}
	if got.AFTDraftWKeel != 6.102 {
		t.Errorf("AFT wKeel: expected 6.102, got %f", got.AFTDraftWKeel)
	}
}

func TestPolarStar_TrimList_MMC(t *testing.T) {
	vessel := getPolarStarTrimListVessel()
	meanDraft := MeanDrafts(getPolarStarTrimListMarks())
	draftsWKeel := CalcDraftsWKeel(meanDraft, CalcFullLBPPPCorrections(meanDraft, vessel), vessel)
	got := CalcMMC(draftsWKeel, vessel.VesselType)
	if got != 4.612 {
		t.Errorf("MMC: expected 4.612, got %f", got)
	}
}

func TestPolarStar_TrimList_Hydrostatics(t *testing.T) {
	vessel := getPolarStarTrimListVessel()
	meanDraft := MeanDrafts(getPolarStarTrimListMarks())
	draftsWKeel := CalcDraftsWKeel(meanDraft, CalcFullLBPPPCorrections(meanDraft, vessel), vessel)
	mmc := CalcMMC(draftsWKeel, vessel.VesselType)
	hydrostatics := CalcHydrostatics(mmc, getPolarStarTrimListHydrostaticRows(), vessel)
	if hydrostatics.Displacement != 19160.1 {
		t.Errorf("Displacement: expected 19160.1, got %f", hydrostatics.Displacement)
	}
	if hydrostatics.TPC != 45.2 {
		t.Errorf("TPC: expected 45.2, got %f", hydrostatics.TPC)
	}
	if hydrostatics.LCF != -6.962 {
		t.Errorf("LCF: expected -6.962, got %f", hydrostatics.LCF)
	}
}

func TestPolarStar_TrimList_FirstTrimCorrection(t *testing.T) {
	vessel := getPolarStarTrimListVessel()
	meanDraft := MeanDrafts(getPolarStarTrimListMarks())
	draftsWKeel := CalcDraftsWKeel(meanDraft, CalcFullLBPPPCorrections(meanDraft, vessel), vessel)
	mmc := CalcMMC(draftsWKeel, vessel.VesselType)
	hydrostatics := CalcHydrostatics(mmc, getPolarStarTrimListHydrostaticRows(), vessel)
	got := CalcFirstTrimCorrection(draftsWKeel, hydrostatics.TPC, hydrostatics.LCF, vessel.LBP)
	if got != -481.481 {
		t.Errorf("1st trim: expected -481.481, got %f", got)
	}
}

func TestPolarStar_TrimList_SecondTrimCorrection(t *testing.T) {
	vessel := getPolarStarTrimListVessel()
	meanDraft := MeanDrafts(getPolarStarTrimListMarks())
	draftsWKeel := CalcDraftsWKeel(meanDraft, CalcFullLBPPPCorrections(meanDraft, vessel), vessel)
	got := CalcSecondTrimCorrection(draftsWKeel, getPolarStarTrimListMTCRows(), vessel.LBP)
	if got != 57.622 {
		t.Errorf("2nd trim: expected 57.622, got %f", got)
	}
}

func TestPolarStar_TrimList_ListCorrection(t *testing.T) {
	got := CalcListCorrection(getPolarStarTrimListMarks(), 45.212, 45.129)
	if got != 0.05 {
		t.Errorf("List corr: expected 0.050, got %f", got)
	}
}

func TestPolarStar_TrimList_DensityCorrection(t *testing.T) {
	vessel := getPolarStarTrimListVessel()
	marks := getPolarStarTrimListMarks()
	meanDraft := MeanDrafts(marks)
	draftsWKeel := CalcDraftsWKeel(meanDraft, CalcFullLBPPPCorrections(meanDraft, vessel), vessel)
	mmc := CalcMMC(draftsWKeel, vessel.VesselType)
	hydrostatics := CalcHydrostatics(mmc, getPolarStarTrimListHydrostaticRows(), vessel)
	firstTrim := CalcFirstTrimCorrection(draftsWKeel, hydrostatics.TPC, hydrostatics.LCF, vessel.LBP)
	secondTrim := CalcSecondTrimCorrection(draftsWKeel, getPolarStarTrimListMTCRows(), vessel.LBP)
	listCorr := CalcListCorrection(marks, 45.212, 45.129)
	got := CalcDensityCorrection(hydrostatics.Displacement, firstTrim, secondTrim, listCorr, 1.017)
	if got != -146.234 {
		t.Errorf("Density corr: expected -146.234, got %f", got)
	}
}
