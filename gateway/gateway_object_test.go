package gateway

import (
	"testing"
)

func Test_Gateway_GetType(t *testing.T) {
	newGateway := NewGateway(DefaultConfig())

	if newGateway.GetType() != ObjectTypeGateway {
		t.Fatal("expected", ObjectTypeGateway, "got", newGateway.GetType())
	}
}

func Test_Gateway_GetID(t *testing.T) {
	firstGateway := NewGateway(DefaultConfig())
	secondGateway := NewGateway(DefaultConfig())

	if firstGateway.GetID() == secondGateway.GetID() {
		t.Fatalf("IDs of jobs are equal")
	}
}
