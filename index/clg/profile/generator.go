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
	Collection  spec.CLGCollection
	Log         spec.Log
	LookupTable spec.CLGLookupTable
}

// DefaultConfig provides a default configuration to create a new CLG profile
// generator object by best effort.
func DefaultConfig() GeneratorConfig {
	newCollection, err := collection.New(collection.DefaultConfig())
	if err != nil {
		panic(err)
	}
	newLookupTable, err := NewLookupTable(DefaultLookupTableConfig())
	if err != nil {
		panic(err)
	}

	newConfig := GeneratorConfig{
		// Dependencies.
		Collection:  newCollection,
		Log:         log.NewLog(log.DefaultConfig()),
		LookupTable: newLookupTable,
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
	if newGenerator.LookupTable == nil {
		return nil, maskAnyf(invalidConfigError, "CLG lookup table must not be empty")
	}

	// Register in logger.
	newGenerator.Log.Register(newGenerator.GetType())

	// Fetch all CLG names.
	newNames, err := namesFromCollection(newGenerator.Collection)
	if err != nil {
		return nil, maskAnyf(invalidConfigError, err.Error())
	}
	newGenerator.CLGNames = newNames

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

	CLGNames    []string
	ID          spec.ObjectID
	LookupTable map[string]string
	Mutex       sync.Mutex
	Type        spec.ObjectType
}

func (g *generator) CreateNameQueue() chan string {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call CreateNameQueue")

	newNameQueue := make(chan string, len(g.CLGNames))

	for _, name := range g.CLGNames {
		newNameQueue <- name
	}

	return newNameQueue
}

func (g *generator) CreateProfileQueue() chan spec.CLGProfile {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call CreateProfileQueue")

	newProfileQueue := make(chan spec.CLGProfile, len(g.CLGNames))
	return newProfileQueue
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

	// Create the new CLG profile.
	var err error
	newConfig := DefaultConfig()
	newConfig.Hash, err = g.getCLGHash(clgName, clgBody)
	if err != nil {
		return nil, false, maskAny(err)
	}
	newConfig.InputTypes, err = g.getCLGInputTypes(methodValue)
	if err != nil {
		return nil, false, maskAny(err)
	}
	newConfig.InputExamples, err = g.getCLGInputExamples(methodValue)
	if err != nil {
		return nil, false, maskAny(err)
	}
	newConfig.MethodName = clgName
	clgBody, ok := g.LookupTable[clgName]
	if !ok {
		return nil, false, maskAnyf(clgBodyNotFoundError, clgName)
	}
	newConfig.MethodBody = clgBody
	newConfig.OutputTypes, err = g.getCLGOutputTypes(methodValue)
	if err != nil {
		return nil, false, maskAny(err)
	}
	newConfig.OutputExamples, err = g.getCLGOutputExamples(methodValue)
	if err != nil {
		return nil, false, maskAny(err)
	}
	newConfig.RightSideNeighbours, err = g.getCLGRightSideNeighbours(collection, clgName, methodValue, canceler)
	if err != nil {
		return nil, false, maskAny(err)
	}
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
func (g *generator) getCLGInputExamples(methodValue reflect.Value) ([]interface{}, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call getCLGInputExamples")

	return nil, nil
}

// TODO
func (g *generator) getCLGInputTypes(methodValue reflect.Value) ([]reflect.Kind, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call getCLGInputTypes")

	// TODO move to generator
	methodValue := reflect.ValueOf(collection).MethodByName(clgName)
	if !g.isMethodValue(methodValue) {
		return nil, maskAnyf(invalidCLGError, clgName)
	}

	return nil, nil
}

// TODO
func (g *generator) getCLGHash(clgName, clgBody string) (string, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call getCLGHash")

	return "", nil
}

// TODO
func (g *generator) getCLGOutputExamples(methodValue reflect.Value) ([]interface{}, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call getCLGOutputExamples")

	return nil, nil
}

// TODO
func (g *generator) getCLGOutputTypes(methodValue reflect.Value) ([]reflect.Kind, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call getCLGOutputTypes")

	return nil, nil
}

// TODO
func (g *generator) getCLGRightSideNeighbours(collection spec.CLGCollection, clgName string, methodValue reflect.Value, canceler <-chan struct{}) ([]string, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call getCLGRightSideNeighbours")

	// TODO
	// Fill a queue with all CLG names.
	clgNameQueue := g.getCLGNameQueue(g.CLGNames)

	// Initialize the profile creation.
	for {
		select {
		case <-canceler:
			return nil, maskAny(workerCanceledError)
		case clgName := <-clgNameQueue:
		}
	}

	//     find right side neighbours for given clg name
	//         if no profile for checked neighbour
	//             push neighbour name back to channel

	return nil, nil
}

// TODO
func (g *generator) isRightSideCLGNeighbour(collection spec.CLGCollection, left, right spec.CLGProfile) (bool, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call isRightSideNeighbour")

	// run clg chain
	// if error
	//     return false

	return false, nil
}

func (g *generator) isMethodValue(v reflect.Value) bool {
	if !v.IsValid() {
		return false
	}

	if v.Kind() != reflect.Func {
		return false
	}

	return true
}
