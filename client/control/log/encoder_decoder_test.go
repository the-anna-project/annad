package log

import (
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func Test_LogControl_resetLevelsEncoder_Error(t *testing.T) {
	ctx := context.Background()
	r := &http.Request{}
	request := make(chan int) // marshaling channels throws json.UnsupportedTypeError

	err := resetLevelsEncoder(ctx, r, request)
	if !IsUnsupportedType(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_LogControl_resetObjectsEncoder_Error(t *testing.T) {
	ctx := context.Background()
	r := &http.Request{}
	request := make(chan int) // marshaling channels throws json.UnsupportedTypeError

	err := resetObjectsEncoder(ctx, r, request)
	if !IsUnsupportedType(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_LogControl_resetVerbosityEncoder_Error(t *testing.T) {
	ctx := context.Background()
	r := &http.Request{}
	request := make(chan int) // marshaling channels throws json.UnsupportedTypeError

	err := resetVerbosityEncoder(ctx, r, request)
	if !IsUnsupportedType(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_LogControl_setLevelsEncoder_Error(t *testing.T) {
	ctx := context.Background()
	r := &http.Request{}
	request := make(chan int) // marshaling channels throws json.UnsupportedTypeError

	err := setLevelsEncoder(ctx, r, request)
	if !IsUnsupportedType(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_LogControl_setObjectsEncoder_Error(t *testing.T) {
	ctx := context.Background()
	r := &http.Request{}
	request := make(chan int) // marshaling channels throws json.UnsupportedTypeError

	err := setObjectsEncoder(ctx, r, request)
	if !IsUnsupportedType(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_LogControl_setVerbosityEncoder_Error(t *testing.T) {
	ctx := context.Background()
	r := &http.Request{}
	request := make(chan int) // marshaling channels throws json.UnsupportedTypeError

	err := setVerbosityEncoder(ctx, r, request)
	if !IsUnsupportedType(err) {
		t.Fatal("expected", true, "got", false)
	}
}
