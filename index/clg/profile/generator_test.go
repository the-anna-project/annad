package profile

import (
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewGenerator(t *testing.T) spec.CLGProfileGenerator {
	newGenerator, err := NewGenerator(DefaultGeneratorConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newGenerator
}

func testMaybeNewProfileWithGenerator(t *testing.T, generator spec.CLGProfileGenerator, clgName string) spec.CLGProfile {
	var canceler chan struct{} // We don't need a canceler here.
	newProfile, err := generator.CreateProfile(clgName, canceler)
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

func Test_Profile_Generator_GetProfileNames(t *testing.T) {
	newGenerator := testMaybeNewGenerator(t)

	expectedCLGNameSubset := []string{
		"AddPositionFeature",
		"AppendFloat64Slice",
		"CallByNameMethod",
		"ContainsString",
		"DiscardInterface",
		"SumInt",
	}

	newProfileNames, err := newGenerator.GetProfileNames()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	for i, e := range expectedCLGNameSubset {
		var found bool

		for _, pn := range newProfileNames {
			if pn == e {
				found = true
				break
			}
		}

		if !found {
			t.Fatal("case", i+1, "expected", true, "got", false)
		}
	}
}
