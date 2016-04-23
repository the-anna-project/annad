package distribution

import (
	"testing"
)

func Test_Common_equalDimensionLength(t *testing.T) {
	if equalDimensionLength([][]float64{}) {
		t.Fatal("expected", false, "got", true)
	}
}

func Test_maxFloat64(t *testing.T) {
	f := maxFloat64(nil)
	if f != 0 {
		t.Fatal("expected", 0, "got", f)
	}
	f = maxFloat64([]float64{})
	if f != 0 {
		t.Fatal("expected", 0, "got", f)
	}
}

func Test_minFloat64(t *testing.T) {
	f := minFloat64(nil)
	if f != 0 {
		t.Fatal("expected", 0, "got", f)
	}
	f = minFloat64([]float64{})
	if f != 0 {
		t.Fatal("expected", 0, "got", f)
	}
	f = minFloat64([]float64{1.1, 2.2, 3.3})
	if f != 1.1 {
		t.Fatal("expected", 1.1, "got", f)
	}
	f = minFloat64([]float64{2.2, 1.1, 3.3})
	if f != 1.1 {
		t.Fatal("expected", 1.1, "got", f)
	}
}
