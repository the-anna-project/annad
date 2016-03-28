// Package clg implementes fundamental actions used to create strategies that
// allow to discover new behavior for problem solving.
package clg

var (
	// Index represents the CLG index containing all CLGs that can be used.
	Index map[string]func(args ...interface{}) ([]interface{}, error)
)

func init() {
	Index = map[string]func(args ...interface{}) ([]interface{}, error){
		// String.
		"ContainsString": ContainsString,
		"LongerString":   LongerString,
		"RepeatString":   RepeatString,
		"ShorterString":  ShorterString,
		"SplitString":    SplitString,

		// String Slice.
		"ContainsStringSlice":  ContainsStringSlice,
		"JoinStringSlice":      JoinStringSlice,
		"SortStringSlice":      SortStringSlice,
		"SwapLeftStringSlice":  SwapLeftStringSlice,
		"SwapRightStringSlice": SwapRightStringSlice,

		// Type.
		"ArgType": ArgType,
	}
}
