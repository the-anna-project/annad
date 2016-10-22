package redis

import (
	"testing"

	"github.com/garyburd/redigo/redis"
	"github.com/rafaeljusto/redigomock"
)

func Test_StringStorage_Get_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("GET", "prefix:foo").Expect("bar")

	newStorage := testMustNewStorageWithConn(t, c)

	value, err := newStorage.Get("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if value != "bar" {
		t.Fatal("expected", "bar", "got", value)
	}
}

func Test_StringStorage_Get_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("GET", "prefix:foo").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.Get("foo")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_Get_Error_NotFound(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("GET", "prefix:foo").ExpectError(redis.ErrNil)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.Get("foo")
	if !IsNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_GetRandom_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("RANDOMKEY").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetRandom()
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_GetRandom_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("RANDOMKEY").Expect("key1")

	newStorage := testMustNewStorageWithConn(t, c)

	randomKey, err := newStorage.GetRandom()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if randomKey != "key1" {
		t.Fatal("expected", "key1", "got", randomKey)
	}
}

func Test_StringStorage_Remove_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("DEL", "prefix:foo").Expect(int64(0))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.Remove("foo")
	if !IsNotFound(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_Remove_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("DEL", "prefix:foo").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.Remove("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_StringStorage_Set_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "prefix:foo", "bar").Expect("OK")

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.Set("foo", "bar")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_StringStorage_Set_NoSuccess(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "prefix:foo", "bar").Expect("invalid")

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.Set("foo", "bar")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}
