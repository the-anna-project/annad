package patnet

import (
	"reflect"
	"sort"
	"strconv"
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeDistribution represents the object type of the distribution
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeDistribution spec.ObjectType = "distribution"
)

// DistributionConfig represents the configuration used to create a new
// distribution object.
type DistributionConfig struct {
	// HashMap provides a way to create a new distribution object out of a given
	// hash map containing bare distribution data. If this is nil or empty, a
	// completely new distribution is created. Otherwise it is tried to create a
	// new distribution using the information of the given hash map.
	HashMap map[string]string

	// Name represents the name of the distribution.
	Name string

	// StaticChannels represents the statically configured channels used to
	// calculate a weighted analysis.
	StaticChannels []float64

	// Vectors represents a list of vector positions within space.
	Vectors [][]float64
}

// DefaultDistributionConfig provides a default configuration to create a new
// distribution object by best effort.
func DefaultDistributionConfig() DistributionConfig {
	newConfig := DistributionConfig{
		Name:           "",
		StaticChannels: []float64{5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100},
		Vectors:        [][]float64{},
	}

	return newConfig
}

// NewDistribution creates a new configured distribution object. A distribution
// represents a weighted analysis of vectors. A two dimensional distribution
// can be seen as bar chart.
//
//     ^
//     |         x
//     |         x
//     |         x    x
//   y |         x    x
//     |         x    x
//     |    x    x    x
//     |    x    x    x    x
//     |    x    x    x    x
//     +------------------------>
//                 x
//
func NewDistribution(config DistributionConfig) (spec.Distribution, error) {
	var newDistribution *distribution

	if config.HashMap != nil {
		newDistribution = &distribution{}

		for key, value := range config.HashMap {
			if key == "name" {
				newDistribution.Name = value
			}
			if key == "id" {
				newDistribution.ID = spec.ObjectID(value)
			}
			if key == "static-channels" {
				newStaticChannels, err := staticChannelsFromString(value)
				if err != nil {
					return nil, maskAnyf(invalidConfigError, err.Error())
				}
				newDistribution.StaticChannels = newStaticChannels
			}
			if key == "vectors" {
				newVectors, err := vectorsFromString(value)
				if err != nil {
					return nil, maskAnyf(invalidConfigError, err.Error())
				}
				newDistribution.Vectors = newVectors
			}
		}
	} else {
		newDistribution = &distribution{
			DistributionConfig: config,
			ID:                 id.NewObjectID(id.Hex128),
			Mutex:              sync.Mutex{},
			Type:               ObjectTypeDistribution,
		}
	}

	if newDistribution.Name == "" {
		return nil, maskAnyf(invalidConfigError, "name must not be empty")
	}
	if len(newDistribution.Vectors) == 0 {
		return nil, maskAnyf(invalidConfigError, "vectors must not be empty")
	}
	if !equalDimensionLength(newDistribution.Vectors) {
		return nil, maskAnyf(invalidConfigError, "vectors must have equal dimensions")
	}
	if len(newDistribution.StaticChannels) == 0 {
		return nil, maskAnyf(invalidConfigError, "vectors must not be empty")
	}
	if !uniqueFloat64(newDistribution.StaticChannels) {
		return nil, maskAnyf(invalidConfigError, "static channels must be unique")
	}

	sort.Float64s(newDistribution.StaticChannels)

	return newDistribution, nil
}

// TODO add moving similarity with threshold
// TODO add surrounding lengths
// TODO add sequence combination distributions
//
//
//
//
//
//
//
//
//
// TODO find a way to find these patterns automatically
type distribution struct {
	DistributionConfig

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (d *distribution) Calculate() []float64 {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()

	eval := mapVectorsToChannels(d.Vectors, d.StaticChannels)

	return eval
}

func (d *distribution) Difference(dist spec.Distribution) ([]float64, error) {
	if !reflect.DeepEqual(d.GetStaticChannels(), dist.GetStaticChannels()) {
		return nil, maskAnyf(channelsDifferError, "channels must be equal")
	}

	perc1 := d.Calculate()
	perc2 := dist.Calculate()
	diff := channelDistance(perc1, perc2)

	return diff, nil
}

func (d *distribution) GetDimensions() int {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()

	return len(d.Vectors[0])
}

func (d *distribution) GetHashMap() map[string]string {
	newHashMap := map[string]string{
		"name":            d.GetName(),
		"id":              string(d.GetID()),
		"static-channels": "",
		"vectors":         "",
	}

	var staticChannels string
	for i, c := range d.GetStaticChannels() {
		if i > 0 {
			staticChannels += ","
		}
		staticChannels += strconv.FormatFloat(c, 'f', -1, 64)
	}
	newHashMap["static-channels"] = string(staticChannels)

	var vectors string
	for i, vector := range d.GetVectors() {
		if i > 0 {
			vectors += "|"
		}
		for j, d := range vector {
			if j > 0 {
				vectors += ","
			}
			vectors += strconv.FormatFloat(d, 'f', -1, 64)
		}
	}
	newHashMap["vectors"] = string(vectors)

	return newHashMap
}

func (d *distribution) GetName() string {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()

	return d.Name
}

func (d *distribution) GetStaticChannels() []float64 {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()

	return d.StaticChannels
}

func (d *distribution) GetVectors() [][]float64 {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()

	return d.Vectors
}
