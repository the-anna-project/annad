package collection

func (c *collection) ForControl(args ...interface{}) ([]interface{}, error) {
	asl, err := ArgToArgsList(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	action, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}
	if len(asl) < 2 {
		return nil, maskAnyf(notEnoughArgumentsError, "expected 2 got %d", len(args))
	}

	var allResults []interface{}
	for _, as := range asl {
		rs, err := c.CallByNameMethod(append([]interface{}{action}, as...)...)
		if err != nil {
			return nil, maskAny(err)
		}

		allResults = append(allResults, rs...)
	}

	return []interface{}{allResults}, nil
}

func (c *collection) IfControl(args ...interface{}) ([]interface{}, error) {
	condition, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	conditionArgs, err := ArgToArgs(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	action, err := ArgToString(args, 2)
	if err != nil {
		return nil, maskAny(err)
	}
	actionArgs, err := ArgToArgs(args, 3)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 4 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 4 got %d", len(args))
	}

	conditionResults, err := c.CallByNameMethod(append([]interface{}{condition}, conditionArgs...)...)
	if err != nil {
		return nil, maskAny(err)
	}
	b, err := ArgToBool(conditionResults, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(conditionResults) > 1 {
		return nil, maskAnyf(tooManyResultsError, "expected 1 got %d", len(args))
	}
	if b {
		actionResults, err := c.CallByNameMethod(append([]interface{}{action}, actionArgs...)...)
		if err != nil {
			return nil, maskAny(err)
		}
		return actionResults, nil
	}

	return []interface{}{}, nil
}

func (c *collection) IfElseControl(args ...interface{}) ([]interface{}, error) {
	condition, err := ArgToString(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	conditionArgs, err := ArgToArgs(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	action, err := ArgToString(args, 2)
	if err != nil {
		return nil, maskAny(err)
	}
	actionArgs, err := ArgToArgs(args, 3)
	if err != nil {
		return nil, maskAny(err)
	}
	alternative, err := ArgToString(args, 4)
	if err != nil {
		return nil, maskAny(err)
	}
	alternativeArgs, err := ArgToArgs(args, 5)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 6 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 6 got %d", len(args))
	}

	conditionResults, err := c.CallByNameMethod(append([]interface{}{condition}, conditionArgs...)...)
	if err != nil {
		return nil, maskAny(err)
	}
	b, err := ArgToBool(conditionResults, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(conditionResults) > 1 {
		return nil, maskAnyf(tooManyResultsError, "expected 1 got %d", len(args))
	}
	if b {
		actionResults, err := c.CallByNameMethod(append([]interface{}{action}, actionArgs...)...)
		if err != nil {
			return nil, maskAny(err)
		}
		return actionResults, nil
	}

	alternativeResults, err := c.CallByNameMethod(append([]interface{}{alternative}, alternativeArgs...)...)
	if err != nil {
		return nil, maskAny(err)
	}
	return alternativeResults, nil
}
