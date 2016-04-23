package spec

// CLGIndex represents the CLG index providing all CLGs that can be used. TODO
// add more CLGs.
type CLGIndex interface {
	CLGControl
	CLGConvert
	CLGFeature
	CLGFeatureSet
	CLGFloat64
	CLGFloat64Slice
	CLGInt
	CLGInterface
	CLGIntSlice
	CLGMethod
	// CLGNetwork
	Object
	CLGString
	CLGStringSlice
	// CLGOS
	// CLGTime

	// TODO Similarity. (manhatten-distance, distribution, syntactic similarity, semantic similarity, combined similarity)
	// TODO Sets for all slice types. (union, intersect, difference)
	// TODO Converting types (use also ArgTo... methods) (string to string slice, float to string, string to int, bool to int)
}

// CLGConvert represents all conversion CLGs that can be used.
type CLGConvert interface {
	// BoolStringConvert provides functionality of strconv.FormatBool.
	BoolStringConvert(args ...interface{}) ([]interface{}, error)

	// Float64StringConvert provides functionality of strconv.FormatFloat.
	Float64StringConvert(args ...interface{}) ([]interface{}, error)

	// IntStringConvert provides functionality of strconv.Itoa.
	IntStringConvert(args ...interface{}) ([]interface{}, error)

	// StringBoolConvert provides functionality of strconv.ParseBool.
	StringBoolConvert(args ...interface{}) ([]interface{}, error)

	// StringFloat64Convert provides functionality of strconv.ParseFloat.
	StringFloat64Convert(args ...interface{}) ([]interface{}, error)

	// StringIntConvert provides functionality of strconv.Atoi.
	StringIntConvert(args ...interface{}) ([]interface{}, error)
}

// CLGControl represents all control flow CLGs that can be used.
type CLGControl interface {
	// ForControl provides code flow functionallity of the for statement to
	// iterate over the given argument lists and applying some action to them
	// identified by a CLG name.
	ForControl(args ...interface{}) ([]interface{}, error)

	// IfControl provides code flow functionallity of the if statement. It
	// expects the following four arguments in the given order.
	//
	//     ConditionCLG, ConditionArgs, ActionCLG, ActionArgs
	//
	IfControl(args ...interface{}) ([]interface{}, error)

	// IfElseControl provides code flow functionallity of the if-else statement.
	// It expects the following six arguments in the given order.
	//
	//     ConditionCLG, ConditionArgs, ActionCLG, ActionArgs, AlternativeCLG, AlternativeArgs
	//
	IfElseControl(args ...interface{}) ([]interface{}, error)

	// TODO control flow ??? switch, goto, strategy modifier
}

// CLGFeatureSet represents all spec.FeatureSet compatible CLGs that can be
// used.
type CLGFeatureSet interface {
	// GetFeaturesFeatureSet returns all features detected during scanning of the
	// configured sequences.
	GetFeaturesFeatureSet(args ...interface{}) ([]interface{}, error)

	// GetFeaturesByCountFeatureSet returns all features detected during scanning
	// of the configured sequences, that occur the given times.
	GetFeaturesByCountFeatureSet(args ...interface{}) ([]interface{}, error)

	// GetFeaturesByLengthFeatureSet returns all features detected during
	// scanning of the configured sequences, that satisfy the given length.
	GetFeaturesByLengthFeatureSet(args ...interface{}) ([]interface{}, error)

	// GetFeaturesBySequenceFeatureSet returns all features detected during
	// scanning of the configured sequences, that satisfy the given glob pattern.
	GetFeaturesBySequenceFeatureSet(args ...interface{}) ([]interface{}, error)

	// GetMaxLengthFeatureSet returns the configured MaxLength of the given
	// feature set.
	GetMaxLengthFeatureSet(args ...interface{}) ([]interface{}, error)

	// GetMinLengthFeatureSet returns the configured MinLength of the given
	// feature set.
	GetMinLengthFeatureSet(args ...interface{}) ([]interface{}, error)

	// GetMinCountFeatureSet returns the configured MinCount of the given
	// feature set.
	GetMinCountFeatureSet(args ...interface{}) ([]interface{}, error)

	// GetNewFeatureSet provides a way to create a new spec.FeatureSet.
	// Optionally decent configuration can be given.
	GetNewFeatureSet(args ...interface{}) ([]interface{}, error)

	// GetSeparatorFeatureSet returns the configured Separator of the given
	// feature set.
	GetSeparatorFeatureSet(args ...interface{}) ([]interface{}, error)

	// GetSequencesFeatureSet returns the configured Sequences of the given
	// feature set.
	GetSequencesFeatureSet(args ...interface{}) ([]interface{}, error)
}

