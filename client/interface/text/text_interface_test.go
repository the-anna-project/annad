package textinterface_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/client"
	serverspec "github.com/xh3b4sd/anna/server/spec"
)

func newTextInterfaceAndServer(t *testing.T, handler http.Handler) (serverspec.TextInterface, *httptest.Server) {
	ts := httptest.NewServer(handler)

	URL, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("url.Parse returned error: %#v", err)
	}

	newTextInterfaceConfig := client.DefaultTextInterfaceConfig()
	newTextInterfaceConfig.URL = URL
	newTextInterface := client.NewTextInterface(newTextInterfaceConfig)

	return newTextInterface, ts
}

// Test_Client_001 checks for TextInterface.ReadPlainWithPlain to work properly
// under normal conditions.
func Test_Client_001(t *testing.T) {
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

// Test_Client_002 checks for TextInterface.ReadPlainWithPlain to handle errors
// properly on plain text responses.
func Test_Client_002(t *testing.T) {
	newTextInterface, ts := newTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	ID, err := newTextInterface.ReadPlainWithPlain(ctx, "hello world")
	if !client.IsInvalidAPIResponse(err) {
		t.Fatalf("TextInterface.ReadPlainWithPlain NOT returned proper error")
	}
	if ID != "" {
		t.Fatalf("expected response ID to be empty, got '%s'", ID)
	}
}

// Test_Client_003 checks for TextInterface.ReadPlainWithPlain to handle errors
// properly on invalid JSON responses.
func Test_Client_003(t *testing.T) {
	newTextInterface, ts := newTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	ID, err := newTextInterface.ReadPlainWithPlain(ctx, "hello world")
	if !client.IsInvalidAPIResponse(err) {
		t.Fatalf("TextInterface.ReadPlainWithPlain NOT returned proper error")
	}
	if ID != "" {
		t.Fatalf("expected response ID to be empty, got '%s'", ID)
	}
}

// Test_Client_004 checks for TextInterface.ReadPlainWithID to work properly
// under normal conditions.
func Test_Client_004(t *testing.T) {
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

// Test_Client_005 checks for TextInterface.ReadPlainWithID to handle errors
// properly on plain text responses.
func Test_Client_005(t *testing.T) {
	newTextInterface, ts := newTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	data, err := newTextInterface.ReadPlainWithID(ctx, "test-id")
	if !client.IsInvalidAPIResponse(err) {
		t.Fatalf("TextInterface.ReadPlainWithID NOT returned proper error")
	}
	if data != "" {
		t.Fatalf("expected response data to be empty, got '%s'", data)
	}
}

// Test_Client_006 checks for TextInterface.ReadPlainWithID to handle errors
// properly on invalid JSON responses.
func Test_Client_006(t *testing.T) {
	newTextInterface, ts := newTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	data, err := newTextInterface.ReadPlainWithID(ctx, "test-id")
	if !client.IsInvalidAPIResponse(err) {
		t.Fatalf("TextInterface.ReadPlainWithID NOT returned proper error")
	}
	if data != "" {
		t.Fatalf("expected response data to be empty, got '%s'", data)
	}
}
