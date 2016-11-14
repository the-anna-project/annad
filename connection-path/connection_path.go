// Package connectionpath implements spec.ConnectionPath for distance
// calculation within the connection space.
package connectionpath

import (
	"encoding/json"
	"math"

	"github.com/xh3b4sd/anna/service/random"
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

	return newConnectionPath, nil
}

// NewFromString provides a convenient way to create a new connection path from
// the raw string representation of its coordinates. Therefore a simple
// json.Unmarshal is used to transform the provided string into [][]float64.
// The vector list is then used to create a new connection path by using New.
func NewFromString(s string) (spec.ConnectionPath, error) {
	var cs [][]float64
	err := json.Unmarshal([]byte(s), &cs)
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

type connectionPath struct {
	Config
}

func (cp *connectionPath) DistanceTo(a spec.ConnectionPath) float64 {
	//
	smallerCoordinates, greaterCoordinates := cp.GetCoordinates(), a.GetCoordinates()
	smallerLength, greaterLength := len(smallerCoordinates), len(greaterCoordinates)

	if smallerLength > greaterLength {
		smallerCoordinates, greaterCoordinates = a.GetCoordinates(), cp.GetCoordinates()
		smallerLength, greaterLength = len(smallerCoordinates), len(greaterCoordinates)
	}

	numPeers := math.Floor(float64(greaterLength) / float64(smallerLength))

	var newCoordinates [][]float64
	for _, vector := range smallerCoordinates {
		for j := 0; j < int(numPeers); j++ {
			newCoordinates = append(newCoordinates, vector)
		}
	}

	if len(newCoordinates) < greaterLength {
		newCoordinates = append(newCoordinates, smallerCoordinates[smallerLength-1])
	}

	var distance float64
	for i, newVector := range newCoordinates {
		for j, newCoordinate := range newVector {
			greaterCoordinate := greaterCoordinates[i][j]

			var d float64
			if newCoordinate > greaterCoordinate {
				d = newCoordinate - greaterCoordinate
			} else {
				d = greaterCoordinate - newCoordinate
			}

			distance += d
		}
	}

	return distance
}

func (cp *connectionPath) GetCoordinates() [][]float64 {
	return cp.Coordinates
}

func (cp *connectionPath) IsCloser(a, b spec.ConnectionPath) (spec.ConnectionPath, error) {
	// At first we calculate the distance a and b have to cp.
	da := cp.DistanceTo(a)
	db := cp.DistanceTo(b)

	if da < db {
		// The sum of the distance of a to cp is the smaller one. Thus we qualified
		// it to be closer to cp and return a.
		return a, nil
	}
	if da > db {
		// The sum of the distance of a to cp is the bigger one. Thus we qualified
		// it to be farther to cp and return b.
		return b, nil
	}
	if da == db {
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
