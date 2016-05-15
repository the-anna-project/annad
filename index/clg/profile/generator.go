package profile

import (
	"encoding/json"
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/index/clg/collection"
	"github.com/xh3b4sd/anna/instrumentation/prometheus"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
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
	Collection      spec.CLGCollection
	Instrumentation spec.Instrumentation
	Log             spec.Log
	Storage         spec.Storage

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

	newPrometheusConfig := prometheus.DefaultConfig()
	newPrometheusConfig.Prefixes = append(newPrometheusConfig.Prefixes, "CLGProfileGenerator")
	newInstrumentation, err := prometheus.New(newPrometheusConfig)
	if err != nil {
		panic(err)
	}

	newConfig := GeneratorConfig{
		// Dependencies.
		Collection:      newCollection,
		Instrumentation: newInstrumentation,
		Log:             log.NewLog(log.DefaultConfig()),
		Storage:         memorystorage.NewMemoryStorage(memorystorage.DefaultConfig()),

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

func (g *generator) CreateProfile(clgName string) (spec.CLGProfile, error) {
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
		newInputs, err := g.CreateInputs(clgName)
		if err != nil {
			return maskAny(err)
		}
		newName, err := g.CreateName(clgName)
		if err != nil {
			return maskAny(err)
		}
		newOutputs, err := g.CreateOutputs(clgName)
		if err != nil {
			return maskAny(err)
		}

		// Create the new CLG profile.
		newConfig := DefaultConfig()
		newConfig.Body = newBody
		newConfig.Hash = newHash
		newConfig.Inputs = newInputs
		newConfig.Name = newName
		newConfig.Outputs = newOutputs
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
