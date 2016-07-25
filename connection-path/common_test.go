package connectionpath

import (
	"testing"
)

func Test_ConnectionPath_equalDimensionLength(t *testing.T) {
	if equalDimensionLength([][]float64{}) {
		t.Fatal("expected", false, "got", true)
	}
}
