package storage

import (
	"reflect"
	"testing"

	"github.com/AVZotov/draft-survey/internal/types"
	"github.com/AVZotov/draft-survey/internal/vessel"
)

const id = "123-456-789"

func getSurvey() *types.Survey {
	return &types.Survey{
		ID:           "123456",
		InitialDraft: types.InitialDraft{},
		FinalDraft:   types.FinalDraft{},
		Job: types.Job{
			JobNumber: 123456,
			DSNumber:  123456,
			Principal: "testPrincipal",
		},
		CargoOperation: types.CargoOperation{
			PlaceOfInspection: "testPlace",
			Destination:       "testDestination",
			Operation:         "testOperation",
			Origin:            "testOrigin",
			Cargo:             "testCargo",
			Packing:           "testPacking",
			Port:              "testPort",
		},
		VesselData: vessel.VesselData{},
	}
}

func TestJSONStore_SaveAndGet(t *testing.T) {
	dir := t.TempDir()
	surveyExpected := getSurvey()
	store := JSONStore{Path: dir, TempPath: dir}
	if err := store.Save(id, surveyExpected); err != nil {
		t.Fatal(err)
	}
	surveyGot, err := store.Get(id)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(surveyExpected, surveyGot) {
		t.Errorf("Expected %v, got %v", surveyExpected, surveyGot)
	}
}

func TestJSONStore_Delete(t *testing.T) {
	dir := t.TempDir()
	surveyExpected := getSurvey()
	store := JSONStore{Path: dir, TempPath: dir}
	if err := store.Save(id, surveyExpected); err != nil {
		t.Fatal(err)
	}

	if err := store.Delete(id); err != nil {
		t.Fatal(err)
	}

	surveyGot, err := store.Get(id)
	if err == nil {
		t.Errorf("Expected error, got %#v", surveyGot)
	}
}
