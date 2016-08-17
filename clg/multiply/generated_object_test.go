package multiply

import (
	"testing"
)

func Test_CLG_GetID(t *testing.T) {
	firstCLG := MustNew()
	secondCLG := MustNew()

	if firstCLG.GetID() == secondCLG.GetID() {
		t.Fatal("expected", false, "got", true)
	}
}

func Test_CLG_GetType(t *testing.T) {
	newCLG := MustNew()
	objectType := newCLG.GetType()

	if objectType != ObjectType {
		t.Fatal("expected", ObjectType, "got", objectType)
	}
}
