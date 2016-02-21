package textinterface

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/juju/errgo"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

// helper

func newTextInterfaceAndServer(t *testing.T, handler http.Handler) (spec.TextInterface, *httptest.Server) {
	ts := httptest.NewServer(handler)

	URL, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("url.Parse returned error: %#v", err)
	}

	newTextInterfaceConfig := DefaultConfig()
	newTextInterfaceConfig.URL = URL
	newTextInterface := NewTextInterface(newTextInterfaceConfig)

	return newTextInterface, ts
}

// read plain with plain

// Test_TextInterface_ReadPlainWithPlain_001 checks for
// TextInterface.ReadPlainWithPlain to work properly under normal conditions
// using api.WithID.
func Test_TextInterface_ReadPlainWithPlain_001(t *testing.T) {
	responseID := "test-id"
	newTextInterface, ts := newTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithID(responseID)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	ID, err := newTextInterface.ReadPlainWithPlain(ctx, "hello world")
	if err != nil {
		t.Fatalf("TextInterface.ReadPlainWithPlain returned error: %#v", err)
	}
	if ID != responseID {
		t.Fatalf("expected response ID to be '%s', got '%s'", responseID, ID)
	}
}

// Test_TextInterface_ReadPlainWithPlain_002 checks for
// TextInterface.ReadPlainWithPlain to handle errors properly on valid error
// responses using api.WithError.
func Test_TextInterface_ReadPlainWithPlain_002(t *testing.T) {
	newTextInterface, ts := newTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(errgo.Newf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	ID, err := newTextInterface.ReadPlainWithPlain(ctx, "hello world")
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("TextInterface.ReadPlainWithPlain NOT returned proper error")
	}
	if ID != "" {
		t.Fatalf("expected response ID to be empty, got '%s'", ID)
	}
}

// Test_TextInterface_ReadPlainWithPlain_003 checks for
// TextInterface.ReadPlainWithPlain to handle errors properly on plain text
// responses.
func Test_TextInterface_ReadPlainWithPlain_003(t *testing.T) {
	newTextInterface, ts := newTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	ID, err := newTextInterface.ReadPlainWithPlain(ctx, "hello world")
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("TextInterface.ReadPlainWithPlain NOT returned proper error")
	}
	if ID != "" {
		t.Fatalf("expected response ID to be empty, got '%s'", ID)
	}
}

// Test_TextInterface_ReadPlainWithPlain_004 checks for
// TextInterface.ReadPlainWithPlain to handle errors properly on invalid JSON
// responses.
func Test_TextInterface_ReadPlainWithPlain_004(t *testing.T) {
	newTextInterface, ts := newTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	ID, err := newTextInterface.ReadPlainWithPlain(ctx, "hello world")
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("TextInterface.ReadPlainWithPlain NOT returned proper error")
	}
	if ID != "" {
		t.Fatalf("expected response ID to be empty, got '%s'", ID)
	}
}

// read plain with ID

// Test_TextInterface_ReadPlainWithID_005 checks for
// TextInterface.ReadPlainWithID to work properly under normal conditions.
func Test_TextInterface_ReadPlainWithID_005(t *testing.T) {
	responseData := "hello world"
	newTextInterface, ts := newTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithData(responseData)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	data, err := newTextInterface.ReadPlainWithID(ctx, "test-id")
	if err != nil {
		t.Fatalf("TextInterface.ReadPlainWithID returned error: %#v", err)
	}
	if data != responseData {
		t.Fatalf("expected response data to be '%s', got '%s'", responseData, data)
	}
}

// Test_TextInterface_ReadPlainWithID_006 checks for
// TextInterface.ReadPlainWithPlain to handle errors properly on valid error
// responses using api.WithError.
func Test_TextInterface_ReadPlainWithID_006(t *testing.T) {
	newTextInterface, ts := newTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(errgo.Newf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	data, err := newTextInterface.ReadPlainWithID(ctx, "test-id")
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("TextInterface.ReadPlainWithID NOT returned proper error")
	}
	if data != "" {
		t.Fatalf("expected response data to be empty, got '%s'", data)
	}
}

// Test_TextInterface_ReadPlainWithID_007 checks for
// TextInterface.ReadPlainWithID to handle errors properly on plain text
// responses.
func Test_TextInterface_ReadPlainWithID_007(t *testing.T) {
	newTextInterface, ts := newTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	data, err := newTextInterface.ReadPlainWithID(ctx, "test-id")
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("TextInterface.ReadPlainWithID NOT returned proper error")
	}
	if data != "" {
		t.Fatalf("expected response data to be empty, got '%s'", data)
	}
}

// Test_TextInterface_ReadPlainWithID_008 checks for
// TextInterface.ReadPlainWithID to handle errors properly on invalid JSON
// responses.
func Test_TextInterface_ReadPlainWithID_008(t *testing.T) {
	newTextInterface, ts := newTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	data, err := newTextInterface.ReadPlainWithID(ctx, "test-id")
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("TextInterface.ReadPlainWithID NOT returned proper error")
	}
	if data != "" {
		t.Fatalf("expected response data to be empty, got '%s'", data)
	}
}
