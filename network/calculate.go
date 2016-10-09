package network

import (
	"reflect"
)

// filterError removes the last element of the given list. Thus filterError
// must only be used if the last element returned by a CLG implements the error
// interface. In case the last element is a non-nil error, this error is
// returned and the given list is discarded.
func filterError(values []reflect.Value) ([]reflect.Value, error) {
	if len(values) == 0 {
		return nil, nil
	}

	lastArg := values[len(values)-1]
	if lastArg.Type().String() == "error" {
		if lastArg.IsValid() && !lastArg.IsNil() {
			switch lastArg.Kind() {
			case reflect.Interface:
				fallthrough
			case reflect.Ptr:
				if err, ok := lastArg.Interface().(error); ok {
					return nil, err
				}
			}
		} else {
			return values[:len(values)-1], nil
		}
	}

	return values, nil
}

// TODO test
//
// import (
// 	"fmt"
// 	"reflect"
// 	"testing"
// )
//
// func Test_CLG_filterError_noValues(t *testing.T) {
// 	out, err := filterError(nil)
// 	if err != nil {
// 		t.Fatal("expected", nil, "got", err)
// 	}
//
// 	if len(out) != 0 {
// 		t.Fatal("expected", 2, "got", len(out))
// 	}
// }
//
// func Test_CLG_filterError_noError(t *testing.T) {
// 	in := []reflect.Value{
// 		reflect.ValueOf("one"),
// 		reflect.ValueOf("two"),
// 		reflect.ValueOf("three"),
// 	}
//
// 	out, err := filterError(in)
// 	if err != nil {
// 		t.Fatal("expected", nil, "got", err)
// 	}
//
// 	if len(out) != 2 {
// 		t.Fatal("expected", 2, "got", len(out))
// 	}
// }
//
// func Test_CLG_filterError_nilError(t *testing.T) {
// 	var err error
// 	in := []reflect.Value{
// 		reflect.ValueOf("one"),
// 		reflect.ValueOf("two"),
// 		reflect.ValueOf(err),
// 	}
//
// 	out, err := filterError(in)
// 	if err != nil {
// 		t.Fatal("expected", nil, "got", err)
// 	}
//
// 	if len(out) != 2 {
// 		t.Fatal("expected", 2, "got", len(out))
// 	}
// }
//
// func Test_CLG_filterError_errgoError(t *testing.T) {
// 	in := []reflect.Value{
// 		reflect.ValueOf("one"),
// 		reflect.ValueOf("two"),
// 		reflect.ValueOf(invalidConfigError),
// 	}
//
// 	_, err := filterError(in)
// 	if !IsInvalidConfig(err) {
// 		t.Fatal("expected", nil, "got", err)
// 	}
// }
//
// func Test_CLG_filterError_fmtError(t *testing.T) {
// 	fmtErr := fmt.Errorf("test error")
// 	in := []reflect.Value{
// 		reflect.ValueOf("one"),
// 		reflect.ValueOf("two"),
// 		reflect.ValueOf(fmtErr),
// 	}
//
// 	_, err := filterError(in)
// 	if err == nil {
// 		t.Fatal("expected", fmtErr, "got", nil)
// 	}
// 	msg := err.Error()
// 	if msg != "test error" {
// 		t.Fatal("expected", "test error", "got", msg)
// 	}
// }
//
// func Test_CLG_filterError_CLG(t *testing.T) {
// 	in := []reflect.Value{
// 		reflect.ValueOf("one"),
// 		reflect.ValueOf("two"),
// 		reflect.ValueOf(MustNew()),
// 	}
//
// 	_, err := filterError(in)
// 	if err != nil {
// 		t.Fatal("expected", nil, "got", err)
// 	}
// }
