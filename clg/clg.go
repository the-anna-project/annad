// Package clg implementes fundamental actions used to create strategies that
// allow to discover new behavior for problem solving.
package clg

var (
	Index map[string]func(args ...interface{}) ([]interface{}, error)
)

func init() {
	Index = map[string]func(args ...interface{}) ([]interface{}, error){
		// String.
		"ContainsString": ContainsString,
		"LongerString":   LongerString,
		"RepeatString":   RepeatString,
		"ShorterString":  ShorterString,

		// String Slice.
		"ContainsStringSlice": ContainsStringSlice,
		"SwapStringSlice":     SwapStringSlice,
	}
}
