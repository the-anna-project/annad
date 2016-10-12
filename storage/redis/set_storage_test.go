package redis

import (
	"testing"

	"github.com/rafaeljusto/redigomock"
)

func Test_SetStorage_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SET", "prefix:foo", "bar").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.Set("foo", "bar")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_SetStorage_GetAllFromSet_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SMEMBERS", "prefix:foo").Expect([]interface{}{
		[]uint8("one"), []uint8("two"),
	})

	newStorage := testMustNewStorageWithConn(t, c)

	values, err := newStorage.GetAllFromSet("foo")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(values) != 2 {
		t.Fatal("expected", 1, "got", len(values))
	}
	if values[0] != "one" {
		t.Fatal("expected", "one", "got", values[0])
	}
	if values[1] != "two" {
		t.Fatal("expected", "two", "got", values[2])
	}
}

func Test_SetStorage_GetAllFromSet_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SMEMBERS", "prefix:foo").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetAllFromSet("foo")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_SetStorage_PushToSet(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SADD", "prefix:test-key", "test-element").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.PushToSet("test-key", "test-element")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_SetStorage_PushToSet_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SADD", "prefix:test-key", "test-element").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.PushToSet("test-key", "test-element")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_SetStorage_RemoveFromSet(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SREM", "prefix:test-key", "test-element").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.RemoveFromSet("test-key", "test-element")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_SetStorage_RemoveFromSet_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SREM", "prefix:test-key", "test-element").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.RemoveFromSet("test-key", "test-element")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_SetStorage_WalkSet(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var element1 string
	err := newStorage.WalkSet("test-key", nil, func(element string) error {
		element1 = element
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element1 != "test-value" {
		t.Fatal("expected", "test-value", "got", element1)
	}
}

func Test_SetStorage_WalkSet_CloseDirectly(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("test-value-2")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	// Directly close and end walking.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	var element1 string
	err := newStorage.WalkSet("test-key", closer, func(element string) error {
		element1 = element
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if element1 != "" {
		t.Fatal("expected", "", "got", element1)
	}
}

func Test_SetStorage_WalkSet_CloseAfterCallback(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("test-value-2")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var count int
	closer := make(chan struct{}, 1)

	err := newStorage.WalkSet("test-key", closer, func(element string) error {
		count++

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
}

func Test_SetStorage_WalkSet_QueryError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SSCAN").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkSet("test-key", nil, nil)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_SetStorage_WalkSet_CallbackError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("SSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("test-value-2")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkSet("test-key", nil, func(element string) error {
		return maskAny(queryExecutionFailedError)
	})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}
