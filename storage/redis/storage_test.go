package redis

import (
	"testing"

	"github.com/garyburd/redigo/redis"

	"github.com/xh3b4sd/anna/spec"
)

func testMustNewStorageWithConn(t *testing.T, c redis.Conn) spec.Storage {
	newStorage, err := NewStorage(DefaultStorageConfigWithConn(c))
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newStorage
}

func Test_Storage_DefaultStorageConfigWithAddr(t *testing.T) {
	newStorageConfig := DefaultStorageConfigWithAddr("foo")
	_, err := newStorageConfig.Pool.Dial()
	if err.Error() != "dial tcp: missing port in address foo" {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Storage_NewRedisStorage_Error_BackOffFactory(t *testing.T) {
	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.BackOffFactory = nil
	_, err := NewStorage(newStorageConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_Storage_NewRedisStorage_Error_Log(t *testing.T) {
	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.Log = nil
	_, err := NewStorage(newStorageConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Storage_NewRedisStorage_Error_Pool(t *testing.T) {
	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.Pool = nil
	_, err := NewStorage(newStorageConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Storage_NewRedisStorage_Error_Prefix(t *testing.T) {
	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.Prefix = ""
	_, err := NewStorage(newStorageConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_Storage_Shutdown(t *testing.T) {
	newStorageConfig := DefaultStorageConfig()
	newStorage, err := NewStorage(newStorageConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newStorage.Shutdown()
	newStorage.Shutdown()
	newStorage.Shutdown()
}
