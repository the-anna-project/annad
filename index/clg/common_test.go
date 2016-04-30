package clg

import (
	"testing"
)

func testMaybeFatalCase(t *testing.T, index int, err error) {
	if err != nil {
		t.Fatal("case", index+1, "expected", nil, "got", err)
	}
}
