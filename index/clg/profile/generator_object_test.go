package profile

import (
	"testing"
)

func Test_Profile_Generator_GetID(t *testing.T) {
	firstGenerator := testMaybeNewGenerator(t)
	secondGenerator := testMaybeNewGenerator(t)

	if firstGenerator.GetID() == secondGenerator.GetID() {
		t.Fatal("expected", false, "got", true)
	}
	if firstGenerator.GetID() != firstGenerator.GetID() {
		t.Fatal("expected", true, "got", false)
	}
	if secondGenerator.GetID() != secondGenerator.GetID() {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Profile_Generator_GetType(t *testing.T) {
	firstGenerator := testMaybeNewGenerator(t)
	secondGenerator := testMaybeNewGenerator(t)

	if firstGenerator.GetType() != ObjectTypeCLGProfileGenerator {
		t.Fatal("expected", true, "got", false)
	}

	if firstGenerator.GetType() != secondGenerator.GetType() {
		t.Fatal("expected", true, "got", false)
	}
	if firstGenerator.GetType() != firstGenerator.GetType() {
		t.Fatal("expected", true, "got", false)
	}
	if secondGenerator.GetType() != secondGenerator.GetType() {
		t.Fatal("expected", true, "got", false)
	}
}
