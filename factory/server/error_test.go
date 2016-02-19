// Note, this test hijacks the factoryserver's package scope. This is
// considered bad practise. We do it anyway to test error handling and to not
// expose it, since error handling is supposed to be package related and
// irrelevant for the public interface.
package factoryserver

import (
	"fmt"
	"testing"
)

func Test_FactoryServer_maskAnyf_001(t *testing.T) {
	testCases := []struct {
		InputError  error
		InputFormat string
		InputArgs   []interface{}
		Expected    error
	}{
		{
			InputError:  nil,
			InputFormat: "",
			InputArgs:   []interface{}{},
			Expected:    nil,
		},
		{
			InputError:  fmt.Errorf("foo"),
			InputFormat: "bar",
			InputArgs:   []interface{}{},
			Expected:    nil,
		},
		{
			InputError:  fmt.Errorf("foo"),
			InputFormat: "bar %s",
			InputArgs:   []interface{}{"baz"},
			Expected:    fmt.Errorf("foo: bar baz"),
		},
	}

	for _, testCase := range testCases {
		var output error
		if len(testCase.InputArgs) == 0 {
			output = maskAnyf(testCase.InputError, testCase.InputFormat)
		} else {
			output = maskAnyf(testCase.InputError, testCase.InputFormat, testCase.InputArgs...)
		}

		if testCase.Expected != nil && output.Error() != testCase.Expected.Error() {
			t.Fatalf("test case %d: output '%s' != expected '%s'", output, testCase.Expected)
		}
	}
}
