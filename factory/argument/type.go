package argument

import (
	"github.com/xh3b4sd/anna/spec"
)

var (
	// Special type.

	// TypeNone is a special type of the argument factory, describing no
	// special concrete type.
	TypeNone spec.PermutationType = "none"

	// TypeArgumentList is a special type of the argument factory, describing no
	// special concrete type.
	TypeArgumentList spec.PermutationType = "argumentlist"

	// Normal type.

	// TypeBool describes the concrete type bool.
	TypeBool spec.PermutationType = "bool"

	// TypeDistribution describes the concrete type distribution.
	TypeDistribution spec.PermutationType = "distribution"

	// TypeFeatureSet describes the concrete type featureset.
	TypeFeatureSet spec.PermutationType = "featureset"

	// TypeFloat64 describes the concrete type float64.
	TypeFloat64 spec.PermutationType = "float64"

	// TypeInt describes the concrete type int.
	TypeInt spec.PermutationType = "int"

	// TypeString describes the concrete type string.
	TypeString spec.PermutationType = "string"

	// TypeFloat64Slice describes the concrete type []float64.
	TypeFloat64Slice spec.PermutationType = "[]float64"

	// TypeIntSlice describes the concrete type []int.
	TypeIntSlice spec.PermutationType = "[]int"

	// TypeStringSlice describes the concrete type []string.
	TypeStringSlice spec.PermutationType = "[]string"

	// TypeFloat64SliceSlice describes the concrete type [][]float64.
	TypeFloat64SliceSlice spec.PermutationType = "[][]float64"

	// TypeIntSliceSlice describes the concrete type [][]int.
	TypeIntSliceSlice spec.PermutationType = "[][]int"

	// TypeStringSliceSlice describes the concrete type [][]string.
	TypeStringSliceSlice spec.PermutationType = "[][]string"
)
