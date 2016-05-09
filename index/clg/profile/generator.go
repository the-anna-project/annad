package profile

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/index/clg/collection"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeCLGProfileGenerator represents the object type of the CLG
	// profile generator object. This is used e.g. to register itself to the
	// logger.
	ObjectTypeCLGProfileGenerator spec.ObjectType = "clg-profile-generator"
)

// GeneratorConfig represents the configuration used to create a new CLG
// profile generator object.
type GeneratorConfig struct {
	// Dependencies.
	Collection spec.CLGCollection
	Log        spec.Log

	// Settings.
	LoaderFileNames func() []string
	LoaderReadFile  func(fileName string) ([]byte, error)
}

// DefaultGeneratorConfig provides a default configuration to create a new CLG
// profile generator object by best effort.
func DefaultGeneratorConfig() GeneratorConfig {
	newCollection, err := collection.New(collection.DefaultConfig())
	if err != nil {
		panic(err)
	}

	newConfig := GeneratorConfig{
		// Dependencies.
		Collection: newCollection,
		Log:        log.NewLog(log.DefaultConfig()),

		// Settings.
		LoaderReadFile:  collection.LoaderReadFile,
		LoaderFileNames: collection.LoaderFileNames,
	}

	return newConfig
}

// NewGenerator creates a new configured CLG profile generator object.
func NewGenerator(config GeneratorConfig) (spec.CLGProfileGenerator, error) {
	// Create new object.
	newGenerator := &generator{
		GeneratorConfig: config,

		ID:    id.NewObjectID(id.Hex128),
		Mutex: sync.Mutex{},
		Type:  ObjectTypeCLGProfileGenerator,
	}

	// Validate new object.
	if newGenerator.Collection == nil {
		return nil, maskAnyf(invalidConfigError, "CLG collection must not be empty")
	}
	if newGenerator.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}

	// Register in logger.
	newGenerator.Log.Register(newGenerator.GetType())

	return newGenerator, nil
}

type generator struct {
	GeneratorConfig

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (g *generator) CreateProfile(clgName string) (spec.CLGProfile, bool, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call CreateProfile")

	// Fetch the CLG profile in advance.
	currentProfile, err := g.GetProfileByName(clgName)
	if IsCLGProfileNotFound(err) {
		// In case the CLG profile cannot be found, we are going ahead to create
		// one.
	} else if err != nil {
		return nil, false, maskAny(err)
	}

	// Create mapper and scanner results for the current profile.
	newBody, err := g.CreateBody(clgName)
	if err != nil {
		return nil, false, maskAny(err)
	}
	newHash, err := g.CreateHash(newBody)
	if err != nil {
		return nil, false, maskAny(err)
	}
	newInputs, err := g.CreateInputs(clgName)
	if err != nil {
		return nil, false, maskAny(err)
	}
	newName, err := g.CreateName(clgName)
	if err != nil {
		return nil, false, maskAny(err)
	}
	newOutputs, err := g.CreateOutputs(clgName)
	if err != nil {
		return nil, false, maskAny(err)
	}

	// Create the new CLG profile.
	newConfig := DefaultConfig()
	newConfig.Hash = newHash
	newConfig.Inputs = newInputs
	newConfig.Body = newBody
	newConfig.Name = newName
	newConfig.Outputs = newOutputs
	newProfile, err := New(newConfig)
	if err != nil {
		return nil, false, maskAny(err)
	}

	if currentProfile != nil && currentProfile.Equals(newProfile) {
		// The CLG profile has not changed. Thus nothing to do here.
		return currentProfile, false, nil
	}

	return newProfile, true, nil
}

// TODO
func (g *generator) GetProfileByName(clgName string) (spec.CLGProfile, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call GetProfileByName")

	return nil, nil
}

func (g *generator) GetProfileNames() ([]string, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call GetProfileNames")

	args, err := g.Collection.GetNamesMethod()
	if err != nil {
		return nil, maskAny(err)
	}
	profileNames, err := collection.ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}

	return profileNames, nil
}

// TODO
func (g *generator) StoreProfile(clgProfile spec.CLGProfile) error {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call StoreProfile")

	return nil
}
