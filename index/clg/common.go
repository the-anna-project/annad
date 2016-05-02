package clg

import (
	"reflect"

	"github.com/xh3b4sd/anna/spec"
)

// TODO
func (i *clgIndex) getCLGInputExamples(methodValue reflect.Value) ([]interface{}, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call getCLGInputExamples")

	return nil, nil
}

// TODO
func (i *clgIndex) getCLGInputTypes(methodValue reflect.Value) ([]reflect.Kind, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call getCLGInputTypes")

	return nil, nil
}

// TODO
func (i *clgIndex) getCLGMethodHash(methodValue reflect.Value) (string, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call getCLGMethodHash")

	return "", nil
}

// TODO
func (i *clgIndex) getCLGRightSideNeighbours(clgCollection spec.CLGCollection, clgName string, methodValue reflect.Value, canceler <-chan struct{}) ([]string, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call getCLGRightSideNeighbours")

	// Fill a queue.
	// TODO create method for this getCLGMethodQueue
	args, err := clgCollection.GetNamesMethod()
	if err != nil {
		return nil, maskAny(err)
	}
	clgNames, err := ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	queue := make(chan string, len(clgNames))
	for _, clgName := range clgNames {
		queue <- clgName
	}

	//     find right side neighbours for given clg name
	//         if no profile for checked neighbour
	//             push neighbour name back to channel

	return nil, nil
}

func (i *clgIndex) isMethodValue(v reflect.Value) bool {
	if !v.IsValid() {
		return false
	}

	if v.Kind() != reflect.Func {
		return false
	}

	return true
}

// TODO
func (i *clgIndex) isRightSideCLGNeighbour(clgCollection spec.CLGCollection, left, right spec.CLGProfile) (bool, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 13}, "call isRightSideNeighbour")

	// run clg chain
	// if error
	//     return false

	return false, nil
}

// maybeReturnAndLogErrors returns the very first error that may be given by
// errors. All upcoming errors may be given by the provided error channel will
// be logged using the configured logger with log level E and verbosity 4.
func (i *clgIndex) maybeReturnAndLogErrors(errors chan error) error {
	var executeErr error

	for err := range errors {
		if IsWorkerCanceled(err) {
			continue
		}

		if executeErr == nil {
			// Only return the first error.
			executeErr = err
		} else {
			// Log all errors but the first one
			i.Log.WithTags(spec.Tags{L: "E", O: i, T: nil, V: 4}, "%#v", maskAny(err))
		}
	}

	if executeErr != nil {
		return maskAny(executeErr)
	}

	return nil
}
