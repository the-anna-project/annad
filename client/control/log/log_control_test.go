package logcontrol_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/client/control/log"
	serverspec "github.com/xh3b4sd/anna/server/spec"
)

// helper

func newLogControlAndServer(t *testing.T, handler http.Handler) (serverspec.LogControl, *httptest.Server) {
	ts := httptest.NewServer(handler)

	URL, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("url.Parse returned error: %#v", err)
	}

	newLogControlConfig := logcontrol.DefaultConfig()
	newLogControlConfig.URL = URL
	newLogControl := logcontrol.NewLogControl(newLogControlConfig)

	return newLogControl, ts
}

// reset levels

// Test_LogControl_001 checks for LogControl.ResetLevels to work properly
// under normal conditions using api.WithSuccess.
func Test_LogControl_001(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

// Test_LogControl_002 checks for LogControl.ResetLevels to handle errors
// properly on valid error responses using api.WithError.
func Test_LogControl_002(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(fmt.Errorf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetLevels(ctx)
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetLevels NOT returned proper error")
	}
}

// Test_LogControl_003 checks for LogControl.ResetLevels to handle errors
// properly on plain text responses.
func Test_LogControl_003(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetLevels(ctx)
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetLevels NOT returned proper error")
	}
}

// Test_LogControl_004 checks for LogControl.ResetLevels to handle errors
// properly on invalid JSON responses.
func Test_LogControl_004(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetLevels(ctx)
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetLevels NOT returned proper error")
	}
}

// reset objects

// Test_LogControl_005 checks for LogControl.ResetObjects to work properly
// under normal conditions using api.WithSuccess.
func Test_LogControl_005(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

// Test_LogControl_006 checks for LogControl.ResetObjects to handle errors
// properly on valid error responses using api.WithError.
func Test_LogControl_006(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(fmt.Errorf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetObjects(ctx)
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetObjects NOT returned proper error")
	}
}

// Test_LogControl_007 checks for LogControl.ResetObjects to handle errors
// properly on plain text responses.
func Test_LogControl_007(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetObjects(ctx)
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetObjects NOT returned proper error")
	}
}

// Test_LogControl_008 checks for LogControl.ResetObjects to handle errors
// properly on invalid JSON responses.
func Test_LogControl_008(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetObjects(ctx)
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetObjects NOT returned proper error")
	}
}

// reset verbosity

// Test_LogControl_009 checks for LogControl.ResetVerbosity to work properly
// under normal conditions using api.WithSuccess.
func Test_LogControl_009(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

// Test_LogControl_010 checks for LogControl.ResetVerbosity to handle errors
// properly on valid error responses using api.WithError.
func Test_LogControl_010(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(fmt.Errorf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetVerbosity(ctx)
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetVerbosity NOT returned proper error")
	}
}

// Test_LogControl_011 checks for LogControl.ResetVerbosity to handle errors
// properly on plain text responses.
func Test_LogControl_011(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetVerbosity(ctx)
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetVerbosity NOT returned proper error")
	}
}

// Test_LogControl_012 checks for LogControl.ResetVerbosity to handle errors
// properly on invalid JSON responses.
func Test_LogControl_012(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.ResetVerbosity(ctx)
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.ResetVerbosity NOT returned proper error")
	}
}

// set levels

// Test_LogControl_013 checks for LogControl.SetLevels to work properly
// under normal conditions using api.WithSuccess.
func Test_LogControl_013(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

// Test_LogControl_014 checks for LogControl.SetLevels to handle errors
// properly on valid error responses using api.WithError.
func Test_LogControl_014(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(fmt.Errorf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetLevels(ctx, "foo,bar")
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetLevels NOT returned proper error")
	}
}

// Test_LogControl_015 checks for LogControl.SetLevels to handle errors
// properly on plain text responses.
func Test_LogControl_015(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetLevels(ctx, "foo,bar")
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetLevels NOT returned proper error")
	}
}

// Test_LogControl_016 checks for LogControl.SetLevels to handle errors
// properly on invalid JSON responses.
func Test_LogControl_016(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetLevels(ctx, "foo,bar")
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetLevels NOT returned proper error")
	}
}

// set objects

// Test_LogControl_017 checks for LogControl.SetObjects to work properly
// under normal conditions using api.WithSuccess.
func Test_LogControl_017(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

// Test_LogControl_018 checks for LogControl.SetObjects to handle errors
// properly on valid error responses using api.WithError.
func Test_LogControl_018(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(fmt.Errorf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetObjects(ctx, "foo,bar")
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetObjects NOT returned proper error")
	}
}

// Test_LogControl_019 checks for LogControl.SetObjects to handle errors
// properly on plain text responses.
func Test_LogControl_019(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetObjects(ctx, "foo,bar")
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetObjects NOT returned proper error")
	}
}

// Test_LogControl_020 checks for LogControl.SetObjects to handle errors
// properly on invalid JSON responses.
func Test_LogControl_020(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetObjects(ctx, "foo,bar")
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetObjects NOT returned proper error")
	}
}

// set verbosity

// Test_LogControl_021 checks for LogControl.SetVerbosity to work properly
// under normal conditions using api.WithSuccess.
func Test_LogControl_021(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

// Test_LogControl_022 checks for LogControl.SetVerbosity to handle errors
// properly on valid error responses using api.WithError.
func Test_LogControl_022(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := api.WithError(fmt.Errorf("test error"))

		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatalf("json.NewEncoder returned error: %#v", err)
		}
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetVerbosity(ctx, 66)
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetVerbosity NOT returned proper error")
	}
}

// Test_LogControl_023 checks for LogControl.SetVerbosity to handle errors
// properly on plain text responses.
func Test_LogControl_023(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "error")
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetVerbosity(ctx, 66)
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetVerbosity NOT returned proper error")
	}
}

// Test_LogControl_024 checks for LogControl.SetVerbosity to handle errors
// properly on invalid JSON responses.
func Test_LogControl_024(t *testing.T) {
	newLogControl, ts := newLogControlAndServer(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error": true}`)
	}))
	defer ts.Close()

	ctx := context.Background()
	err := newLogControl.SetVerbosity(ctx, 66)
	if !logcontrol.IsInvalidAPIResponse(err) {
		t.Fatalf("LogControl.SetVerbosity NOT returned proper error")
	}
}
