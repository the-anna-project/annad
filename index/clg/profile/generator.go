package profile

import (
	"encoding/json"
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/factory/random"
	"github.com/xh3b4sd/anna/index/clg/collection"
	memoryinstrumentation "github.com/xh3b4sd/anna/instrumentation/memory"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	memorystorage "github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypeCLGProfileGenerator represents the object type of the CLG
	// profile generator object. This is used e.g. to register itself to the
	// logger.
	ObjectTypeCLGProfileGenerator spec.ObjectType = "clg-profile-generator"
)

var (
	// maxArgs represents the maximum number of arguments allowed used to find
	// out how many input arguments a CLG expects. Usually CLGs do not expect
	// more than 5 input arguments. In case a CLG expects 5 or more arguments, we
	// assume it expects infinite arguments.
	//
	// The reason why we limit so strictly is the vast amount of possibilities.
	// Imagine we have a list of about 100 different arguments that can be
	// combined with each other. In case of creating permutations of up to 5
	// arguments within one argument list, we can have more than 10 billion
	// possible combinations.
	//
	//     1+100^1+100^2+100^3+100^4+100^5 = 10.101.010.101
	//
	// This would be the amount of iterations we would need to make to find all
	// possible inputs of one single CLG. The goal is to have houndrets of CLGs.
	// This is why the process of finding out how the interface of an CLG looks
	// like needs to be optimized. The limitation of the number of input
	// arguments is one thing.
	//
	// Another optimization is to stop the discovery process here in case an
	// error is returned signaling too many arguments were used.
	//
	// Further, every now and then, Anna should try what else works out for a
	// CLG. How such a discovery looks like in detail is to be defined.
	//
	// Note that this number is experimental and might change if necessary.
	maxArgs = 5
)

// GeneratorConfig represents the configuration used to create a new CLG
// profile generator object.
type GeneratorConfig struct {
	// Dependencies.
	ArgumentListFactory func() (spec.PermutationList, error)
	Collection          spec.CLGCollection
	Instrumentation     spec.Instrumentation
	Log                 spec.Log
	PermutationFactory  spec.PermutationFactory
	RandomFactory       spec.RandomFactory
	Storage             spec.Storage

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

	newInstrumentation, err := memoryinstrumentation.NewInstrumentation(memoryinstrumentation.DefaultInstrumentationConfig())
	if err != nil {
		panic(err)
	}

	newPermutationFactory, err := permutation.NewFactory(permutation.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}

	newRandomFactory, err := random.NewFactory(random.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}

	newStorage, err := memorystorage.NewStorage(memorystorage.DefaultStorageConfig())
	if err != nil {
		panic(err)
	}

	newConfig := GeneratorConfig{
		// Dependencies.
		ArgumentListFactory: createArgumentListFactory(),
		Collection:          newCollection,
		Instrumentation:     newInstrumentation,
		Log:                 log.NewLog(log.DefaultConfig()),
		PermutationFactory:  newPermutationFactory,
		RandomFactory:       newRandomFactory,
		Storage:             newStorage,

		// Settings.
		LoaderReadFile:  collection.LoaderReadFile,
		LoaderFileNames: collection.LoaderFileNames,
	}

	return newConfig
}

// NewGenerator creates a new configured CLG profile generator object.
func NewGenerator(config GeneratorConfig) (spec.CLGProfileGenerator, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	// Create new object.
	newGenerator := &generator{
		GeneratorConfig: config,

		ID:    newID,
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

func (g *generator) CreateProfile(clgName string, canceler <-chan struct{}) (spec.CLGProfile, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call CreateProfile")

	var newProfile spec.CLGProfile
	action := func() error {
		// Fetch the CLG profile in advance.
		currentProfile, err := g.GetProfileByName(clgName)
		if IsCLGProfileNotFound(err) {
			// In case the CLG profile cannot be found, we are going ahead to create
			// one.
		} else if err != nil {
			return maskAny(err)
		}

		// Create mapper and scanner results for the current profile.
		newBody, err := g.CreateBody(clgName)
		if err != nil {
			return maskAny(err)
		}
		newHash, err := g.CreateHash(newBody)
		if err != nil {
			return maskAny(err)
		}
		newInputsOutputs, err := g.CreateInputsOutputs(clgName, canceler)
		if err != nil {
			return maskAny(err)
		}
		newName, err := g.CreateName(clgName)
		if err != nil {
			return maskAny(err)
		}

		// Create the new CLG profile.
		newConfig := DefaultConfig()
		newConfig.Body = newBody
		newConfig.Hash = newHash
		newConfig.InputsOutputs = newInputsOutputs
		newConfig.Name = newName
		newProfile, err = New(newConfig)
		if err != nil {
			return maskAny(err)
		}

		if currentProfile != nil {
			if !currentProfile.Equals(newProfile) {
				// The new profile differs from the current one. Thus we mark it as
				// having changed.
				newProfile.SetHashChanged(true)
			}

			// There is already a profile known. No matter if it changed or not, to not
			// change the ID we set it in all cases.
			newProfile.SetID(currentProfile.GetID())
		} else {
			// There is no profile known yet. Thus we mark it as having changed.
			newProfile.SetHashChanged(true)
		}

		return nil
	}

	err := g.Instrumentation.ExecFunc("CreateProfile", action)
	if err != nil {
		return nil, maskAny(err)
	}

	return newProfile, nil
}

func (g *generator) GetProfileByName(clgName string) (spec.CLGProfile, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call GetProfileByName")

	value, err := g.Storage.Get(g.key("profile:%s", string(clgName)))
	if err != nil {
		return nil, maskAny(err)
	}

	if value == "" {
		return nil, maskAny(clgProfileNotFoundError)
	}

	newProfile := NewEmptyProfile()
	err = json.Unmarshal([]byte(value), newProfile)
	if err != nil {
		return nil, maskAny(err)
	}

	return newProfile, nil
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

func (g *generator) StoreProfile(clgProfile spec.CLGProfile) error {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call StoreProfile")

	raw, err := json.Marshal(clgProfile)
	if err != nil {
		return maskAny(err)
	}

	err = g.Storage.Set(g.key("profile:%s", clgProfile.GetName()), string(raw))
	if err != nil {
		return maskAny(err)
	}

	return nil
}