// CLGFeature represents all spec.Feature compatible CLGs that can be used.
type CLGFeature interface {
	// AddPositionFeature provides a way to add a position to the given the
	// spec.Feature.
	AddPositionFeature(args ...interface{}) ([]interface{}, error)

	// GetCountFeature returns the spec.Feature's configured count. That is, the
	// number of configured positions.
	GetCountFeature(args ...interface{}) ([]interface{}, error)

	// GetDistributionFeature returns the spec.Feature's configured
	// spec.Distribution.
	GetDistributionFeature(args ...interface{}) ([]interface{}, error)

	// GetNewFeature provides a way to create a new spec.Feature. Optionally
	// decent configuration can be given.
	GetNewFeature(args ...interface{}) ([]interface{}, error)

	// GetPositionsFeature returns the spec.Feature's configured positions.
	GetPositionsFeature(args ...interface{}) ([]interface{}, error)

	// GetSequenceFeature returns the spec.Feature's configured sequence.
	GetSequenceFeature(args ...interface{}) ([]interface{}, error)
}

// CLGFloat64 represents all float64 compatible CLGs that can be used.
type CLGFloat64 interface {
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

	// RoundFloat64 returns a number rounded using the given precision.
	RoundFloat64(args ...interface{}) ([]interface{}, error)

	// SqrtFloat64 provides functionality of math.Sqrt, but for float64s.
	SqrtFloat64(args ...interface{}) ([]interface{}, error)

	// SubtractFloat64 creates the difference of the given float64s.
	SubtractFloat64(args ...interface{}) ([]interface{}, error)

	// SumFloat64 creates the sum of the given float64s.
	SumFloat64(args ...interface{}) ([]interface{}, error)
}

// CLGFloat64Slice represents all float64 slice compatible CLGs that can be used.
type CLGFloat64Slice interface {
	// AppendFloat64Slice provides functionality of append for float64 slices.
	AppendFloat64Slice(args ...interface{}) ([]interface{}, error)

	// ContainsFloat64Slice provides functionality to check if a float64 slice
	// contains a certain member.
	ContainsFloat64Slice(args ...interface{}) ([]interface{}, error)

	// CountFloat64Slice returns the number of elements in args.
	CountFloat64Slice(args ...interface{}) ([]interface{}, error)

	// EqualMatcherFloat64Slice takes a float64 slice and a float64. It then
	// returns two float64 slices, where the first one contains all items
	// matching the given float64, and the second float64 slice contains all
	// items not matching the given float64.
	EqualMatcherFloat64Slice(args ...interface{}) ([]interface{}, error)

	// GlobMatcherFloat64Slice takes a float64 slice and a float64. It then
	// returns two float64 slices, where the first one contains all items
	// satisfying strings.Contains after converting the ints to strings, and the
	// second float64 slice contains all items not satisfying strings.Contains
	// after converting the ints to strings.
	GlobMatcherFloat64Slice(args ...interface{}) ([]interface{}, error)

	// IndexFloat64Slice returns the element under the given index.
	IndexFloat64Slice(args ...interface{}) ([]interface{}, error)

	// IsUniqueFloat64Slice checks whether the given float64 slice only contains
	// unique members.
	IsUniqueFloat64Slice(args ...interface{}) ([]interface{}, error)

	// MaxFloat64Slice returns the highest number of a list.
	MaxFloat64Slice(args ...interface{}) ([]interface{}, error)

	// MinFloat64Slice returns the lowest number of a list.
	MinFloat64Slice(args ...interface{}) ([]interface{}, error)

	// NewFloat64Slice returns a new float64 slice.
	NewFloat64Slice(args ...interface{}) ([]interface{}, error)

	// SortFloat64Slice provides functionality of strings.Contains.
	SortFloat64Slice(args ...interface{}) ([]interface{}, error)

	// SwapLeftFloat64Slice provides functionality to move the first member of a
	// float64 slice to the left, that is, the end of the float64 slice.
	SwapLeftFloat64Slice(args ...interface{}) ([]interface{}, error)

	// SwapRightFloat64Slice provides functionality to move the last member of a
	// float64 slice to the right, that is, the beginning of the float64 slice.
	SwapRightFloat64Slice(args ...interface{}) ([]interface{}, error)

	// UniqueFloat64Slice returns an float64 slice only having unique members.
	UniqueFloat64Slice(args ...interface{}) ([]interface{}, error)
}

// CLGInt represents all int compatible CLGs that can be used.
type CLGInt interface {
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
}

// CLGInterface represents all interface compatible CLGs that can be used.
type CLGInterface interface {
	// DiscardInterface does nothing. It discards the given arguments and returns
	// nil.
	DiscardInterface(args ...interface{}) ([]interface{}, error)

	// EqualInterface provides functionality of reflect.DeepEqual.
	EqualInterface(args ...interface{}) ([]interface{}, error)

	// InsertArgInterface inserts certain arguments in a specific order and
	// returns the manipulated arguments.
	InsertArgInterface(args ...interface{}) ([]interface{}, error)

	// ReturnInterface returns the arguments it receives.
	ReturnInterface(args ...interface{}) ([]interface{}, error)

	// SwapInterface returns the two arguments it receives, but in swapped order.
	SwapInterface(args ...interface{}) ([]interface{}, error)

	// TypeInterface returns the string representation of the given arg's type.
	TypeInterface(args ...interface{}) ([]interface{}, error)
}

