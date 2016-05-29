package log

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

func testMaybeNewLogControlAndServer(t *testing.T, handler http.Handler) (spec.LogControl, *httptest.Server) {
	ts := httptest.NewServer(handler)

	URL, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("expected", nil, "got", err)
	}

	newControlConfig := DefaultControlConfig()
	newControlConfig.URL = URL
	newControl, err := NewControl(newControlConfig)
	if err != nil {
		t.Fatalf("expected", nil, "got", err)
	}

	return newControl, ts
}

// reset levels

// Test_Log_LogControl_ResetLevels_001 checks for LogControl.ResetLevels to work
// properly under normal conditions using api.WithSuccess.
func Test_Log_LogControl_ResetLevels_001(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithSuccess()

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetLevels(ctx)
	if err != nil {
		t.Fatalf("LogControl.ResetLevels returned error: %#v", err)
	}
}

// Test_Log_LogControl_ResetLevels_002 checks for LogControl.ResetLevels to handle
// errors properly on valid error responses using api.WithError.
func Test_Log_LogControl_ResetLevels_002(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(errgo.Newf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetLevels(ctx)
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetLevels NOT returned proper error")
	}
}

// Test_Log_LogControl_ResetLevels_003 checks for LogControl.ResetLevels to handle
// errors properly on plain text responses.
func Test_Log_LogControl_ResetLevels_003(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetLevels(ctx)
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetLevels NOT returned proper error")
	}
}

// Test_Log_LogControl_ResetLevels_004 checks for LogControl.ResetLevels to handle
// errors properly on invalid JSON responses.
func Test_Log_LogControl_ResetLevels_004(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetLevels(ctx)
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetLevels NOT returned proper error")
	}
}

// reset objects

// Test_Log_LogControl_ResetObjects_005 checks for LogControl.ResetObjects to work
// properly under normal conditions using api.WithSuccess.
func Test_Log_LogControl_ResetObjects_005(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithSuccess()

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetObjects(ctx)
	if err != nil {
		t.Fatalf("LogControl.ResetObjects returned error: %#v", err)
	}
}

// Test_Log_LogControl_ResetObjects_006 checks for LogControl.ResetObjects to
// handle errors properly on valid error responses using api.WithError.
func Test_Log_LogControl_ResetObjects_006(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(errgo.Newf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetObjects(ctx)
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetObjects NOT returned proper error")
	}
}

// Test_Log_LogControl_ResetObjects_007 checks for LogControl.ResetObjects to
// handle errors properly on plain text responses.
func Test_Log_LogControl_ResetObjects_007(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetObjects(ctx)
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetObjects NOT returned proper error")
	}
}

// Test_Log_LogControl_ResetObjects_008 checks for LogControl.ResetObjects to
// handle errors properly on invalid JSON responses.
func Test_Log_LogControl_ResetObjects_008(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetObjects(ctx)
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetObjects NOT returned proper error")
	}
}

// reset verbosity

// Test_Log_LogControl_ResetVerbosity_009 checks for LogControl.ResetVerbosity to
// work properly under normal conditions using api.WithSuccess.
func Test_Log_LogControl_ResetVerbosity_009(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithSuccess()

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetVerbosity(ctx)
	if err != nil {
		t.Fatalf("LogControl.ResetVerbosity returned error: %#v", err)
	}
}

// Test_Log_LogControl_ResetVerbosity_010 checks for LogControl.ResetVerbosity to
// handle errors properly on valid error responses using api.WithError.
func Test_Log_LogControl_ResetVerbosity_010(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(errgo.Newf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetVerbosity(ctx)
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetVerbosity NOT returned proper error")
	}
}

// Test_Log_LogControl_ResetVerbosity_011 checks for LogControl.ResetVerbosity to
// handle errors properly on plain text responses.
func Test_Log_LogControl_ResetVerbosity_011(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetVerbosity(ctx)
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetVerbosity NOT returned proper error")
	}
}

