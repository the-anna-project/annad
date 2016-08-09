package text

import (
	"net/http"
	"testing"

	"golang.org/x/net/context"
)

func Test_TextInterface_streamTextEncoder_Error(t *testing.T) {
	ctx := context.Background()
	r := &http.Request{}
	request := make(chan int) // marshaling channels throws json.UnsupportedTypeError

	err := streamTextEncoder(ctx, r, request)
	if !IsUnsupportedType(err) {
		t.Fatal("expected", true, "got", false)
	}
}
