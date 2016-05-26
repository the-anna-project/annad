package profile

import (
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewGenerator(t *testing.T) spec.CLGProfileGenerator {
	newGeneratorConfig, err := DefaultGeneratorConfig()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newGenerator, err := NewGenerator(newGeneratorConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newGenerator
}

func testMaybeNewProfileWithGenerator(t *testing.T, generator spec.CLGProfileGenerator, clgName string) spec.CLGProfile {
	newProfile, err := generator.CreateProfile(clgName)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newProfile
}

func Test_Profile_Generator_CreateProfile(t *testing.T) {
	newGenerator := testMaybeNewGenerator(t)

	clgName := "DiscardInterface"

	firstProfile := testMaybeNewProfileWithGenerator(t, newGenerator, clgName)
	secondProfile := testMaybeNewProfileWithGenerator(t, newGenerator, clgName)

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
