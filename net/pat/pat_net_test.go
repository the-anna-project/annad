package patnet

import (
	"testing"
)

func Test_PatNet_BootShutdown_Single(t *testing.T) {
	newNet, err := NewPatNet(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newNet.Boot()
	newNet.Shutdown()
}

func Test_PatNet_BootShutdown_Multi(t *testing.T) {
	newNet, err := NewPatNet(DefaultConfig())
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	newNet.Boot()
	newNet.Boot()
	newNet.Boot()

	newNet.Shutdown()
	newNet.Shutdown()
	newNet.Shutdown()
}
