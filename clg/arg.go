package clg

func ArgToInt(args []interface{}, index int) (int, error) {
	if len(args) < index+1 {
		return 0, maskAnyf(notEnoughArgumentsError, "expected %d got %d", index+1, len(args))
	}

	if i, ok := args[index].(int); ok {
		return i, nil
	}

	return 0, maskAnyf(wrongArgumentTypeError, "expected %T got %T", "", args[index])
}

func ArgToString(args []interface{}, index int) (string, error) {
	if len(args) < index+1 {
		return "", maskAnyf(notEnoughArgumentsError, "expected %d got %d", index+1, len(args))
	}

	if s, ok := args[index].(string); ok {
		return s, nil
	}

	return "", maskAnyf(wrongArgumentTypeError, "expected %T got %T", "", args[index])
}

func ArgToStringSlice(args []interface{}, index int) ([]string, error) {
	if len(args) < index+1 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected %d got %d", index+1, len(args))
	}

	if ss, ok := args[index].([]string); ok {
		return ss, nil
	}

	return nil, maskAnyf(wrongArgumentTypeError, "expected %T got %T", "", args[index])
}
