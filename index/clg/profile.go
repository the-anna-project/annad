package clg

import (
	"reflect"

	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeCLGProfile represents the object type of the CLG profile object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeCLGProfile spec.ObjectType = "clg-profile"
)

// ProfileConfig represents the configuration used to create a new CLG profile
// object.
type ProfileConfig struct {
	// Settings.

	// MethodName represents the CLG's implemented method name.
	MethodName string

	// MethodHash represents the hashed value of the CLG's implemented method
	// body.
	MethodHash string

	// InputTypes represents the CLG's implemented method input parameter types.
	InputTypes []reflect.Kind

	// InputExamples represents the CLG's implemented method input parameter
	// examples.
	InputExamples []interface{}

	// RightSideNeighbours represents the CLG's that can be used to combine it
	// with the current CLG.
	//
	//     Input -> CurrentCLG -> Output
	//                            Input -> RightSideNeighbourCLG -> Output
	//
	RightSideNeighbours []string
}

// DefaultCLGProfileConfig provides a default configuration to create a new CLG
// index object by best effort.
func DefaultCLGProfileConfig() ProfileConfig {
	newConfig := ProfileConfig{
		// Settings.
		MethodName:          "",
		MethodHash:          "",
		InputTypes:          nil,
		InputExamples:       nil,
		RightSideNeighbours: nil,
	}

	return newConfig
}

// NewCLGProfile creates a new configured CLG index object.
func NewCLGProfile(config ProfileConfig) (spec.CLGProfile, error) {
	newCLGProfile := &clgProfile{
		ProfileConfig: config,

		Type: ObjectTypeCLGProfile,
	}

	if newCLGProfile.MethodName == "" {
		return nil, maskAnyf(invalidConfigError, "method name of CLG profile must not be empty")
	}
	if newCLGProfile.MethodHash == "" {
		return nil, maskAnyf(invalidConfigError, "method hash of CLG profile must not be empty")
	}

	return newCLGProfile, nil
}

type clgProfile struct {
	ProfileConfig

	Type spec.ObjectType
}

func (p *clgProfile) Equals(other spec.CLGProfile) bool {
	if p.GetMethodName() != other.GetMethodName() {
		return false
	}
	if p.GetMethodHash() != other.GetMethodHash() {
		return false
	}
	if !reflect.DeepEqual(p.GetInputTypes(), other.GetInputTypes()) {
		return false
	}
	if !reflect.DeepEqual(p.GetInputExamples(), other.GetInputExamples()) {
		return false
	}
	if !reflect.DeepEqual(p.GetRightSideNeighbours(), other.GetRightSideNeighbours()) {
		return false
	}

	return true
}

func (p *clgProfile) GetMethodName() string {
	return p.MethodName
}

func (p *clgProfile) GetMethodHash() string {
	return p.MethodHash
}

func (p *clgProfile) GetInputTypes() []reflect.Kind {
	return p.InputTypes
}

func (p *clgProfile) GetInputExamples() []interface{} {
	return p.InputExamples
}

func (p *clgProfile) GetRightSideNeighbours() []string {
	return p.RightSideNeighbours
}
