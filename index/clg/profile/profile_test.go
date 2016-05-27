package profile

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

// testMaybeNewProfiles returns two different profiles. We use them below to
// test different profile methods.
func testMaybeNewProfiles(t *testing.T) (spec.CLGProfile, spec.CLGProfile) {
	firstConfig := DefaultConfig()
	firstConfig.Body = "firstBody"
	firstConfig.HasChanged = true
	firstConfig.Hash = "firstHash"
	firstConfig.InputsOutputs = spec.InputsOutputs{
		InsOuts: []spec.InOut{
			{
				In:  []string{"firstIn"},
				Out: []string{"firstOut"},
			},
		},
	}
	firstConfig.Name = "firstName"
	firstProfile, err := New(firstConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	secondConfig := DefaultConfig()
	secondConfig.Body = "secondBody"
	secondConfig.HasChanged = false
	secondConfig.Hash = "secondHash"
	secondConfig.InputsOutputs = spec.InputsOutputs{
		InsOuts: []spec.InOut{
			{
				In:  []string{"secondIn"},
				Out: []string{"secondOut"},
			},
		},
	}
	secondConfig.Name = "secondName"
	secondProfile, err := New(secondConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return firstProfile, secondProfile
}

func Test_Profile_Profile_Equals(t *testing.T) {
	firstProfile, secondProfile := testMaybeNewProfiles(t)

	if firstProfile.Equals(secondProfile) {
		t.Fatal("expected", false, "got", true)
	}
	if !firstProfile.Equals(firstProfile) {
		t.Fatal("expected", true, "got", false)
	}
	if !secondProfile.Equals(secondProfile) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Profile_Profile_GetBody(t *testing.T) {
	firstProfile, secondProfile := testMaybeNewProfiles(t)

	if firstProfile.GetBody() == secondProfile.GetBody() {
		t.Fatal("expected", false, "got", true)
	}
	if firstProfile.GetBody() != firstProfile.GetBody() {
		t.Fatal("expected", true, "got", false)
	}
	if secondProfile.GetBody() != secondProfile.GetBody() {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Profile_Profile_GetHasChanged(t *testing.T) {
	firstProfile, secondProfile := testMaybeNewProfiles(t)

	if firstProfile.GetHasChanged() == secondProfile.GetHasChanged() {
		t.Fatal("expected", false, "got", true)
	}
	if firstProfile.GetHasChanged() != firstProfile.GetHasChanged() {
		t.Fatal("expected", true, "got", false)
	}
	if secondProfile.GetHasChanged() != secondProfile.GetHasChanged() {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Profile_Profile_GetHash(t *testing.T) {
	firstProfile, secondProfile := testMaybeNewProfiles(t)

	if firstProfile.GetHash() == secondProfile.GetHash() {
		t.Fatal("expected", false, "got", true)
	}
	if firstProfile.GetHash() != firstProfile.GetHash() {
		t.Fatal("expected", true, "got", false)
	}
	if secondProfile.GetHash() != secondProfile.GetHash() {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Profile_Profile_GetInputsOutputs(t *testing.T) {
	firstProfile, secondProfile := testMaybeNewProfiles(t)

	if reflect.DeepEqual(firstProfile.GetInputsOutputs(), secondProfile.GetInputsOutputs()) {
		t.Fatal("expected", false, "got", true)
	}
	if !reflect.DeepEqual(firstProfile.GetInputsOutputs(), firstProfile.GetInputsOutputs()) {
		t.Fatal("expected", true, "got", false)
	}
	if !reflect.DeepEqual(secondProfile.GetInputsOutputs(), secondProfile.GetInputsOutputs()) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Profile_Profile_GetName(t *testing.T) {
	firstProfile, secondProfile := testMaybeNewProfiles(t)

	if firstProfile.GetName() == secondProfile.GetName() {
		t.Fatal("expected", false, "got", true)
	}
	if firstProfile.GetName() != firstProfile.GetName() {
		t.Fatal("expected", true, "got", false)
	}
	if secondProfile.GetName() != secondProfile.GetName() {
		t.Fatal("expected", true, "got", false)
	}
}
