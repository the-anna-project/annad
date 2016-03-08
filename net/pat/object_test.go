package patnet

import (
	"testing"
)

func Test_PatNet_GetType(t *testing.T) {
	newNet, err := NewPatNet(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if newNet.GetType() != ObjectTypePatNet {
		t.Fatal("expected", ObjectTypePatNet, "got", newNet.GetType())
	}
}

func Test_PatNet_GetID(t *testing.T) {
	firstNet, err := NewPatNet(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	secondNet, err := NewPatNet(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	if firstNet.GetID() == secondNet.GetID() {
		t.Fatalf("IDs of jobs are equal")
	}
}
