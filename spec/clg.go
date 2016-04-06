package spec

// Index represents the CLG index providing all CLGs that can be used.
type Index interface {
	// CLG.

	// CallCLGByName provides a way to dynamically execute any index's CLG by
	// providing the CLG's name as string argument.
	CallCLGByName(args ...interface{}) ([]interface{}, error)

	// GetCLGNames provides a way to fetch all CLG names. Optionally a glob
	// pattern can be provided to limit the returned names.
	GetCLGNames(args ...interface{}) ([]interface{}, error)

	// Int.

	// DivideInt creates the difference of the given ints.
	DivideInt(args ...interface{}) ([]interface{}, error)

	// GreaterThanInt returns the int that is greater than the other.
	GreaterThanInt(args ...interface{}) ([]interface{}, error)

	// LesserThanInt returns the int that is lesser than the other.
	LesserThanInt(args ...interface{}) ([]interface{}, error)

	// MultiplyInt creates the product of the given ints.
	MultiplyInt(args ...interface{}) ([]interface{}, error)

	// PowInt provides functionality of math.Pow, but for ints.
	PowInt(args ...interface{}) ([]interface{}, error)

	// SqrtInt provides functionality of math.Sqrt, but for ints.
	SqrtInt(args ...interface{}) ([]interface{}, error)

	// SubtractInt creates the difference of the given ints.
	SubtractInt(args ...interface{}) ([]interface{}, error)

	// SumInt creates the sum of the given ints.
	SumInt(args ...interface{}) ([]interface{}, error)

	// Int Slice.

	// ContainsIntSlice provides functionality to check if a int slice
	// contains a certain member.
	ContainsIntSlice(args ...interface{}) ([]interface{}, error)

	// CountIntSlice returns the number of elements in args.
	CountIntSlice(args ...interface{}) ([]interface{}, error)

	// EqualMatcherIntSlice takes a int slice and a int. It then returns two int
	// slices, where the first one contains all items matching the given int, and
	// the second int slice contains all items not matching the given int.
	EqualMatcherIntSlice(args ...interface{}) ([]interface{}, error)

	// GlobMatcherIntSlice takes a int slice and a int. It then returns two int
	// slices, where the first one contains all items satisfying strings.Contains
	// after converting the ints to strings, and the second int slice contains
	// all items not satisfying strings.Contains after converting the ints to
	// strings.
	GlobMatcherIntSlice(args ...interface{}) ([]interface{}, error)

	// IndexIntSlice returns the element under the given index.
	IndexIntSlice(args ...interface{}) ([]interface{}, error)

	// JoinIntSlice provides functionality of strings.Join after converting ints
	// to strings.
	JoinIntSlice(args ...interface{}) ([]interface{}, error)

	// MaxIntSlice returns the highest number of a list.
	MaxIntSlice(args ...interface{}) ([]interface{}, error)

	// MinIntSlice returns the lowest number of a list.
	MinIntSlice(args ...interface{}) ([]interface{}, error)

	// SortIntSlice provides functionality of strings.Contains.
	SortIntSlice(args ...interface{}) ([]interface{}, error)

	// SwapLeftIntSlice provides functionality to move the first member of a
	// int slice to the left, that is, the end of the int slice.
	SwapLeftIntSlice(args ...interface{}) ([]interface{}, error)

	// SwapRightIntSlice provides functionality to move the last member of a
	// int slice to the right, that is, the beginning of the int slice.
	SwapRightIntSlice(args ...interface{}) ([]interface{}, error)

	// UniqueIntSlice returns an int slice only having unique members.
	UniqueIntSlice(args ...interface{}) ([]interface{}, error)
	// TODO Set. (union, intersect, difference)

	// Float64.

	// DivideFloat64 creates the difference of the given float64s.
	DivideFloat64(args ...interface{}) ([]interface{}, error)

	// GreaterThanFloat64 returns the int that is greater than the other.
	GreaterThanFloat64(args ...interface{}) ([]interface{}, error)

	// LesserThanFloat64 returns the int that is lesser than the other.
	LesserThanFloat64(args ...interface{}) ([]interface{}, error)

	// MultiplyFloat64 creates the product of the given float64s.
	MultiplyFloat64(args ...interface{}) ([]interface{}, error)

	// PowFloat64 provides functionality of math.Pow, but for float64s.
	PowFloat64(args ...interface{}) ([]interface{}, error)

	// SqrtFloat64 provides functionality of math.Sqrt, but for float64s.
	SqrtFloat64(args ...interface{}) ([]interface{}, error)

	// SubtractFloat64 creates the difference of the given float64s.
	SubtractFloat64(args ...interface{}) ([]interface{}, error)

	// SumFloat64 creates the sum of the given float64s.
	SumFloat64(args ...interface{}) ([]interface{}, error)

	// TODO Network.

	// Object.

	Object

	// TODO OS.

	// TODO Pattern. (stem, distance, distribution)
	// TODO we need more similarity measures

	// String.

	// ContainsString provides functionality of strings.Contains.
	ContainsString(args ...interface{}) ([]interface{}, error)

	// ContainsString provides functionality to check if one string is longer
	// than the other.
	LongerString(args ...interface{}) ([]interface{}, error)

	// NewIDString creates a new random hash using the id package.
	NewIDString(args ...interface{}) ([]interface{}, error)

	// ContainsString provides functionality of strings.Repeat.
	RepeatString(args ...interface{}) ([]interface{}, error)

	// ContainsString provides functionality to check if one string is shorter
	// than the other.
	ShorterString(args ...interface{}) ([]interface{}, error)

	// ContainsString provides functionality of strings.Split.
	SplitString(args ...interface{}) ([]interface{}, error)

	// ToLowerString provides functionality of strings.ToLower.
	ToLowerString(args ...interface{}) ([]interface{}, error)

	// ToUpperString provides functionality of strings.ToUpper.
	ToUpperString(args ...interface{}) ([]interface{}, error)

	// String slice.

	// ContainsStringSlice provides functionality to check if a string slice
	// contains a certain member.
	ContainsStringSlice(args ...interface{}) ([]interface{}, error)

	// CountStringSlice returns the number of elements in args.
	CountStringSlice(args ...interface{}) ([]interface{}, error)

	// EqualMatcherStringSlice takes a string slice and a string. It then returns
	// two string slices, where the first one contains all items matching the
	// given string, and the second string slice contains all items not matching
	// the given string.
	EqualMatcherStringSlice(args ...interface{}) ([]interface{}, error)

	// GlobMatcherStringSlice takes a string slice and a string. It then returns
	// two string slices, where the first one contains all items satisfying
	// strings.Contains, and the second string slice contains all items not
	// satisfying strings.Contains.
	GlobMatcherStringSlice(args ...interface{}) ([]interface{}, error)

	// IndexStringSlice returns the element under the given index.
	IndexStringSlice(args ...interface{}) ([]interface{}, error)

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

	// UniqueStringSlice returns a string slice only having unique members.
	UniqueStringSlice(args ...interface{}) ([]interface{}, error)

	// TODO Time.

	// Interface.

	// DiscardInterface does nothing. It discards the given arguments and returns
	// nil.
	DiscardInterface(args ...interface{}) ([]interface{}, error)

	// EqualInterface provides functionality of reflect.DeepEqual.
	EqualInterface(args ...interface{}) ([]interface{}, error)

	// TypeInterface returns the string representation of the given arg's type.
	TypeInterface(args ...interface{}) ([]interface{}, error)

	// TODO control flow ??? for, if, else, switch, goto, strategy modifier
	// TODO knowledge network ??? create, delete, find peers and connections
}
