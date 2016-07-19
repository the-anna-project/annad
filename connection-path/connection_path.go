// Package connectionpath implementes spec.ConnectionPath for distance
// calculation within the connection space.
package connectionpath

import (
	"encoding/json"

	"github.com/xh3b4sd/anna/factory/random"
	"github.com/xh3b4sd/anna/spec"
)

// Config represents the configuration used to create a new connection path
// object.
type Config struct {
	// Settings.

	// Coordinates represents the location of the connection path within the
	// connection space. This location is defined by a list of multi dimensional
	// vectors. The amount of dimensions does not matter as soon as at least one
	// dimension is provided and all vectors have the same amount of dimensions.
	// In case this is not true, the connection path is considered invalid, and
	// calls to Validate will return errors.
	Coordinates [][]float64
}

// DefaultConfig provides a default configuration to create a new connection
// path object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Settings.
		Coordinates: [][]float64{},
	}

	return newConfig
}

// New creates a new configured connection path object.
func New(config Config) (spec.ConnectionPath, error) {
	newConnectionPath := &connectionPath{
		Config: config,
	}

	err := newConnectionPath.Validate()
	if err != nil {
		return nil, maskAnyf(invalidConfigError, "Validate failed: %s", err.Error())
	}

	return newConnectionPath, nil
}

// NewFromString provides a convenient way to create a new connection path from
// the raw string representation of its coordinates. Therefore a simple
// json.Unmarshal is used to transform the provided string into [][]float64.
// The vector list is then used to create a new connection path by using New.
func NewFromString(s string) (spec.ConnectionPath, error) {
	var cs [][]float64
	err := json.Unmarshal([]byte(string), &cs)
	if err != nil {
		return nil, maskAnyf(invalidConfigError, "cannot parse string to [][]float64: %s", err.Error())
	}

	newConfig := DefaultConfig()
	newConfig.Coordinates = cs
	newConnectionPath, err := New(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newConnectionPath, nil
}

// TODO write tests
type connectionPath struct {
	Config
}

// TODO
func (cp *connectionPath) DistanceTo(a spec.ConnectionPath) ([]float64, error) {
	return nil, nil
}

func (cp *connectionPath) GetCoordinates() [][]float64 {
	return cp.Coordinates
}

func (cp *connectionPath) IsCloser(a, b spec.ConnectionPath) (spec.ConnectionPath, error) {
	// At first we calculate the distance a and b have to cp.
	dta, err := cp.DistanceTo(a)
	if err != nil {
		return nil, maskAny(err)
	}
	dtb, err := cp.DistanceTo(a)
	if err != nil {
		return nil, maskAny(err)
	}

	// For simplicity we sum the distances to have a comparable value.
	var sa float64
	for _, f := range dta {
		sa += f
	}
	var sb float64
	for _, f := range dtb {
		sb += f
	}

	if sa < sb {
		// The sum of the distance of a to cp is the smaller one. Thus we qualified
		// it to be closer to cp and return a.
		return a, nil
	}
	if sa > sb {
		// The sum of the distance of a to cp is the bigger one. Thus we qualified
		// it to be farther to cp and return b.
		return b, nil
	}
	if sa == sb {
		if random.Bit() == 0 {
			// The sum of the distance of a to cp is equal to the sum of the distance
			// of b to cp. We rolled the dice and got a 0. Thus we return a.
			return a, nil
		}
	}
	// The sum of the distance of a to cp is equal to the sum of the distance of
	// b to cp. We rolled the dice and got a 1. Thus we return b.
	return b, nil
}

func (cp *connectionPath) String() (string, error) {
	b, err := json.Marshal(cp.GetCoordinates())
	if err != nil {
		return "", maskAny(err)
	}

	return string(b), nil
}

func (cp *connectionPath) Validate() error {
	if len(cp.GetCoordinates()) == 0 {
		return maskAnyf(invalidConnectionPathError, "coordinates must not be empty")
	}

	if !equalDimensionLength(cp.GetCoordinates()) {
		return maskAnyf(invalidConnectionPathError, "coordinates must have equal vector lenghts")
	}

	return nil
}