// CLGIntSlice represents all int slice compatible CLGs that can be used.
type CLGIntSlice interface {
	// AppendIntSlice provides functionality of append for int slices.
	AppendIntSlice(args ...interface{}) ([]interface{}, error)

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

	// IsUniqueIntSlice checks whether the given int slice only contains unique
	// members.
	IsUniqueIntSlice(args ...interface{}) ([]interface{}, error)

	// JoinIntSlice provides functionality of strings.Join after converting ints
	// to strings.
	JoinIntSlice(args ...interface{}) ([]interface{}, error)

	// MaxIntSlice returns the highest number of a list.
	MaxIntSlice(args ...interface{}) ([]interface{}, error)

	// MinIntSlice returns the lowest number of a list.
	MinIntSlice(args ...interface{}) ([]interface{}, error)

	// NewIntSlice returns a new int slice.
	NewIntSlice(args ...interface{}) ([]interface{}, error)

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
}

// CLGMethod represents all CLGs that can be used to operate on CLG methods.
type CLGMethod interface {
	// CallMethodByName provides a way to dynamically execute any index's CLG by
	// providing the method's name as string argument.
	CallMethodByName(args ...interface{}) ([]interface{}, error)

	// GetMethodNames provides a way to fetch all method names. Optionally a glob
	// pattern can be provided to limit the returned names.
	GetMethodNames(args ...interface{}) ([]interface{}, error)

	// GetNumMethods returns the number of available methods the CLG index
	// provides.
	GetNumMethods(args ...interface{}) ([]interface{}, error)
}

// CLGString represents all string compatible CLGs that can be used.
type CLGString interface {
	// ContainsString provides functionality of strings.Contains.
	ContainsString(args ...interface{}) ([]interface{}, error)

	// CountCharacterString returns a map of characters expressing their
	// corresponding occurrence count within the given string.
	CountCharacterString(args ...interface{}) ([]interface{}, error)

	// EditDistanceString implementes the Levenshtein distance to measure
	// similarity between two strings. Here all edit operations are weighted with
	// the cost 1. See http://en.wikipedia.org/wiki/Levenshtein_distance.
	EditDistanceString(args ...interface{}) ([]interface{}, error)

	// ContainsString provides functionality to check if one string is longer
	// than the other.
	LongerString(args ...interface{}) ([]interface{}, error)

	// NewIDString creates a new random hash using the id package.
	NewIDString(args ...interface{}) ([]interface{}, error)

	// ContainsString provides functionality of strings.Repeat.
	RepeatString(args ...interface{}) ([]interface{}, error)

	// ReverseString reverses the order of characters of the given string.
	ReverseString(args ...interface{}) ([]interface{}, error)

	// ContainsString provides functionality to check if one string is shorter
	// than the other.
	ShorterString(args ...interface{}) ([]interface{}, error)

	// SplitString provides functionality of strings.Split.
	SplitString(args ...interface{}) ([]interface{}, error)

	// SplitEqualString splits the given string into a given number of equal
	// parts.
	SplitEqualString(args ...interface{}) ([]interface{}, error)

	// ToLowerString provides functionality of strings.ToLower.
	ToLowerString(args ...interface{}) ([]interface{}, error)

	// ToUpperString provides functionality of strings.ToUpper.
	ToUpperString(args ...interface{}) ([]interface{}, error)
}

// CLGStringSlice represents all string slice compatible CLGs that can be used.
type CLGStringSlice interface {
	// AppendStringSlice provides functionality of append for string slices.
	AppendStringSlice(args ...interface{}) ([]interface{}, error)

	// ContainsStringSlice provides functionality to check if a string slice
	// contains a certain member.
	ContainsStringSlice(args ...interface{}) ([]interface{}, error)

	// CountCharacterStringSlice returns a map of characters expressing their
	// corresponding occurrence count within the given string slice.
	CountCharacterStringSlice(args ...interface{}) ([]interface{}, error)

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

	// IsUniqueStringSlice checks whether the given string slice only contains
	// unique members.
	IsUniqueStringSlice(args ...interface{}) ([]interface{}, error)

	// JoinStringSlice provides functionality of strings.Join.
	JoinStringSlice(args ...interface{}) ([]interface{}, error)

	// NewStringSlice returns a new string slice.
	NewStringSlice(args ...interface{}) ([]interface{}, error)

	// ReverseStringSlice reverses the order of the given list.
	ReverseStringSlice(args ...interface{}) ([]interface{}, error)

	// StemStringSlice returns the word stem that all words provided by the given
	// list have in common. Having the following list.
	//
	//     abc
	//     abcd
	//     abcde
	//     abcdef
	//
	// Results in the word stem "abc".
	//
	StemStringSlice(args ...interface{}) ([]interface{}, error)

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
}
