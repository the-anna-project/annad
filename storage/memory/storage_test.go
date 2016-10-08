package memory

import (
	"testing"
)

func Test_Memory_String_GetSetGet(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()

	_, err := newStorage.Get("foo")
	if !IsNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}

	err = newStorage.Set("foo", "bar")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	value, err := newStorage.Get("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if value != "bar" {
		t.Fatal("expected", "bar", "got", value)
	}
}

func Test_Memory_StringMap_GetSetGet(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()

	value, err := newStorage.GetStringMap("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(value) != 0 {
		t.Fatal("expected", 0, "got", len(value))
	}

	err = newStorage.SetStringMap("foo", map[string]string{"bar": "baz"})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	value, err = newStorage.GetStringMap("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(value) != 1 {
		t.Fatal("expected", 1, "got", len(value))
	}
	v, ok := value["bar"]
	if !ok {
		t.Fatal("expected", true, "got", false)
	}
	if v != "baz" {
		t.Fatal("expected", "baz", "got", v)
	}
}
