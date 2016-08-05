package text

import (
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func Test_TextInterface_getResponseForIDEncoder_Error(t *testing.T) {
	ctx := context.Background()
	r := &http.Request{}
	request := make(chan int) // marshaling channels throws json.UnsupportedTypeError

	err := getResponseForIDEncoder(ctx, r, request)
	if !IsUnsupportedType(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_TextInterface_readCoreRequestEncoder_Error(t *testing.T) {
	ctx := context.Background()
	r := &http.Request{}
	request := make(chan int) // marshaling channels throws json.UnsupportedTypeError

	err := readCoreRequestEncoder(ctx, r, request)
	if !IsUnsupportedType(err) {
		t.Fatal("expected", true, "got", false)
	}
}
