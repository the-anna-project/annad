package profile

import (
	"testing"
)

func Test_Profile_Profile_GetID(t *testing.T) {
	firstProfile, secondProfile := testMaybeNewProfiles(t)

	if firstProfile.GetID() == secondProfile.GetID() {
		t.Fatal("expected", false, "got", true)
	}
	if firstProfile.GetID() != firstProfile.GetID() {
		t.Fatal("expected", true, "got", false)
	}
	if secondProfile.GetID() != secondProfile.GetID() {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Profile_Profile_GetType(t *testing.T) {
	firstProfile, secondProfile := testMaybeNewProfiles(t)

	if firstProfile.GetType() != ObjectTypeCLGProfile {
		t.Fatal("expected", true, "got", false)
	}

	if firstProfile.GetType() != secondProfile.GetType() {
		t.Fatal("expected", true, "got", false)
	}
	if firstProfile.GetType() != firstProfile.GetType() {
		t.Fatal("expected", true, "got", false)
	}
	if secondProfile.GetType() != secondProfile.GetType() {
		t.Fatal("expected", true, "got", false)
	}
}
