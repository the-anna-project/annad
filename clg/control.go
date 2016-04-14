package clg

import ()

func (i *clgIndex) IfControl(args ...interface{}) ([]interface{}, error) {
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

	conditionResults, err := i.CallCLGByName(append([]interface{}{condition}, conditionArgs...)...)
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
		actionResults, err := i.CallCLGByName(append([]interface{}{action}, actionArgs...)...)
		if err != nil {
			return nil, maskAny(err)
		}
		return actionResults, nil
	}

	return []interface{}{}, nil
}
