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

func Test_StringStorage_WalkKeys(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SCAN", int64(0), "MATCH", "*", "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-key")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var count int
	var element1 string

	err := newStorage.WalkKeys("*", nil, func(key string) error {
		count++
		element1 = key
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if count != 1 {
		t.Fatal("expected", 1, "got", count)
	}
	if element1 != "test-key" {
		t.Fatal("expected", "test-key", "got", element1)
	}
}

func Test_StringStorage_WalkKeys_CloseDirectly(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SCAN", int64(0), "MATCH", "*", "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-key")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var count int
	// Directly close and end walking.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	err := newStorage.WalkKeys("*", closer, func(key string) error {
		count++
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if count != 0 {
		t.Fatal("expected", 0, "got", count)
	}
}

func Test_StringStorage_WalkKeys_CloseAfterCallback(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SCAN", int64(0), "MATCH", "*", "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-key")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var count int
	var element1 string
	closer := make(chan struct{}, 1)

	err := newStorage.WalkKeys("*", closer, func(key string) error {
		count++
		element1 = key

		// Close and end walking.
		closer <- struct{}{}

		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if count != 1 {
		t.Fatal("expected", 1, "got", count)
	}
	if element1 != "test-key" {
		t.Fatal("expected", "test-key", "got", element1)
	}
}

func Test_StringStorage_WalkKeys_QueryError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SCAN").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkKeys("*", nil, nil)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_StringStorage_WalkKeys_CallbackError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SCAN", int64(0), "MATCH", "*", "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-key")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkKeys("*", nil, func(key string) error {
		return maskAny(queryExecutionFailedError)
	})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}
