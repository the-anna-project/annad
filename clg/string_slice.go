package clg

func ContainsStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	str, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	var contains bool
	for _, s := range ss {
		if s == str {
			contains = true
			break
		}
	}

	return []interface{}{contains}, nil
}

// TODO JoinStringSlice

// TODO SortStringSlice

func SwapStringSlice(args ...interface{}) ([]interface{}, error) {
	ss, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(ss) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected at least 2 got %d", len(ss))
	}

	newStringSlice := append([]string{ss[len(ss)-1]}, ss[:len(ss)-1]...)

	return []interface{}{newStringSlice}, nil
}
