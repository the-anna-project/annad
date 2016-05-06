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

	// Hash represents the hashed value of the CLG's implemented method.
	Hash string

	// InputTypes represents the CLG's implemented method input parameter types.
	InputTypes []reflect.Kind

	// InputExamples represents the CLG's implemented method input parameter
	// examples.
	InputExamples []interface{}

	// MethodName represents the CLG's implemented method name.
	MethodName string

	// MethodBody represents the CLG's implemented method body.
	MethodBody string

	// OutputTypes represents the CLG's implemented method output parameter types.
	OutputTypes []reflect.Kind

	// OutputExamples represents the CLG's implemented method output parameter
	// examples.
	OutputExamples []interface{}

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
		Hash:                "",
		InputTypes:          nil,
		InputExamples:       nil,
		MethodName:          "",
		MethodBody:          "",
		OutputTypes:         nil,
		OutputExamples:      nil,
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

	if newCLGProfile.Hash == "" {
		return nil, maskAnyf(invalidConfigError, "hash of CLG profile must not be empty")
	}
	if newCLGProfile.MethodName == "" {
		return nil, maskAnyf(invalidConfigError, "method name of CLG profile must not be empty")
	}
	if newCLGProfile.MethodBody == "" {
		return nil, maskAnyf(invalidConfigError, "method body of CLG profile must not be empty")
	}
	if len(newCLGProfile.OutputTypes) == 0 {
		return nil, maskAnyf(invalidConfigError, "output types of CLG profile must not be empty")
	}
	if len(newCLGProfile.OutputExamples) == 0 {
		return nil, maskAnyf(invalidConfigError, "output examples of CLG profile must not be empty")
	}

	return newCLGProfile, nil
}

type clgProfile struct {
	ProfileConfig

	Type spec.ObjectType
}

func (p *clgProfile) Equals(other spec.CLGProfile) bool {
	if p.GetHash() != other.GetHash() {
		return false
	}
	if !reflect.DeepEqual(p.GetInputTypes(), other.GetInputTypes()) {
		return false
	}
	if !reflect.DeepEqual(p.GetInputExamples(), other.GetInputExamples()) {
		return false
	}
	if p.GetMethodName() != other.GetMethodName() {
		return false
	}
	if p.GetMethodBody() != other.GetMethodBody() {
		return false
	}
	if !reflect.DeepEqual(p.GetOutputTypes(), other.GetOutputTypes()) {
		return false
	}
	if !reflect.DeepEqual(p.GetOutputExamples(), other.GetOutputExamples()) {
		return false
	}
	if !reflect.DeepEqual(p.GetRightSideNeighbours(), other.GetRightSideNeighbours()) {
		return false
	}

	return true
}

func (p *clgProfile) GetHash() string {
	return p.Hash
}

func (p *clgProfile) GetInputTypes() []reflect.Kind {
	return p.InputTypes
}

func (p *clgProfile) GetInputExamples() []interface{} {
	return p.InputExamples
}

func (p *clgProfile) GetMethodName() string {
	return p.MethodName
}

func (p *clgProfile) GetMethodBody() string {
	return p.MethodBody
}

func (p *clgProfile) GetOutputTypes() []reflect.Kind {
	return p.OutputTypes
}

func (p *clgProfile) GetOutputExamples() []interface{} {
	return p.OutputExamples
}

func (p *clgProfile) GetRightSideNeighbours() []string {
	return p.RightSideNeighbours
}
