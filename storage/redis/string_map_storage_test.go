package redis

import (
	"reflect"
	"testing"

	"github.com/rafaeljusto/redigomock"
)

func Test_StringMapStorage_GetStringMap_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HGETALL", "prefix:foo").Expect([]interface{}{[]byte("k1"), []byte("v1")})

	newStorage := testMustNewStorageWithConn(t, c)

	value, err := newStorage.GetStringMap("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if !reflect.DeepEqual(value, map[string]string{"k1": "v1"}) {
		t.Fatal("expected", map[string]string{"k1": "v1"}, "got", value)
	}
}

func Test_StringMapStorage_GetStringMap_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HGETALL", "prefix:foo").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetStringMap("foo")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringMapStorage_SetStringMap_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HMSET", "prefix:foo", "k1", "v1").Expect("OK")

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetStringMap("foo", map[string]string{"k1": "v1"})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_StringMapStorage_SetStringMap_NotOK(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HMSET", "prefix:foo", "k1", "v1").Expect("Not OK")

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetStringMap("foo", map[string]string{"k1": "v1"})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringMapStorage_SetStringMap_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("HMSET", "prefix:foo", "k1", "v1").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetStringMap("foo", map[string]string{"k1": "v1"})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}