// Test_Log_LogControl_ResetVerbosity_012 checks for LogControl.ResetVerbosity to
// handle errors properly on invalid JSON responses.
func Test_Log_LogControl_ResetVerbosity_012(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetVerbosity(ctx)
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetVerbosity NOT returned proper error")
	}
}

// set levels

// Test_Log_LogControl_SetLevels_013 checks for LogControl.SetLevels to work
// properly under normal conditions using api.WithSuccess.
func Test_Log_LogControl_SetLevels_013(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithSuccess()

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetLevels(ctx, "foo,bar")
	if err != nil {
		t.Fatalf("LogControl.SetLevels returned error: %#v", err)
	}
}

// Test_Log_LogControl_SetLevels_014 checks for LogControl.SetLevels to handle
// errors properly on valid error responses using api.WithError.
func Test_Log_LogControl_SetLevels_014(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(errgo.Newf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetLevels(ctx, "foo,bar")
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetLevels NOT returned proper error")
	}
}

// Test_Log_LogControl_SetLevels_015 checks for LogControl.SetLevels to handle
// errors properly on plain text responses.
func Test_Log_LogControl_SetLevels_015(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetLevels(ctx, "foo,bar")
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetLevels NOT returned proper error")
	}
}

// Test_Log_LogControl_SetLevels_016 checks for LogControl.SetLevels to handle
// errors properly on invalid JSON responses.
func Test_Log_LogControl_SetLevels_016(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetLevels(ctx, "foo,bar")
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetLevels NOT returned proper error")
	}
}

// set objects

// Test_Log_LogControl_SetObjects_017 checks for LogControl.SetObjects to work
// properly under normal conditions using api.WithSuccess.
func Test_Log_LogControl_SetObjects_017(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithSuccess()

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetObjects(ctx, "foo,bar")
	if err != nil {
		t.Fatalf("LogControl.SetObjects returned error: %#v", err)
	}
}

// Test_Log_LogControl_SetObjects_018 checks for LogControl.SetObjects to handle
// errors properly on valid error responses using api.WithError.
func Test_Log_LogControl_SetObjects_018(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(errgo.Newf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetObjects(ctx, "foo,bar")
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetObjects NOT returned proper error")
	}
}

// Test_Log_LogControl_SetObjects_019 checks for LogControl.SetObjects to handle
// errors properly on plain text responses.
func Test_Log_LogControl_SetObjects_019(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetObjects(ctx, "foo,bar")
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetObjects NOT returned proper error")
	}
}

// Test_Log_LogControl_SetObjects_020 checks for LogControl.SetObjects to handle
// errors properly on invalid JSON responses.
func Test_Log_LogControl_SetObjects_020(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetObjects(ctx, "foo,bar")
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetObjects NOT returned proper error")
	}
}

// set verbosity

// Test_Log_LogControl_SetVerbosity_021 checks for LogControl.SetVerbosity to work
// properly under normal conditions using api.WithSuccess.
func Test_Log_LogControl_SetVerbosity_021(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithSuccess()

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetVerbosity(ctx, 66)
	if err != nil {
		t.Fatalf("LogControl.SetVerbosity returned error: %#v", err)
	}
}

// Test_Log_LogControl_SetVerbosity_022 checks for LogControl.SetVerbosity to
// handle errors properly on valid error responses using api.WithError.
func Test_Log_LogControl_SetVerbosity_022(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(errgo.Newf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetVerbosity(ctx, 66)
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetVerbosity NOT returned proper error")
	}
}

// Test_Log_LogControl_SetVerbosity_023 checks for LogControl.SetVerbosity to
// handle errors properly on plain text responses.
func Test_Log_LogControl_SetVerbosity_023(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetVerbosity(ctx, 66)
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetVerbosity NOT returned proper error")
	}
}

// Test_Log_LogControl_SetVerbosity_024 checks for LogControl.SetVerbosity to
// handle errors properly on invalid JSON responses.
func Test_Log_LogControl_SetVerbosity_024(t *testing.T) {
	newLogControl, ts := testMaybeNewLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetVerbosity(ctx, 66)
	if !IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetVerbosity NOT returned proper error")
	}
}
