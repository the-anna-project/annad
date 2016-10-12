package redis

import (
	"reflect"
	"testing"

	"github.com/rafaeljusto/redigomock"
)

func Test_ScoredSetStorage_GetElementsByScore_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGEBYSCORE", "prefix:foo", 0.8, 0.8, "LIMIT", 0, 3).Expect([]interface{}{[]uint8("bar")})

	newStorage := testMustNewStorageWithConn(t, c)

	values, err := newStorage.GetElementsByScore("foo", 0.8, 3)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(values) != 1 {
		t.Fatal("expected", 1, "got", len(values))
	}
	if values[0] != "bar" {
		t.Fatal("expected", "bar", "got", values[0])
	}
}

func Test_ScoredSetStorage_GetElementsByScore_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGEBYSCORE", "prefix:foo", 0.8, 0.8, "LIMIT", 0, 3).ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetElementsByScore("foo", 0.8, 3)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ScoredSetStorage_GetHighestScoredElements_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGE", "prefix:foo", 0, 2, "WITHSCORES").Expect([]interface{}{
		[]uint8("one"), []uint8("0.8"), []uint8("two"), []uint8("0.5"),
	})

	newStorage := testMustNewStorageWithConn(t, c)

	values, err := newStorage.GetHighestScoredElements("foo", 2)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if len(values) != 4 {
		t.Fatal("expected", 1, "got", len(values))
	}
	if values[0] != "one" {
		t.Fatal("expected", "one", "got", values[0])
	}
	if values[1] != "0.8" {
		t.Fatal("expected", "0.8", "got", values[1])
	}
	if values[2] != "two" {
		t.Fatal("expected", "two", "got", values[2])
	}
	if values[3] != "0.5" {
		t.Fatal("expected", "0.5", "got", values[3])
	}
}

func Test_ScoredSetStorage_GetHighestScoredElements_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREVRANGE", "prefix:foo", 0, 2, "WITHSCORES").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	_, err := newStorage.GetHighestScoredElements("foo", 2)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ScoredSetStorage_SetElementByScore_Success(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZADD", "prefix:key", 0.8, "element").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetElementByScore("key", "element", 0.8)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_ScoredSetStorage_SetElementByScore_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZADD", "prefix:key", 0.8, "element").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.SetElementByScore("key", "element", 0.8)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ScoredSetStorage_RemoveScoredElement(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREM", "prefix:test-key", "test-element").Expect(int64(1))

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.RemoveScoredElement("test-key", "test-element")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func Test_ScoredSetStorage_RemoveScoredElement_Error(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZREM", "prefix:test-key", "test-element").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.RemoveScoredElement("test-key", "test-element")
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ScoredSetStorage_WalkScoredSet(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("0.8"), []uint8("test-value-2"), []uint8("0.8")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var values []interface{}
	err := newStorage.WalkScoredSet("test-key", nil, func(element string, score float64) error {
		values = append(values, element, score)
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if !reflect.DeepEqual(values, []interface{}{"test-value-1", 0.8, "test-value-2", 0.8}) {
		t.Fatal("expected", []interface{}{"test-value-1", 0.8, "test-value-2", 0.8}, "got", values)
	}
}

func Test_ScoredSetStorage_WalkScoredSet_CloseDirectly(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("0.8"), []uint8("test-value-2"), []uint8("0.8")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	// Directly close and end walking.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	var values []interface{}
	err := newStorage.WalkScoredSet("test-key", closer, func(element string, score float64) error {
		values = append(values, element, score)
		return nil
	})
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	if values != nil {
		t.Fatal("expected", nil, "got", values)
	}
}

func Test_ScoredSetStorage_WalkScoredSet_CloseAfterCallback(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("0.8"), []uint8("test-value-2"), []uint8("0.8")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	var count int
	closer := make(chan struct{}, 1)

	err := newStorage.WalkScoredSet("test-key", closer, func(element string, score float64) error {
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

func Test_ScoredSetStorage_WalkScoredSet_QueryError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN").ExpectError(queryExecutionFailedError)

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkScoredSet("test-key", nil, nil)
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}

func Test_ScoredSetStorage_WalkScoredSet_CallbackError(t *testing.T) {
	c := redigomock.NewConn()
	c.Command("ZSCAN", "prefix:test-key", int64(0), "COUNT", 100).Expect([]interface{}{
		[]uint8("0"),
		[]interface{}{[]uint8("test-value-1"), []uint8("0.8"), []uint8("test-value-2"), []uint8("0.8")},
	})

	newStorage := testMustNewStorageWithConn(t, c)

	err := newStorage.WalkScoredSet("test-key", nil, func(element string, score float64) error {
		return maskAny(queryExecutionFailedError)
	})
	if !IsQueryExecutionFailed(err) {
		t.Fatal("expected", true, "got", false)
	}
}
