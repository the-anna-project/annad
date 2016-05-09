package profile

import (
	"reflect"

	"github.com/xh3b4sd/anna/index/clg/collection"
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
	Mapper     spec.CLGMapper
	Scanner    spec.CLGScanner
}

// DefaultConfig provides a default configuration to create a new CLG profile
// generator object by best effort.
func DefaultConfig() GeneratorConfig {
	newCollection, err := collection.New(collection.DefaultConfig())
	if err != nil {
		panic(err)
	}
	newMapper, err := NewMapper(DefaultMapperConfig())
	if err != nil {
		panic(err)
	}
	newScanner, err := NewScanner(DefaultScannerConfig())
	if err != nil {
		panic(err)
	}

	newConfig := GeneratorConfig{
		// Dependencies.
		Collection: newCollection,
		Log:        log.NewLog(log.DefaultConfig()),
		Mapper:     newMapper,
		Scanner:    newScanner,
	}

	return newConfig
}

// NewGenerator creates a new configured CLG profile generator object.
func NewGenerator(config GeneratorConfig) (spec.CLGProfileGenerator, error) {
	// Create new object.
	newGenerator := &generator{
		GeneratorConfig: config,

		ID:           id.NewObjectID(id.Hex128),
		Mutex:        sync.Mutex{},
		ProfileNames: nil,
		Type:         ObjectTypeCLGProfileGenerator,
	}

	// Validate new object.
	if newGenerator.Collection == nil {
		return nil, maskAnyf(invalidConfigError, "CLG collection must not be empty")
	}
	if newGenerator.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newGenerator.Mapper == nil {
		return nil, maskAnyf(invalidConfigError, "CLG mapper must not be empty")
	}
	if newGenerator.Scanner == nil {
		return nil, maskAnyf(invalidConfigError, "CLG scanner must not be empty")
	}

	// Register in logger.
	newGenerator.Log.Register(newGenerator.GetType())

	// Create CLG lookup table.
	newLookupTable, err := createLookupTable()
	if err != nil {
		return nil, maskAnyf(invalidConfigError, err.Error())
	}
	newGenerator.LookupTable = lookupTable

	return newGenerator, nil
}

type generator struct {
	GeneratorConfig

	ID           spec.ObjectID
	Mutex        sync.Mutex
	ProfileNames []string
	Type         spec.ObjectType
}

func (g *generator) GetProfileNames() ([]string, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call CreateNameQueue")

	if g.ProfileNames != nil {
		return g.ProfileNames, nil
	}

	profileNames, err := profileNamesFromCollection(g.Collection)
	if err != nil {
		return nil, maskAny(err)
	}
	g.ProfileNames = profileNames

	return g.ProfileNames, nil
}

func (g *generator) CreateProfile(clgName string, canceler <-chan struct{}) (spec.CLGProfile, bool, error) {
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
	newMapResult, err = g.Mapper.CreateResult(clgName, canceler)
	if err != nil {
		return nil, false, maskAny(err)
	}
	newScanResult, err = g.Scanner.CreateResult(clgName, canceler)
	if err != nil {
		return nil, false, maskAny(err)
	}

	// Create the new CLG profile.
	newConfig := DefaultConfig()
	newConfig.Hash = newMapResult.Hash
	newConfig.InputTypes = newMapResult.InputTypes
	newConfig.MethodBody = newMapResult.MethodBody
	newConfig.MethodName = newMapResult.MethodName
	newConfig.OutputTypes = newMapResult.OutputTypes
	newConfig.RightSideNeighbours = newScanResult.RightSideNeighbours
	newProfile, err := New(newConfig)
	if err != nil {
		return nil, false, maskAny(err)
	}

	if currentProfile != nil && currentProfile.Equals(newProfile) {
		// The CLG profile has not changed. Thus nothing to do here.
		return currentProfile, false, nil
	}

	return newProfile, nil
}

// TODO
func (g *generator) GetProfileByName(clgName string) (spec.CLGProfile, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call StoreCLGProfile")

	return nil, nil
}

// TODO
func (g *generator) StoreProfile(clgProfile spec.CLGProfile) error {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call StoreProfile")

	return nil
}
