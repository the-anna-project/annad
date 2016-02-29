// Package common provides basic functionality that is of general interest for
// the factory's internal subpackages.
package common

import (
	"github.com/xh3b4sd/anna/spec"
)

const (
	ObjectTypeCore            spec.ObjectType = "core"
	ObjectTypeImpulse         spec.ObjectType = "impulse"
	ObjectTypeRedisStorage    spec.ObjectType = "redis-storage"
	ObjectTypeStrategyNetwork spec.ObjectType = "strategy-network"
)
