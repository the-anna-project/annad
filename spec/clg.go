package spec

// Index represents the CLG index providing all CLGs that can be used.
type Index interface {
	// Call provides a way to dynamically execute any index's method by providing
	// the method's name as string argument.
	Call(methodName string, args ...interface{}) ([]interface{}, error)

	// String.

	// ContainsString provides functionality of strings.Contains.
	ContainsString(args ...interface{}) ([]interface{}, error)

	// ContainsString provides functionality to check if one string is longer
	// than the other.
	LongerString(args ...interface{}) ([]interface{}, error)

	// ContainsString provides functionality of strings.Repeat.
	RepeatString(args ...interface{}) ([]interface{}, error)

	// ContainsString provides functionality to check if one string is shorter
	// than the other.
	ShorterString(args ...interface{}) ([]interface{}, error)

	// ContainsString provides functionality of strings.Split.
	SplitString(args ...interface{}) ([]interface{}, error)

	// String slice.

	// ContainsStringSlice provides functionality to check if a string slice
	// contains a certain member.
	ContainsStringSlice(args ...interface{}) ([]interface{}, error)

	// JoinStringSlice provides functionality of strings.Join.
	JoinStringSlice(args ...interface{}) ([]interface{}, error)

	// SortStringSlice provides functionality of sort.Strings.
	SortStringSlice(args ...interface{}) ([]interface{}, error)

	// SwapLeftStringSlice provides functionality to move the first member of a
	// string slice to the left, that is, the end of the string slice.
	SwapLeftStringSlice(args ...interface{}) ([]interface{}, error)

	// SwapRightStringSlice provides functionality to move the last member of a
	// string slice to the right, that is, the beginning of the string slice.
	SwapRightStringSlice(args ...interface{}) ([]interface{}, error)

	// Type.
	ArgType(args ...interface{}) ([]interface{}, error)
}
