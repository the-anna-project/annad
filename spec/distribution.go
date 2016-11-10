package spec

// Distribution is a calculation object that describes the distribution of
// vectors. This vectors can be of arbitrary dimensions. That way
// characteristics of features and their location in space can be represented.
type Distribution interface {
	// Calculate calculates the current distribution based on the given channel
	// and vector configuration. Here is calculated which vectors are located
	// within which channel. No matter how many dimensions a vector has and how
	// many of these are located within a channel, the weight of each vector is
	// 1. Is a vector located within multiple channels, the vector's weight is
	// divided across them.
	Calculate() []float64

	// Difference calculates the difference by simple subtraction of
	// distribution values. Is the given distribution above the receiving
	// distribution, the channel value will be positive, negative otherwise.
	Difference(dist Distribution) ([]float64, error)

	// GetDimensions returns the dimensions of the distribution's vectors. Note
	// that vectors within a distribution must have the same amount of dimensions.
	GetDimensions() int

	// GetName returns the distribution's name. Note this might be a sequence of
	// a feature.
	GetName() string

	// GetStaticChannels returns the distribution's static channels. These define
	// the block wise separation of calculated vector weights.
	GetStaticChannels() []float64

	// GetVectors returns the distribution's vectors. Their weights are
	// calculated based on their channel locations. Note that each vector has a
	// weight of 1.
	GetVectors() [][]float64
}
