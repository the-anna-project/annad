package context

import (
	nativecontext "context"
	"encoding/json"
)

type Context interface {
	nativecontext.Context
	json.Marshaler
	json.Unmarshaler
}
