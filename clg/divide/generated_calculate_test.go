package divide

// This file is generated by the CLG generator. Don't edit it manually. The CLG
// generator is invoked by go generate. For more information about the usage of
// the CLG generator check https://github.com/xh3b4sd/clggen or have a look at
// the clg package. There is the go generate statement placed to invoke clggen.

import (
	"reflect"
	"testing"
)

func Test_CLG_Factory(t *testing.T) {
	newCLG := MustNew()
	if newCLG.Service() == nil {
		t.Fatal("expected", false, "got", true)
	}
}

func Test_CLG_Gateway(t *testing.T) {
	newCLG := MustNew()
	if newCLG.Gateway() == nil {
		t.Fatal("expected", false, "got", true)
	}
}

func Test_CLG_GetCalculate(t *testing.T) {
	newCLG := MustNew()
	if reflect.TypeOf(newCLG.GetCalculate()).Kind() != reflect.Func {
		t.Fatal("expected", false, "got", true)
	}
}

func Test_CLG_GetName(t *testing.T) {
	newCLG := MustNew()
	clgName := newCLG.GetName()
	if clgName != "divide" {
		t.Fatal("expected", "divide", "got", clgName)
	}
}

func Test_CLG_New_ServiceCollection_Error(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.ServiceCollection = nil
	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_New_GatewayCollection_Error(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.GatewayCollection = nil
	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_New_LogError(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.Log = nil
	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_New_StorageCollection_Error(t *testing.T) {
	newConfig := DefaultConfig()
	newConfig.StorageCollection = nil
	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_CLG_SetServiceCollection(t *testing.T) {
	newCLG := MustNew()
	var rawCLG *clg

	switch c := newCLG.(type) {
	case *clg:
		// all good
		rawCLG = newCLG.(*clg)
	default:
		t.Fatal("expected", "*clg", "got", c)
	}

	if rawCLG.ServiceCollection == nil {
		t.Fatal("expected", "spec.ServiceCollection", "got", nil)
	}

	newCLG.SetServiceCollection(nil)

	if rawCLG.ServiceCollection != nil {
		t.Fatal("expected", nil, "got", "spec.ServiceCollection")
	}
}

func Test_CLG_SetGatewayCollection(t *testing.T) {
	newCLG := MustNew()
	var rawCLG *clg

	switch c := newCLG.(type) {
	case *clg:
		// all good
		rawCLG = newCLG.(*clg)
	default:
		t.Fatal("expected", "*clg", "got", c)
	}

	if rawCLG.GatewayCollection == nil {
		t.Fatal("expected", "spec.GatewayCollection", "got", nil)
	}

	newCLG.SetGatewayCollection(nil)

	if rawCLG.GatewayCollection != nil {
		t.Fatal("expected", nil, "got", "spec.GatewayCollection")
	}
}

func Test_CLG_SetLog(t *testing.T) {
	newCLG := MustNew()
	var rawCLG *clg

	switch c := newCLG.(type) {
	case *clg:
		// all good
		rawCLG = newCLG.(*clg)
	default:
		t.Fatal("expected", "*clg", "got", c)
	}

	if rawCLG.Log == nil {
		t.Fatal("expected", "spec.Log", "got", nil)
	}

	newCLG.SetLog(nil)

	if rawCLG.Log != nil {
		t.Fatal("expected", nil, "got", "spec.Log")
	}
}

func Test_CLG_SetStorageCollection(t *testing.T) {
	newCLG := MustNew()
	var rawCLG *clg

	switch c := newCLG.(type) {
	case *clg:
		// all good
		rawCLG = newCLG.(*clg)
	default:
		t.Fatal("expected", "*clg", "got", c)
	}

	if rawCLG.StorageCollection == nil {
		t.Fatal("expected", "spec.StorageCollection", "got", nil)
	}

	newCLG.SetStorageCollection(nil)

	if rawCLG.StorageCollection != nil {
		t.Fatal("expected", nil, "got", "spec.StorageCollection")
	}
}

func Test_CLG_Storage(t *testing.T) {
	newCLG := MustNew()
	if newCLG.Storage() == nil {
		t.Fatal("expected", false, "got", true)
	}
}
