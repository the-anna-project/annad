package redisstorage

import (
	"net"
	"testing"
)

func Test_Storage_NewDial_Error_Addr(t *testing.T) {
	newDialConfig := DefaultRedisDialConfig()
	newDialConfig.Addr = "foo"
	newDial := NewDial(newDialConfig)

	_, err := newDial()
	if e, ok := err.(*net.OpError); ok {
		if e.Op != "dial" {
			t.Fatal("expected", "dial", "got", e.Op)
		}
	} else {
		t.Fatal("expected", "*net.OpError", "got", err)
	}
}
