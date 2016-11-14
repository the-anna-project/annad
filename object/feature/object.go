package feature

import (
	"sync"

	objectspec "github.com/xh3b4sd/anna/object/spec"
)

// New creates a new feature object. A feature represents a differentiable part
// of an information sequence.
func New() objectspec.Feature {
	return &object{}
}

type object struct {
	// Settings.

	// positions represents the index locations of a detected feature.
	positions [][]float64
	// sequence represents the input sequence being detected as feature. That
	// means, the sequence of a feature object is the feature itself.
	sequence string
	mutex    sync.Mutex
}

func (o *object) AddPosition(position []float64) error {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	if len(o.positions) > 0 && len(o.positions[0]) != len(position) {
		return maskAnyf(invalidPositionError, "must have length of %d", len(o.positions))
	}

	o.positions = append(o.positions, position)

	return nil
}

func (o *object) Count() int {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return len(o.positions)
}

func (o *object) Positions() [][]float64 {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.positions
}

func (o *object) Sequence() string {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	return o.sequence
}

func (o *object) SetPositions(ps [][]float64) {
	o.positions = ps
}

func (o *object) SetSequence(s string) {
	o.sequence = s
}
