package clg

import (
	"reflect"
	"sort"
)

func (i *clgIndex) DiscardInterface(args ...interface{}) ([]interface{}, error) {
	return nil, nil
}

func (i *clgIndex) EqualInterface(args ...interface{}) ([]interface{}, error) {
	if len(args) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected 2 got %d", len(args))
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	t := reflect.DeepEqual(args[0], args[1])

	return []interface{}{t}, nil
}

func (i *clgIndex) InsertArgInterface(args ...interface{}) ([]interface{}, error) {
	scopeArgs, err := ArgToArgs(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	scopeArg, err := ArgToArg(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	indizes, err := ArgToIntSlice(args, 2)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 3 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 3 got %d", len(args))
	}

	seen := map[int]struct{}{}
	for _, indiz := range indizes {
		if _, ok := seen[indiz]; ok {
			return nil, maskAnyf(duplicatedMemberError, "members of %#v must be unique", indizes)
		}
		seen[indiz] = struct{}{}
	}
	sort.Ints(indizes)

	newArgs := scopeArgs
	for _, indiz := range indizes {
		if indiz > len(newArgs) {
			return nil, maskAny(indexOutOfRangeError)
		}
		newArgs = append(newArgs, 0)
		copy(newArgs[indiz+1:], newArgs[indiz:])
		newArgs[indiz] = scopeArg
	}

	return []interface{}{newArgs}, nil
}

func (i *clgIndex) ReturnInterface(args ...interface{}) ([]interface{}, error) {
	return args, nil
}

func (i *clgIndex) SwapInterface(args ...interface{}) ([]interface{}, error) {
	if len(args) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected 2 got %d", len(args))
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	args[0], args[1] = args[1], args[0]
	return args, nil
}

func (i *clgIndex) TypeInterface(args ...interface{}) ([]interface{}, error) {
	if len(args) < 1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected 1 got %d", len(args))
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	t := reflect.TypeOf(args[0]).String()

	return []interface{}{t}, nil
}
