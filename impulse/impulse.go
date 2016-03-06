// Package impulse implementes spec.Impulse. An impulse can walk through any
// spec.Core, spec.Network and spec.Neuron. Concrete implementations and their
// dynamic state decide about the way an impulse is going, resulting in
// behaviour.
package impulse

import (
	"sync"

	"github.com/xh3b4sd/anna/ctx"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	ObjectTypeImpulse spec.ObjectType = "impulse"
)

type Config struct {
	Log spec.Log

	Input  string
	Output string
}

func DefaultConfig() Config {
	newConfig := Config{
		Input:  "",
		Log:    log.NewLog(log.DefaultConfig()),
		Output: "",
	}

	return newConfig
}

func NewImpulse(config Config) (spec.Impulse, error) {
	newImpulse := &impulse{
		Config: config,
		Ctxs:   map[spec.ObjectID]spec.Ctx{},
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeImpulse,
	}

	newImpulse.Log.Register(newImpulse.GetType())

	return newImpulse, nil
}

type impulse struct {
	Config

	Ctxs map[spec.ObjectID]spec.Ctx

	ID spec.ObjectID `json:"id"`

	Mutex sync.Mutex `json:"-"`

	Type spec.ObjectType `json:"type"`
}

func (i *impulse) GetInput() (string, error) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	return i.Input, nil
}

func (i *impulse) GetOutput() (string, error) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	return i.Output, nil
}

func (i *impulse) GetCtx(object spec.Object) spec.Ctx {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	if ctx, ok := i.Ctxs[object.GetID()]; ok {
		return ctx
	}

	newCtxConfig := ctx.DefaultConfig()
	newCtxConfig.Object = object
	newCtx := ctx.NewCtx(newCtxConfig)
	i.Ctxs[object.GetID()] = newCtx

	return newCtx
}

func (i *impulse) SetID(ID spec.ObjectID) error {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	i.ID = ID
	return nil
}

func (i *impulse) SetInput(input string) error {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	i.Input = input
	return nil
}

func (i *impulse) SetOutput(output string) error {
	i.Mutex.Lock()
	i.Output = output
	return nil
}
