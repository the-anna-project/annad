package text

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

func testMaybeNewTextInterfaceAndServer(t *testing.T, handler http.Handler) (spec.TextInterface, *httptest.Server) {
	ts := httptest.NewServer(handler)

	URL, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newInterfaceConfig := DefaultInterfaceConfig()
	newInterfaceConfig.URL = URL
	newInterface, err := NewInterface(newInterfaceConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newInterface, ts
}

// read core request

func Test_Text_TextInterface_ReadCoreRequest_ReceiveID(t *testing.T) {
	responseID := "test-id"
	newTextInterface, ts := testMaybeNewTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithID(responseID)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	request := api.CoreRequest{
		Input: "test input",
	}
	sessionID := "session id"
	ID, err := newTextInterface.ReadCoreRequest(ctx, request, sessionID)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if ID != responseID {
		t.Fatal("expected", responseID, "got", ID)
	}
}

func Test_Text_TextInterface_ReadCoreRequest_ErrorResponse(t *testing.T) {
	newTextInterface, ts := testMaybeNewTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(errgo.Newf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	request := api.CoreRequest{
		Input: "test input",
	}
	sessionID := "session id"
	ID, err := newTextInterface.ReadCoreRequest(ctx, request, sessionID)
	if !IsInvalidAPIResponse(err) {
		t.Fatal("expected", nil, "got", err)
	}
	if ID != "" {
		t.Fatal("expected", "", "got", ID)
	}
}

func Test_Text_TextInterface_ReadCoreRequest_PlainTextError(t *testing.T) {
	newTextInterface, ts := testMaybeNewTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	request := api.CoreRequest{
		Input: "test input",
	}
	sessionID := "session id"
	ID, err := newTextInterface.ReadCoreRequest(ctx, request, sessionID)
	if !IsInvalidAPIResponse(err) {
		t.Fatal("expected", nil, "got", err)
	}
	if ID != "" {
		t.Fatal("expected", "", "got", ID)
	}
}

func Test_Text_TextInterface_ReadCoreRequest_InvalidJSON(t *testing.T) {
	newTextInterface, ts := testMaybeNewTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	request := api.CoreRequest{
		Input: "test input",
	}
	sessionID := "session id"
	ID, err := newTextInterface.ReadCoreRequest(ctx, request, sessionID)
	if !IsInvalidAPIResponse(err) {
		t.Fatal("expected", nil, "got", err)
	}
	if ID != "" {
		t.Fatal("expected", "", "got", ID)
	}
}

func Test_Text_TextInterface_ReadCoreRequest_RequestTransport(t *testing.T) {
	newTextInterface, ts := testMaybeNewTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request api.ReadCoreRequestRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		if request.CoreRequest.Input != "test input" {
			t.Fatal("expected", "test input", "got", request.CoreRequest.Input)
		}
		if request.SessionID != "session id" {
			t.Fatal("expected", "session id", "got", request.SessionID)
		}

		response := api.WithID("test-id")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	request := api.CoreRequest{
		Input: "test input",
	}
	sessionID := "session id"
	_, err := newTextInterface.ReadCoreRequest(ctx, request, sessionID)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

// get response for id

func Test_Text_TextInterface_GetResponseForID_ReceiveResponse(t *testing.T) {
	responseData := "hello world"
	newTextInterface, ts := testMaybeNewTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithData(responseData)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	data, err := newTextInterface.GetResponseForID(ctx, "test-id")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if data != responseData {
		t.Fatal("expected", responseData, "got", data)
	}
}

func Test_Text_TextInterface_GetResponseForID_ErrorResponse(t *testing.T) {
	newTextInterface, ts := testMaybeNewTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(errgo.Newf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	data, err := newTextInterface.GetResponseForID(ctx, "test-id")
	if !IsInvalidAPIResponse(err) {
		t.Fatal("expected", true, "got", false)
	}
	if data != "" {
		t.Fatal("expected", "", "got", data)
	}
}

func Test_Text_TextInterface_GetResponseForID_PlainTextError(t *testing.T) {
	newTextInterface, ts := testMaybeNewTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	data, err := newTextInterface.GetResponseForID(ctx, "test-id")
	if !IsInvalidAPIResponse(err) {
		t.Fatal("expected", true, "got", false)
	}
	if data != "" {
		t.Fatal("expected", "", "got", data)
	}
}

func Test_Text_TextInterface_GetResponseForID_InvalidJSON(t *testing.T) {
	newTextInterface, ts := testMaybeNewTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	data, err := newTextInterface.GetResponseForID(ctx, "test-id")
	if !IsInvalidAPIResponse(err) {
		t.Fatal("expected", true, "got", false)
	}
	if data != "" {
		t.Fatal("expected", "", "got", data)
	}
}

func Test_Text_TextInterface_GetResponseForID_RequestTransport(t *testing.T) {
	newTextInterface, ts := testMaybeNewTextInterfaceAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request api.GetResponseForIDRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Fatal("expected", nil, "got", err)
		}

		if request.ID != "test-id" {
			t.Fatal("expected", "", "got", request.ID)
		}

		response := api.WithData("test response")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal("expected", nil, "got", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	_, err := newTextInterface.GetResponseForID(ctx, "test-id")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}
